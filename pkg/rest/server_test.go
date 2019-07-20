package rest

import (
	"testing"

	"github.com/franela/goblin"

	"encoding/base64"
	"fmt"
	"time"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"

	"github.com/studtool/go-errs/pkg/errs"
	"github.com/studtool/go-logs/pkg/logs"

	"github.com/studtool/go-rest/pkg/rest/tests"
)

var (
	//ldflags
	//nolint:gochecknoglobals
	testServerAddress = ""
)

func TestServer(t *testing.T) {
	g := goblin.Goblin(t)

	proto := "http"
	pathPrefix := "/api/vX"

	testCases := []testCase{
		{
			path:   "/get",
			method: MethodGet,
			handler: func(ctx *Context) {
				cookie := Cookie{
					Key:      tests.TestCookie,
					Value:    "some_request_value",
					Expire:   time.Now().Add(time.Hour),
					HTTPOnly: true,
					Secure:   false,
				}
				ctx.SetResponseCookie(&cookie)

				v := tests.TestingJSONType{
					SomeField: "some_response_value",
				}
				ctx.WriteOKJSON(v)
			},
			expectedStatus: StatusOK,
		},
		{
			path:   "/post/:id",
			method: MethodPost,
			handler: func(ctx *Context) {
				_ = ctx.PathParam("id")
				if ctx.RequestCookie(tests.TestCookie) == "" {
					ctx.WriteErrJSON(errs.NewBadFormat(5, "bad format"))
					return
				}
				var v tests.TestingJSONType
				err := ctx.ReadJSON(&v,
					errs.NewInvalidFormat(0, "invalid format"),
				)
				if err != nil {
					ctx.WriteErrJSON(err)
					return
				}
				ctx.WriteErrJSON(errs.NewNotFound(0, "not found"))
			},
			middleware: []Middleware{
				WithBasicAuth(BasicAuthMiddlewareParams{
					KnownClients: map[string]string{
						tests.TestClientID: tests.TestClientSecret,
					},
					AuthErr: errs.NewPermissionDenied(0, "unknown client"),
				}),
			},
			withBasicAuth:     true,
			withRequestBody:   true,
			withRequestCookie: true,
			expectedStatus:    StatusNotFound,
		},
		{
			path:   "/panic",
			method: MethodPatch,
			handler: func(ctx *Context) {
				panic("test-error")
			},
			middleware: []Middleware{
				WithAuth,
			},
			withAuth:       true,
			expectedStatus: StatusInternalServerError,
		},
	}

	handlers := make(map[string][]RequestHandler)
	for _, testCase := range testCases {
		handlers[testCase.path] = []RequestHandler{
			{
				Method:     testCase.method,
				Handler:    testCase.handler,
				Middleware: testCase.middleware,
			},
		}
	}

	routerConfig := RouterConfig{
		PathPrefix: pathPrefix,
		Handlers:   handlers,
		Middleware: []Middleware{
			WithRecover,
			WithLogs(LogsMiddlewareParams{
				RequestTypeDetector: func(ctx *Context) string {
					return RequestTypeTesting
				},
				RequestACLGroupDetector: func(ctx *Context) string {
					return ACLGroupCommon
				},
			}),
		},
	}

	srv := NewServer(ServerParams{
		Address:      testServerAddress,
		RouterConfig: routerConfig,
		LogPack: LogPack{
			StructLogger:     logs.NewStructLogger(logs.StructLoggerParams{}),
			StackTraceLogger: logs.NewStackTraceLogger(logs.StackTraceLoggerParams{}),
			RequestLogger:    logs.NewRequestLogger(logs.RequestLoggerParams{}),
		},
	})

	if err := srv.Run(); err != nil {
		panic(err)
	}
	defer func() {
		if err := srv.Shutdown(); err != nil {
			panic(err)
		}
	}()

	g.Describe("Server", func() {
		for _, tCase := range testCases {
			tCase.basePath = fmt.Sprintf(
				"%s://%s%s", proto, srv.address, pathPrefix,
			)
			doTest(g, tCase)
		}
	})
}

type testCase struct {
	path              string
	basePath          string
	method            string
	handler           Handler
	middleware        []Middleware
	withAuth          bool
	withBasicAuth     bool
	withRequestBody   bool
	withRequestCookie bool
	expectedStatus    int
}

func doTest(
	g *goblin.G,
	tCase testCase,
) {
	name := fmt.Sprintf("%s %s: should return %d",
		tCase.method, tCase.path, tCase.expectedStatus,
	)
	g.It(name, func() {
		status, _ := doRequest(tCase)
		g.Assert(status).Equal(tCase.expectedStatus)
	})

	if tCase.withAuth {
		name = fmt.Sprintf("%s %s: should return %d without %s header",
			tCase.method, tCase.path, StatusUnauthorized, HeaderXUserID,
		)
		g.It(name, func() {
			c := tCase
			c.withAuth = false

			status, _ := doRequest(c)
			g.Assert(status).Equal(StatusUnauthorized)
		})
	}

	if tCase.withBasicAuth {
		name = fmt.Sprintf("%s %s: should return %d without BasicAuth",
			tCase.method, tCase.path, StatusForbidden,
		)
		g.It(name, func() {
			c := tCase
			c.withBasicAuth = false

			status, _ := doRequest(c)
			g.Assert(status).Equal(StatusForbidden)
		})
	}

	if tCase.withRequestBody {
		name = fmt.Sprintf("%s %s: should return %d without request body",
			tCase.method, tCase.path, StatusUnprocessableEntity,
		)
		g.It(name, func() {
			c := tCase
			c.withRequestBody = false

			status, _ := doRequest(c)
			g.Assert(status).Equal(StatusUnprocessableEntity)
		})
	}

	if tCase.withRequestCookie {
		name = fmt.Sprintf("%s %s: should return %d without request cookie",
			tCase.method, tCase.path, StatusBadRequest,
		)
		g.It(name, func() {
			c := tCase
			c.withRequestCookie = false

			status, _ := doRequest(c)
			g.Assert(status).Equal(StatusBadRequest)
		})
	}

	name = fmt.Sprintf("%s %s: should return %d",
		MethodOptions, tCase.path, StatusOK,
	)
	g.It(name, func() {
		c := tCase
		c.method = MethodOptions

		status, _ := doRequest(c)
		g.Assert(status).Equal(StatusOK)
	})
}

func doRequest(tCase testCase) (int, []byte) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(tCase.basePath + tCase.path)
	req.Header.SetMethod(tCase.method)

	if tCase.withAuth {
		req.Header.Set(
			HeaderXUserID, tests.TestUserID,
		)
	}

	if tCase.withBasicAuth {
		req.Header.Set(
			HeaderAuthorization,
			makeBasicAuth(tests.TestClientID, tests.TestClientSecret),
		)
	}

	if tCase.withRequestBody {
		v := tests.TestingJSONType{
			SomeField: "some_value",
		}
		b, err := easyjson.Marshal(v)
		if err != nil {
			panic(err)
		}
		req.SetBody(b)
	}

	if tCase.withRequestCookie {
		req.Header.SetCookie(tests.TestCookie, "some-cookie-value")
	}

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		panic(err)
	}

	return resp.StatusCode(), resp.Body()
}

func makeBasicAuth(clientID, clientSecret string) string {
	auth := clientID + ":" + clientSecret
	base64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + base64Auth
}
