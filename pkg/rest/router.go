package rest

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type Router struct {
	handler fasthttp.RequestHandler
}

type RouterParams struct {
	logs   *LogPack
	config RouterConfig
}

type RouterConfig struct {
	PathPrefix string
	Handlers   map[string][]RequestHandler
	Middleware []Middleware
}

type RequestHandler struct {
	Method     string
	Handler    Handler
	Middleware []Middleware
}

func NewRouter(params RouterParams) *Router {
	r := fasthttprouter.New()
	r.HandleOPTIONS = true

	cfg := params.config
	for path, handlers := range cfg.Handlers {
		for _, handler := range handlers {
			r.Handle(
				handler.Method,
				cfg.PathPrefix+path,
				makeHandler(
					handler.Handler,
					handler.Middleware,
					params.logs,
				),
			)
		}
	}

	return &Router{
		handler: addMiddleware(
			r.Handler, cfg.Middleware, params.logs,
		),
	}
}
