package handler

import (
	"net/http"

	"github.com/flazhgrowth/fg-tamagochi/pkg/http/request"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/response"
)

type HTTPHandlerFunc func(r request.Request, w response.Response)

func HandleHTTPHandler(h HTTPHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wHandler := response.New(w)
		rHandler := request.New(r)
		h(rHandler, wHandler)
	}
}
