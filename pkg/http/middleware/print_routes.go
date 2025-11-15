package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/flazhgrowth/fg-tamagochi/pkg/config"
	"github.com/go-chi/chi/v5"
)

func PrintRoutes(r chi.Router) {
	// for debugging purpose. No need to walk the routes on production
	if config.GetConfig().IsEnvProduction() {
		return
	}

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.Contains(route, "pprof") {
			return nil
		}
		val := fmt.Sprintf("[%s] %s", method, route)
		fmt.Println(val)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Error walking routes: %s\n", err.Error())
	}
}
