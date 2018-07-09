package hedrs

import "net/http"

// request constants ...
const (
	Accept                      = "Accept"
	AcceptEncoding              = "Accept-Encoding"
	AcceptVersion               = "Accept-Version"
	AccessControlRequestHeaders = "Access-Control-Request-Headers"
	AccessControlRequestMethod  = "Access-Control-Request-Method"
	APIVersion                  = "X-Api-Version"
	Authorization               = "Authorization"
	ContentLength               = "Content-Length"
	ContentMD5                  = "Content-MD5"
	ContentType                 = "Content-Type"
	CSRFToken                   = "X-CSRF-Token"
	Date                        = "Date"
	Origin                      = "Origin"
	RequestedWith               = "X-Requested-With"
)

// response constants ...
const (
	AccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	AccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	AccessControlAllowMethods     = "Access-Control-Allow-Methods"
	AccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	AccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	AccessControlMaxAge           = "Access-Control-Max-Age"
)

var (
	// AllMethods ...
	AllMethods = []string{
		http.MethodOptions,
		http.MethodGet,
		http.MethodPut,
		http.MethodHead,
		http.MethodPost,
		http.MethodDelete,
		http.MethodPatch,
		http.MethodTrace,
	}
)
