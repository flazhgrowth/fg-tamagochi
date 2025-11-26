package app

import (
	"fmt"
	"net/http"

	"github.com/flazhgrowth/fg-tamagochi/appconfig"
	"github.com/flazhgrowth/fg-tamagochi/pkg/cache"
	"github.com/flazhgrowth/fg-tamagochi/pkg/config"
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator"
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator/sqltx"
	"github.com/flazhgrowth/fg-tamagochi/pkg/featureflag"
	fgmw "github.com/flazhgrowth/fg-tamagochi/pkg/http/middleware"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/router"
	"github.com/flazhgrowth/fg-tamagochi/pkg/vault"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type App struct {
	appCfg    *appconfig.AppConfig
	sqlator   sqlator.SQLator
	txsqlator sqltx.SQLTx
	cache     cache.Cache
}

func New(appCfg *appconfig.AppConfig) *App {
	// initialize configs, vault, and featureflags
	if err := config.New(); err != nil {
		panic(fmt.Sprintf("error initializing configs: %s", err.Error()))
	}
	if err := vault.New(); err != nil {
		panic(fmt.Sprintf("error initializing vaults: %s", err.Error()))
	}
	if err := featureflag.New(); err != nil {
		panic(fmt.Sprintf("error initializing featureflags: %s", err.Error()))
	}

	// init db
	sqlator, txsqlator := sqlator.New(sqlator.SQLatorConfig{
		Driver:    vault.GetVault().Database.Driver,
		WriterDSN: vault.GetVault().Database.WriterDSN,
		ReaderDSN: vault.GetVault().Database.ReaderDSN,
	})

	corsBaseOpts := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token", "Accept-Encoding"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
	if appCfg.CorsOpt == nil {
		appCfg.CorsOpt = &fgmw.CorsOpt{
			Opts: corsBaseOpts,
		}
	}

	// default middlewares registered by fg-tamagochi
	mws := []fgmw.RegisterMiddlewaresArg{
		{
			Name:    fgmw.MIDDLEWARE_BASIC_BEARER_AUTH,
			Handler: fgmw.BasicBearerAuthMiddleware,
		},
		{
			Name:    fgmw.MIDDLEWARE_RECOVER_PANIC,
			Handler: middleware.Recoverer,
		},
		{
			Name:    fgmw.MIDDLEWARE_CORS,
			Handler: fgmw.Cors(*appCfg.CorsOpt),
		},
		{
			Name:    fgmw.MIDDLEWARE_REQUESTID,
			Handler: middleware.RequestID,
		},
		{
			Name:    fgmw.MIDDLEWARE_REALIP,
			Handler: middleware.RealIP,
		},
		{
			Name:    fgmw.MIDDLEWARE_LOGGER,
			Handler: middleware.Logger,
		},
		{
			Name:    fgmw.MIDDLEWARE_BASIC_API_KEY,
			Handler: fgmw.BasicAPIKeyMiddleware("X-API-Key"),
		},
	}
	// register additional middlewares defined by user
	if len(appCfg.Middlewares) > 0 {
		for key, mw := range appCfg.Middlewares {
			mws = append(mws, fgmw.RegisterMiddlewaresArg{
				Name:    key,
				Handler: mw,
			})
		}
	}
	fgmw.RegisterMiddlewares(mws...)

	var cacheClient cache.Cache
	if appCfg.UseCache {
		cacheClient = cache.New(cache.CacheConfig{})
	}

	return &App{
		appCfg:    appCfg,
		sqlator:   sqlator,
		txsqlator: txsqlator,
		cache:     cacheClient,
	}
}

func (app *App) Cfg() *appconfig.AppConfig {
	return app.appCfg
}

func (app *App) SetRouter(rtr router.Router) *Server {
	fgmw.PrintRoutes(rtr.Routes())

	return &Server{
		appCfg:       app.appCfg,
		sqlator:      app.sqlator,
		serverRouter: rtr,
	}
}
