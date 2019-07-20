package rest

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/pprofhandler"
)

type Handler func(ctx *Context)

func makeHandler(
	handler Handler,
	middleware []Middleware,
	logs *LogPack,
) fasthttp.RequestHandler {
	for _, mw := range middleware {
		handler = mw(handler)
	}

	return func(ctx *fasthttp.RequestCtx) {
		rCtx := Context{
			ctx:  ctx,
			logs: *logs,
		}
		handler(&rCtx)
	}
}

func addMiddleware(
	backendHandler fasthttp.RequestHandler,
	middleware []Middleware,
	logs *LogPack,
) fasthttp.RequestHandler {
	h := func(ctx *Context) {
		backendHandler(ctx.ctx)
	}

	for _, mw := range middleware {
		h = mw(h)
	}

	return func(ctx *fasthttp.RequestCtx) {
		rCtx := Context{
			ctx:  ctx,
			logs: *logs,
		}
		h(&rCtx)
	}
}

func PProfHandler() Handler {
	return func(ctx *Context) {
		pprofhandler.PprofHandler(ctx.ctx)
	}
}
