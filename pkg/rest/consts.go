package rest

import (
	"github.com/go-http-utils/headers"
	"github.com/valyala/fasthttp"
)

const (
	MethodGet     = fasthttp.MethodGet
	MethodHead    = fasthttp.MethodHead
	MethodPost    = fasthttp.MethodPost
	MethodPut     = fasthttp.MethodPut
	MethodPatch   = fasthttp.MethodPatch
	MethodDelete  = fasthttp.MethodDelete
	MethodConnect = fasthttp.MethodConnect
	MethodOptions = fasthttp.MethodOptions
	MethodTrace   = fasthttp.MethodTrace
)

const (
	StatusContinue           = fasthttp.StatusContinue
	StatusSwitchingProtocols = fasthttp.StatusSwitchingProtocols
	StatusProcessing         = fasthttp.StatusProcessing

	StatusOK                   = fasthttp.StatusOK
	StatusCreated              = fasthttp.StatusCreated
	StatusAccepted             = fasthttp.StatusAccepted
	StatusNonAuthoritativeInfo = fasthttp.StatusNonAuthoritativeInfo
	StatusNoContent            = fasthttp.StatusNoContent
	StatusResetContent         = fasthttp.StatusResetContent
	StatusPartialContent       = fasthttp.StatusPartialContent
	StatusMultiStatus          = fasthttp.StatusMultiStatus
	StatusAlreadyReported      = fasthttp.StatusAlreadyReported
	StatusIMUsed               = fasthttp.StatusIMUsed

	StatusMultipleChoices   = fasthttp.StatusMultipleChoices
	StatusMovedPermanently  = fasthttp.StatusMovedPermanently
	StatusFound             = fasthttp.StatusFound
	StatusSeeOther          = fasthttp.StatusSeeOther
	StatusNotModified       = fasthttp.StatusNotModified
	StatusUseProxy          = fasthttp.StatusUseProxy
	StatusTemporaryRedirect = fasthttp.StatusTemporaryRedirect
	StatusPermanentRedirect = fasthttp.StatusPermanentRedirect

	StatusBadRequest                   = fasthttp.StatusBadRequest
	StatusUnauthorized                 = fasthttp.StatusUnauthorized
	StatusPaymentRequired              = fasthttp.StatusPaymentRequired
	StatusForbidden                    = fasthttp.StatusForbidden
	StatusNotFound                     = fasthttp.StatusNotFound
	StatusMethodNotAllowed             = fasthttp.StatusMethodNotAllowed
	StatusNotAcceptable                = fasthttp.StatusNotAcceptable
	StatusProxyAuthRequired            = fasthttp.StatusProxyAuthRequired
	StatusRequestTimeout               = fasthttp.StatusRequestTimeout
	StatusConflict                     = fasthttp.StatusConflict
	StatusGone                         = fasthttp.StatusGone
	StatusLengthRequired               = fasthttp.StatusLengthRequired
	StatusPreconditionFailed           = fasthttp.StatusPreconditionFailed
	StatusRequestEntityTooLarge        = fasthttp.StatusRequestEntityTooLarge
	StatusRequestURITooLong            = fasthttp.StatusRequestURITooLong
	StatusUnsupportedMediaType         = fasthttp.StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable = fasthttp.StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed            = fasthttp.StatusExpectationFailed
	StatusTeapot                       = fasthttp.StatusTeapot
	StatusUnprocessableEntity          = fasthttp.StatusUnprocessableEntity
	StatusLocked                       = fasthttp.StatusLocked
	StatusFailedDependency             = fasthttp.StatusFailedDependency
	StatusUpgradeRequired              = fasthttp.StatusUpgradeRequired
	StatusPreconditionRequired         = fasthttp.StatusPreconditionRequired
	StatusTooManyRequests              = fasthttp.StatusTooManyRequests
	StatusRequestHeaderFieldsTooLarge  = fasthttp.StatusRequestHeaderFieldsTooLarge
	StatusUnavailableForLegalReasons   = fasthttp.StatusUnavailableForLegalReasons

	StatusInternalServerError           = fasthttp.StatusInternalServerError
	StatusNotImplemented                = fasthttp.StatusNotImplemented
	StatusBadGateway                    = fasthttp.StatusBadGateway
	StatusServiceUnavailable            = fasthttp.StatusServiceUnavailable
	StatusGatewayTimeout                = fasthttp.StatusGatewayTimeout
	StatusHTTPVersionNotSupported       = fasthttp.StatusHTTPVersionNotSupported
	StatusVariantAlsoNegotiates         = fasthttp.StatusVariantAlsoNegotiates
	StatusInsufficientStorage           = fasthttp.StatusInsufficientStorage
	StatusLoopDetected                  = fasthttp.StatusLoopDetected
	StatusNotExtended                   = fasthttp.StatusNotExtended
	StatusNetworkAuthenticationRequired = fasthttp.StatusNetworkAuthenticationRequired
)

