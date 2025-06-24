package middleware

import (
	"net/http"

	"github.com/flazhgrowth/fg-tamagochi/pkg/http/request"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/response"
	"github.com/go-chi/cors"
)

type (
	CorsOpt struct {
		Opts              cors.Options
		ValidatorHandlers []FnCorsAdditionValidator
	}
	FnCorsAdditionValidator func(request.Request) error
)

func Cors(opts CorsOpt) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := request.New(r)
			resp := response.New(w)

			for _, handler := range opts.ValidatorHandlers {
				if err := handler(req); err != nil {
					resp.Respond(nil, err)
					return
				}
			}

			cors.Handler(opts.Opts)(next).ServeHTTP(w, r)
		})
	}
}
