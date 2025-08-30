package router

import (
	"net/http"

	"github.com/flazhgrowth/fg-tamagochi/pkg/http/handler"
	fgmw "github.com/flazhgrowth/fg-tamagochi/pkg/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router interface {
	Get(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Post(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Put(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Patch(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Delete(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Options(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Use(handlersNames ...fgmw.HTTPMiddleware)

	Group(pattern string, fn func(r Router)) Router

	Scope(fn func(r Router))

	Mount(pattern string, fn http.Handler)

	ServeDocs()

	ServeProfiler(pattern ...string)

	Routes() *chi.Mux
}

type RouterImpl struct {
	mux *chi.Mux
}

func NewRouter() Router {
	mux := chi.NewRouter()
	return &RouterImpl{
		mux: mux,
	}
}

func (r *RouterImpl) ServeDocs() {
	r.mux.Get("/docs/*", httpSwagger.WrapHandler)
}

func (r *RouterImpl) ServeProfiler(pattern ...string) {
	if len(pattern) > 0 {
		r.mux.Mount(pattern[0], middleware.Profiler())
		return
	}
	r.mux.Mount("/healthcheck/profiler", middleware.Profiler())
}

func (r *RouterImpl) Routes() *chi.Mux {
	return r.mux
}
