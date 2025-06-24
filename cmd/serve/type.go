package serve

import (
	"net/http"

	"github.com/flazhgrowth/fg-tamagotchi/app"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/middleware"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/router"
)

type ServeCmdArgs struct {
	GetRoutesFn func(app *app.App) router.Router
	Middlewares map[middleware.HTTPMiddleware]func(next http.Handler) http.Handler
	CorsOpts    *middleware.CorsOpt
}
