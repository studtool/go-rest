package rest

import (
	"time"

	"github.com/valyala/fasthttp"
)

type Cookie struct {
	Key      string
	Value    string
	Expire   time.Time
	HTTPOnly bool
	Secure   bool
	Domain   string
	Path     string
}

func (c *Cookie) initCookie(fc *fasthttp.Cookie) {
	fc.SetKey(c.Key)
	fc.SetValue(c.Value)
	fc.SetExpire(c.Expire)
	fc.SetHTTPOnly(c.HTTPOnly)
	fc.SetSecure(c.Secure)
	fc.SetDomain(c.Domain)
	fc.SetPath(c.Path)
}
