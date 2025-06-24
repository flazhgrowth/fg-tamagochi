package serve

import (
	"net/http"

	"github.com/flazhgrowth/fg-tamagochi/app"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/middleware"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/router"
)

type ServeCmdArgs struct {
	GetRoutesFn func(app *app.App) router.Router
	Middlewares map[middleware.HTTPMiddleware]func(next http.Handler) http.Handler
	CorsOpts    *middleware.CorsOpt
}