const (
	HeaderAccept                        = headers.Accept
	HeaderAcceptCharset                 = headers.AcceptCharset
	HeaderAcceptEncoding                = headers.AcceptEncoding
	HeaderAcceptLanguage                = headers.AcceptLanguage
	HeaderAuthorization                 = headers.Authorization
	HeaderCacheControl                  = headers.CacheControl
	HeaderContentLength                 = headers.ContentLength
	HeaderContentMD5                    = headers.ContentMD5
	HeaderContentType                   = headers.ContentType
	HeaderDoNotTrack                    = headers.DoNotTrack
	HeaderIfMatch                       = headers.IfMatch
	HeaderIfModifiedSince               = headers.IfModifiedSince
	HeaderIfNoneMatch                   = headers.IfNoneMatch
	HeaderIfRange                       = headers.IfRange
	HeaderIfUnmodifiedSince             = headers.IfUnmodifiedSince
	HeaderMaxForwards                   = headers.MaxForwards
	HeaderProxyAuthorization            = headers.ProxyAuthorization
	HeaderPragma                        = headers.Pragma
	HeaderRange                         = headers.Range
	HeaderReferer                       = headers.Referer
	HeaderUserAgent                     = headers.UserAgent
	HeaderTE                            = headers.TE
	HeaderVia                           = headers.Via
	HeaderWarning                       = headers.Warning
	HeaderCookie                        = headers.Cookie
	HeaderOrigin                        = headers.Origin
	HeaderAcceptDatetime                = headers.AcceptDatetime
	HeaderXRequestedWith                = headers.XRequestedWith
	HeaderAccessControlAllowOrigin      = headers.AccessControlAllowOrigin
	HeaderAccessControlAllowMethods     = headers.AccessControlAllowMethods
	HeaderAccessControlAllowHeaders     = headers.AccessControlAllowHeaders
	HeaderAccessControlAllowCredentials = headers.AccessControlAllowCredentials
	HeaderAccessControlExposeHeaders    = headers.AccessControlExposeHeaders
	HeaderAccessControlMaxAge           = headers.AccessControlMaxAge
	HeaderAccessControlRequestMethod    = headers.AccessControlRequestMethod
	HeaderAccessControlRequestHeaders   = headers.AccessControlRequestHeaders
	HeaderAcceptPatch                   = headers.AcceptPatch
	HeaderAcceptRanges                  = headers.AcceptRanges
	HeaderAllow                         = headers.Allow
	HeaderContentEncoding               = headers.ContentEncoding
	HeaderContentLanguage               = headers.ContentLanguage
	HeaderContentLocation               = headers.ContentLocation
	HeaderContentDisposition            = headers.ContentDisposition
	HeaderContentRange                  = headers.ContentRange
	HeaderETag                          = headers.ETag
	HeaderExpires                       = headers.Expires
	HeaderLastModified                  = headers.LastModified
	HeaderLink                          = headers.Link
	HeaderLocation                      = headers.Location
	HeaderP3P                           = headers.P3P
	HeaderProxyAuthenticate             = headers.ProxyAuthenticate
	HeaderRefresh                       = headers.Refresh
	HeaderRetryAfter                    = headers.RetryAfter
	HeaderServer                        = headers.Server
	HeaderSetCookie                     = headers.SetCookie
	HeaderStrictTransportSecurity       = headers.StrictTransportSecurity
	HeaderTransferEncoding              = headers.TransferEncoding
	HeaderUpgrade                       = headers.Upgrade
	HeaderVary                          = headers.Vary
	HeaderWWWAuthenticate               = headers.WWWAuthenticate

	HeaderXFrameOptions          = headers.XFrameOptions
	HeaderXXSSProtection         = headers.XXSSProtection
	HeaderContentSecurityPolicy  = headers.ContentSecurityPolicy
	HeaderXContentSecurityPolicy = headers.XContentSecurityPolicy
	HeaderXWebKitCSP             = headers.XWebKitCSP
	HeaderXContentTypeOptions    = headers.XContentTypeOptions
	HeaderXPoweredBy             = headers.XPoweredBy
	HeaderXUACompatible          = headers.XUACompatible
	HeaderXForwardedProto        = headers.XForwardedProto
	HeaderXHTTPMethodOverride    = headers.XHTTPMethodOverride
	HeaderXForwardedFor          = headers.XForwardedFor
	HeaderXRealIP                = headers.XRealIP
	HeaderXCSRFToken             = headers.XCSRFToken
	HeaderXRatelimitLimit        = headers.XRatelimitLimit
	HeaderXRatelimitRemaining    = headers.XRatelimitRemaining
	HeaderXRatelimitReset        = headers.XRatelimitReset

	HeaderXUserID    = "X-User-ID"
	HeaderXClientID  = "X-Client-ID"
	HeaderXRequestID = "X-Request-ID"
)

const (
	MimeApplicationJSON = "application/json"

	MimeTextPlain      = "text/plain"
	MimeTextHTML       = "text/html"
	MimeTextCSS        = "text/css"
	MimeTextJavaScript = "text/javascript"

	MimeImageBMP   = "image/bmp"
	MimeImageGIF   = "image/gif"
	MimeImageXIcon = "image/x-icon"
	MimeImageJPEG  = "image/jpeg"
)
