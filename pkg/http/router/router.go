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
	Get(pattern string, h handler.HTTPHandlerFunc)

	Post(pattern string, h handler.HTTPHandlerFunc)

	Put(pattern string, h handler.HTTPHandlerFunc)

	Patch(pattern string, h handler.HTTPHandlerFunc)

	Delete(pattern string, h handler.HTTPHandlerFunc)

	Options(pattern string, h handler.HTTPHandlerFunc)

	Use(handlersNames ...fgmw.HTTPMiddleware)

	Group(pattern string, fn func(r Router)) Router

	Scope(fn func(r Router))

	Mount(pattern string, fn http.Handler)

	ServeDocs(pattern ...string)

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

func (r *RouterImpl) ServeDocs(pattern ...string) {
	if len(pattern) > 0 {
		r.mux.Get(pattern[0], httpSwagger.WrapHandler)
		return
	}
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
