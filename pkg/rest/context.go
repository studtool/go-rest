package rest

import (
	"net"

	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"

	"github.com/studtool/go-conv/pkg/conv"
	"github.com/studtool/go-errs/pkg/errs"
)

type Context struct {
	logs LogPack
	ctx  *fasthttp.RequestCtx
}

func (ctx *Context) RemoteIP() net.IP {
	return ctx.ctx.RemoteIP()
}

func (ctx *Context) Method() string {
	return conv.BytesToString(ctx.ctx.Method())
}

func (ctx *Context) Path() string {
	return conv.BytesToString(ctx.ctx.Request.RequestURI())
}

func (ctx *Context) PathParam(name string) string {
	return ctx.ctx.UserValue(name).(string)
}

func (ctx *Context) QueryParam(name string) string {
	return conv.BytesToString(ctx.ctx.QueryArgs().Peek(name))
}

func (ctx *Context) RequestBody() []byte {
	return ctx.ctx.PostBody()
}

func (ctx *Context) SetRequestBody(body []byte) {
	ctx.ctx.Request.SetBody(body)
}

func (ctx *Context) RequestHeader(name string) string {
	return conv.BytesToString(ctx.ctx.Request.Header.Peek(name))
}

func (ctx *Context) SetRequestHeader(name, value string) {
	ctx.ctx.Request.Header.Set(name, value)
}

func (ctx *Context) RequestCookie(name string) string {
	return conv.BytesToString(ctx.ctx.Request.Header.Cookie(name))
}

func (ctx *Context) ResponseStatus() int {
	return ctx.ctx.Response.StatusCode()
}

func (ctx *Context) SetResponseStatus(status int) {
	ctx.ctx.SetStatusCode(status)
}

func (ctx *Context) ResponseHeader(name string) string {
	return conv.BytesToString(ctx.ctx.Response.Header.Peek(name))
}

func (ctx *Context) SetResponseHeader(name string, value string) {
	ctx.ctx.Response.Header.Set(name, value)
}

func (ctx *Context) SetResponseCookie(cookie *Cookie) {
	var fc fasthttp.Cookie
	cookie.initCookie(&fc)
	ctx.ctx.Response.Header.SetCookie(&fc)
}

func (ctx *Context) ResponseBody() []byte {
	return ctx.ctx.Response.Body()
}

func (ctx *Context) SetResponseBody(body []byte) {
	ctx.ctx.Response.SetBodyRaw(body)
}

func (ctx *Context) SetResponseContentType(mime string) {
	ctx.SetResponseHeader(HeaderContentType, mime)
}

func (ctx *Context) ReadJSON(
	v easyjson.Unmarshaler, err *errs.Error,
) *errs.Error {
	if easyjson.Unmarshal(ctx.RequestBody(), v) != nil {
		return err
	}
	return nil
}

func (ctx *Context) WriteJSON(status int, v easyjson.Marshaler) {
	ctx.SetResponseStatus(status)
	ctx.SetResponseContentType(MimeApplicationJSON)

	data, _ := easyjson.Marshal(v)
	ctx.SetResponseBody(data)
}

func (ctx *Context) WriteOK() {
	ctx.SetResponseStatus(StatusOK)
}

func (ctx *Context) WriteUnauthorized() {
	ctx.SetResponseStatus(StatusUnauthorized)
}

func (ctx *Context) WriteNotImplemented() {
	ctx.SetResponseStatus(StatusNotImplemented)
}

func (ctx *Context) WriteInternalServerError() {
	ctx.SetResponseStatus(StatusInternalServerError)
}

func (ctx *Context) WriteOKJSON(v easyjson.Marshaler) {
	ctx.WriteJSON(StatusOK, v)
}

func (ctx *Context) WriteErrStatus(err *errs.Error) {
	status := selectErrStatus(err)
	ctx.SetResponseStatus(status)
}

func (ctx *Context) WriteErrJSON(err *errs.Error) {
	status := selectErrStatus(err)

	ctx.SetResponseStatus(status)
	ctx.SetResponseContentType(MimeApplicationJSON)

	ctx.SetResponseBody(err.JSON())
}

func selectErrStatus(err *errs.Error) int {
	switch err.Type {
	case errs.Internal:
		return StatusInternalServerError
	case errs.BadFormat:
		return StatusBadRequest
	case errs.InvalidFormat:
		return StatusUnprocessableEntity
	case errs.Conflict:
		return StatusConflict
	case errs.NotFound:
		return StatusNotFound
	case errs.NotAuthorized:
		return StatusUnauthorized
	case errs.PermissionDenied:
		return StatusForbidden
	case errs.NotImplemented:
		return StatusNotImplemented
	default:
		panic(err)
	}
}
