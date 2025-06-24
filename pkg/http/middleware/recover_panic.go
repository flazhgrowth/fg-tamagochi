package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func RecoverPanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				trace := fmt.Sprintf("%s\n%s", fmt.Errorf("%s", err).Error(), debug.Stack())
				fmt.Println(trace) // TODO: log, instead of io prints
				w.Header().Set("Content-Type", "application/json")

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
