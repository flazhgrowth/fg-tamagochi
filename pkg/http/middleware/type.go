package middleware

import "net/http"

type HTTPMiddleware string

const (
	MIDDLEWARE_BASIC_BEARER_AUTH HTTPMiddleware = "basic_bearer_auth"
	MIDDLEWARE_RECOVER_PANIC     HTTPMiddleware = "recover_panic"
	MIDDLEWARE_CORS              HTTPMiddleware = "cors"
)

type RegisterMiddlewaresArg struct {
	Name    HTTPMiddleware
	Handler func(next http.Handler) http.Handler
}
