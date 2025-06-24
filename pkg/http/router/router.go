package router

import (
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/handler"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router interface {
	Get(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Post(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Put(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Patch(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Delete(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Options(pattern string, h handler.HTTPHandlerFunc, opts *RouterOpts)

	Use(handlersNames ...middleware.HTTPMiddleware)

	Group(pattern string, fn func(r Router)) Router

	Scope(fn func(r Router))

	ServeDocs()

	Routes() *chi.Mux

	Path() string
}

type RouterImpl struct {
	mux  *chi.Mux
	path Path
}

func NewRouter(pattern string) Router {
	mux := chi.NewRouter()
	return &RouterImpl{
		mux:  mux,
		path: Path(pattern),
	}
}

func (r *RouterImpl) ServeDocs() {
	r.mux.Get("/docs/*", httpSwagger.WrapHandler)
}

func (r *RouterImpl) Routes() *chi.Mux {
	return r.mux
}

func (r *RouterImpl) Path() string {
	return string(r.path)
}
