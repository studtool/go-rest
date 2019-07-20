package rest

import (
	"encoding/base64"
	"strings"

	"github.com/studtool/go-conv/pkg/conv"
	"github.com/studtool/go-errs/pkg/errs"
)

func IsInformational(ctx *Context) bool {
	return ctx.ResponseStatus() >= 100 &&
		ctx.ResponseStatus() < 200
}

func IsSuccessful(ctx *Context) bool {
	return ctx.ResponseStatus() >= 200 &&
		ctx.ResponseStatus() < 300
}

func IsRedirection(ctx *Context) bool {
	return ctx.ResponseStatus() >= 300 &&
		ctx.ResponseStatus() < 400
}

func IsClientError(ctx *Context) bool {
	return ctx.ResponseStatus() >= 400 &&
		ctx.ResponseStatus() < 500
}

func IsServerError(ctx *Context) bool {
	return ctx.ResponseStatus() >= 500
}

type ClientCredentials struct {
	ID     string
	Secret string
}

func parseBasicAuth(
	ctx *Context,
	credentials *ClientCredentials,
	err *errs.Error,
) *errs.Error {
	auth := ctx.RequestHeader(HeaderAuthorization)
	if auth == "" {
		return err
	}

	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return err
	}

	c, e := base64.StdEncoding.
		DecodeString(auth[len(prefix):])
	if e != nil {
		return err
	}

	cs := conv.BytesToString(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return err
	}

	credentials.ID = cs[:s]
	credentials.Secret = cs[s+1:]

	return nil
}
