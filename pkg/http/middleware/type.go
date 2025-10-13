package middleware

import "net/http"

type HTTPMiddleware string

const (
	MIDDLEWARE_BASIC_BEARER_AUTH HTTPMiddleware = "basic_bearer_auth"
	MIDDLEWARE_BASIC_API_KEY     HTTPMiddleware = "basic_api_key"
	MIDDLEWARE_RECOVER_PANIC     HTTPMiddleware = "recover_panic"
	MIDDLEWARE_CORS              HTTPMiddleware = "cors"
	MIDDLEWARE_REQUESTID         HTTPMiddleware = "request_id"
	MIDDLEWARE_REALIP            HTTPMiddleware = "real_ip"
	MIDDLEWARE_LOGGER            HTTPMiddleware = "mwlogger"
)

type RegisterMiddlewaresArg struct {
	Name    HTTPMiddleware
	Handler func(next http.Handler) http.Handler
}
