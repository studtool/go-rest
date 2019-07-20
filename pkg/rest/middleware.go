package rest

import (
	"time"

	"github.com/studtool/go-errs/pkg/errs"
	"github.com/studtool/go-logs/pkg/logs"
)

type Middleware func(h Handler) Handler

func WithRecover(h Handler) Handler {
	return func(ctx *Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.logs.StackTraceLogger.Errorf("panic: ['%v']", r)
				ctx.WriteInternalServerError()
			}
		}()

		h(ctx)
	}
}

func WithAuth(h Handler) Handler {
	return func(ctx *Context) {
		if ctx.RequestHeader(HeaderXUserID) == "" {
			ctx.WriteUnauthorized()
			return
		}
		h(ctx)
	}
}

type BasicAuthMiddlewareParams struct {
	AuthErr      *errs.Error
	KnownClients map[string]string
}

func WithBasicAuth(params BasicAuthMiddlewareParams) Middleware {
	return func(h Handler) Handler {
		return func(ctx *Context) {
			var client ClientCredentials
			if parseBasicAuth(ctx, &client, params.AuthErr) != nil {
				ctx.WriteErrJSON(params.AuthErr)
				return
			}

			if client.Secret != params.KnownClients[client.ID] {
				ctx.WriteErrJSON(params.AuthErr)
				return
			}

			ctx.SetRequestHeader(HeaderXClientID, client.ID)
			h(ctx)
		}
	}
}

const (
	RequestTypeMetrics   = "metrics"
	RequestTypeProfiling = "profiling"
	RequestTypeInternal  = "internal"
	RequestTypePrivate   = "private"
	RequestTypeProtected = "protected"
	RequestTypePublic    = "public"
	RequestTypeTesting   = "testing"
)

const (
	ACLGroupCommon  = "common"
	ACLGroupTesting = "testing"
)

type LogsMiddlewareParams struct {
	RequestTypeDetector     func(*Context) string
	RequestACLGroupDetector func(*Context) string
}

func WithLogs(params LogsMiddlewareParams) Middleware {
	return func(h Handler) Handler {
		return func(ctx *Context) {
			startTime := time.Now()
			h(ctx)
			requestTime := time.Since(startTime)

			logFunc := ctx.logs.
				RequestLogger.Info

			if IsClientError(ctx) {
				logFunc = ctx.logs.
					RequestLogger.Warning
			} else if IsServerError(ctx) {
				logFunc = ctx.logs.
					RequestLogger.Error
			}

			p := logs.RequestParams{
				Method:      ctx.Method(),
				Path:        ctx.Path(),
				Status:      ctx.ResponseStatus(),
				Type:        params.RequestTypeDetector(ctx),
				UserID:      ctx.RequestHeader(HeaderXUserID),
				ClientID:    ctx.RequestHeader(HeaderXClientID),
				RequestIP:   ctx.RemoteIP().String(),
				XRealIP:     ctx.RequestHeader(HeaderXRealIP),
				ACLGroup:    params.RequestACLGroupDetector(ctx),
				UserAgent:   ctx.RequestHeader(HeaderUserAgent),
				RequestTime: requestTime,
			}

			logFunc(&p)
		}
	}
}
