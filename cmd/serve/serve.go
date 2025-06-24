package serve

import (
	"fmt"

	"github.com/flazhgrowth/fg-tamagochi/app"
	"github.com/flazhgrowth/fg-tamagochi/appconfig"
	"github.com/flazhgrowth/fg-tamagochi/pkg/config"
	"github.com/spf13/cobra"
)

func Command(cmdArgs ServeCmdArgs) *cobra.Command {
	commands := &cobra.Command{
		Use:   "serve",
		Short: "serve the app",
		Run: func(cmd *cobra.Command, args []string) {
			serve(cmdArgs)
		},
	}

	return commands
}

func serve(cmdArgs ServeCmdArgs) {
	if err := config.New(); err != nil {
		panic(fmt.Sprintf("error initializing configs: %s", err.Error()))
	}
	app := app.New(&appconfig.AppConfig{
		HTTPServer: appconfig.HTTPServerConfig{
			Timeout: appconfig.HTTPServerTimeoutConfig{
				WriteTimeout: appconfig.Timeout(config.GetConfig().GetInt64WithDefault("http.timeout.write", 10)),
				ReadTimeout:  appconfig.Timeout(config.GetConfig().GetInt64WithDefault("http.timeout.read", 10)),
				IdleTimeout:  appconfig.Timeout(config.GetConfig().GetInt64WithDefault("http.timeout.idle", 30)),
				Unit:         config.GetConfig().GetStringWithDefault("http.timeout.unit", "second"),
			},
			Server: config.GetConfig().GetStringWithDefault("http.server", ":8000"),
		},
		Middlewares: cmdArgs.Middlewares,
		CorsOpt:     cmdArgs.CorsOpts,
	})

	if err := app.DefineRoutes(cmdArgs.GetRoutesFn(app)).Run(); err != nil {
		panic(err)
	}
}
