package middleware

import (
	"net/http"
)

var middlewaresMap map[HTTPMiddleware]func(next http.Handler) http.Handler = map[HTTPMiddleware]func(next http.Handler) http.Handler{}

// RegisterMiddlewares registers all the middleware that is available for use in routes
/*
	eg: router.Use("basic_bearer_auth")
*/
func RegisterMiddlewares(handlers ...RegisterMiddlewaresArg) {
	for _, handl := range handlers {
		middlewaresMap[handl.Name] = handl.Handler
	}
}

func GetMiddleware(name HTTPMiddleware) func(next http.Handler) http.Handler {
	handler, found := middlewaresMap[name]
	if !found {
		return fallbackMiddleware
	}

	return handler
}

func fallbackMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
