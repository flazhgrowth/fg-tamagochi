package router

import (
	"net/http"

	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/handler"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/middleware"
	"github.com/go-chi/chi/v5"
)

func (r *RouterImpl) Get(path string, h handler.HTTPHandlerFunc, opts *RouterOpts) {
	r.mux.Get(path, handle(h))
}

func (r *RouterImpl) Post(path string, h handler.HTTPHandlerFunc, opts *RouterOpts) {
	r.mux.Post(path, handle(h))
}

func (r *RouterImpl) Put(path string, h handler.HTTPHandlerFunc, opts *RouterOpts) {
	r.mux.Put(path, handle(h))
}

func (r *RouterImpl) Patch(path string, h handler.HTTPHandlerFunc, opts *RouterOpts) {
	r.mux.Patch(path, handle(h))
}

func (r *RouterImpl) Delete(path string, h handler.HTTPHandlerFunc, opts *RouterOpts) {
	r.mux.Delete(path, handle(h))
}

func (r *RouterImpl) Options(path string, h handler.HTTPHandlerFunc, opts *RouterOpts) {
	r.mux.Options(path, handle(h))
}

func (r *RouterImpl) Use(handlersNames ...middleware.HTTPMiddleware) {
	handlers := []func(http.Handler) http.Handler{}
	for _, handlerName := range handlersNames {
		handlers = append(handlers, middleware.GetMiddleware(handlerName))
	}
	r.mux.Use(handlers...)
}

func (r *RouterImpl) Group(path string, fn func(r Router)) Router {
	subrouter := NewRouter("")
	if fn == nil {
		panic("cannot group routes on nil fn")
	}
	fn(subrouter)
	r.mux.Mount(path, subrouter.Routes())

	return subrouter
}

func (r *RouterImpl) Scope(fn func(r Router)) {
	r.mux.Group(func(chiR chi.Router) {
		mux := &RouterImpl{
			mux: chiR.(*chi.Mux),
		}
		fn(mux)
	})
}
