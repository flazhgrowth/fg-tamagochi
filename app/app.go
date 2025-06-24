package app

import (
	"fmt"
	"net/http"

	"github.com/flazhgrowth/fg-tamagotchi/appconfig"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/db/sqlator"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/db/sqlator/sqltx"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/middleware"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/router"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/vault"
	"github.com/go-chi/cors"
)

type App struct {
	appCfg    *appconfig.AppConfig
	sqlator   sqlator.SQLator
	txsqlator sqltx.SQLTx
}

func New(appCfg *appconfig.AppConfig) *App {
	if err := vault.New(); err != nil {
		panic(fmt.Sprintf("error initializing vaults: %s", err.Error()))
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
		appCfg.CorsOpt = &middleware.CorsOpt{
			Opts: corsBaseOpts,
		}
	}

	mws := []middleware.RegisterMiddlewaresArg{
		{
			Name:    middleware.MIDDLEWARE_BASIC_BEARER_AUTH,
			Handler: middleware.BasicBearerAuthMiddleware,
		},
		{
			Name:    middleware.MIDDLEWARE_RECOVER_PANIC,
			Handler: middleware.RecoverPanicMiddleware,
		},
		{
			Name:    middleware.MIDDLEWARE_CORS,
			Handler: middleware.Cors(*appCfg.CorsOpt),
		},
	}
	// register middlewares
	if len(appCfg.Middlewares) > 0 {
		for key, mw := range appCfg.Middlewares {
			mws = append(mws, middleware.RegisterMiddlewaresArg{
				Name:    key,
				Handler: mw,
			})
		}
	}
	middleware.RegisterMiddlewares(mws...)

	return &App{
		appCfg:    appCfg,
		sqlator:   sqlator,
		txsqlator: txsqlator,
	}
}

func (app *App) Cfg() *appconfig.AppConfig {
	return app.appCfg
}

func (app *App) DefineRoutes(rtr router.Router) *Server {
	middleware.PrintRoutes(rtr.Routes())

	return &Server{
		appCfg:       app.appCfg,
		sqlator:      app.sqlator,
		serverRouter: rtr,
	}
}
