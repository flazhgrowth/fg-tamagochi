package router

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/apierrors"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/handler"
	fgmw "github.com/flazhgrowth/fg-tamagochi/pkg/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
)

type Router interface {
	Get(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

	Post(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

	Put(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

	Patch(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

	Delete(pattern string, h handler.HTTPHandlerFunc, docs ...RouterDocs)

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
	mux              *chi.Mux
	openapireflector *openapi3.Reflector
	prevPath         string
}

type OpenAPISpecInfo struct {
	Title   string
	Version string
	Desc    string
}

func NewRouter(specInfo OpenAPISpecInfo) Router {
	mux := chi.NewRouter()
	opsReflector := openapi3.NewReflector()
	opsReflector.Spec.SetTitle(specInfo.Title)
	opsReflector.Spec.SetVersion(specInfo.Version)
	opsReflector.Spec.SetDescription(specInfo.Desc)
	opsReflector.Spec.SetHTTPBearerTokenSecurity(string(SecurityBearerAuth), "JWT", "User Access Token")
	opsReflector.Spec.SetAPIKeySecurity(string(SecurityAPIKey), "X-API-Key", openapi.InHeader, "API Key")
	routerImpl := &RouterImpl{
		mux:              mux,
		openapireflector: opsReflector,
	}

	return routerImpl
}

func (r *RouterImpl) ServeDocs(pattern ...string) {
	ptrn := "/docs/*"
	if len(pattern) > 0 {
		ptrn = pattern[0]
	}

	// save specs
	specJsonByte, err := r.openapireflector.Spec.MarshalJSON()
	if err != nil {
		apierrors.LogPath("openapi spec marshaller").LogError(context.Background(), "error on marshalling openapi spec", err)
		return
	}
	if err = os.WriteFile("./docs/swagger.json", specJsonByte, 0644); err != nil {
		apierrors.LogPath("openapi spec writer failed").LogError(context.Background(), "error on writing openapi spec file", err)
	}

	r.mux.Get(ptrn, func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL:  "./docs/swagger.json",
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	})
}

func (r *RouterImpl) ServeProfiler(pattern ...string) {
	if len(pattern) > 0 {
		r.mux.Mount(pattern[0], middleware.Profiler())
		return
	}
	r.mux.Mount("/pprof/profiler", middleware.Profiler())
}

func (r *RouterImpl) Routes() *chi.Mux {
	return r.mux
}
