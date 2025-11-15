# FG-TAMAGOCHI
Honestly, a rather simple wrapper around the amazing [go-chi](https://github.com/go-chi/chi).

The main idea would be, how to get started, fast, using go-chi, with a few commands up on its sleeves.

## Requirements
1. Golang (version 1.21 or later)

## Install Conjurer CLI
Think of Conjurer as the main tools to do anything that related to a project, whether it is to initiate a project, make docs, and run migrations (well, currently the available commands on Conjurer are the three mentioned before).

Hence, to install Conjurer CLI, use the command below:
```
go install github.com/flazhgrowth/fg-tamagochi/cmd/conjurer@v0.1.0
```
This will install conjurer cli. You can type `conjurer conjure` in your terminal to see all the available commands on `conjurer`. Note that we use Tamagochi has not yet reach a stable version.

As of now, there are 4 available commands in `conjurer`:
1. `conjurer conjure init` => use this command to instantiate a new project inside your project directory. (eg: `conjurer conjure init --packagename="github.com/flazhgrowth/supertesterapp"`)
2. `conjurer conjure mocks` => use this command to generate mocks. Conjurer CLI uses [mockery](https://vektra.github.io/mockery/latest/installation/). When you use `init` command, it will try to detect if your system has `mockery` already. If you have not, it will then try to install `mockery` to you system. (eg: `conjurer conjure mocks`)
3. `conjurer conjure migration` => use this command to do anything related to migration. Conjurer CLI uses [go-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate). For go-migrate, Conjurer CLI will only tells you if you don't have the `migrate` CLI tool installed. Please refer to the github page on how to install the CLI tool.

NB: command `init` also accept argument `with-wire` (eg: `conjurer conjure init --packagename="github.com/flazhgrowth/supertesterapp" --with-wire`). This will also install [wire cli](github.com/google/wire/cmd/wire@latest). Wire can be use to ease up the process of wiring dependencies across your project (please refer to the github/docs on how to use wire if you prefer to use wire). For now, Tamagochi only provide `wire`. No other DI wiring tools planned for Tamagochi (like UberFx).

Under the hood, command `init` prepares the basic or the bare minimum to run an app. The provided argument `packagename` when you run `init` command is used to `go mod init` the project. Command `init` will also do `go mod tidy` due to the nature Tamagochi `init` command is templating, hence any used dependecies in Tamagochi, needs to be installed beforehands.

## Running the app
To run the app, simply run this in your terminal
```
go run main.go conjure serve
```

Notice that after main.go, we provide another command `conjure` followed by `serve`.

The base command is `conjure`, which practically does nothing. Like, `conjure` sounds cool tho, so we put it in there lol before anything else.

## The Basics
If you take a peek on the generated main.go file, you can see that the code looks like this
```
package main

import (
	"github.com/flazhgrowth/fg-tamagochi/app"
	"github.com/flazhgrowth/fg-tamagochi/cmd"
	"github.com/flazhgrowth/fg-tamagochi/cmd/serve"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/router"
)

func main() {
	cmd.Conjure(cmd.CmdArgs{
		ServeCmdArgs: serve.ServeCmdArgs{
			GetRoutesFn: func(app *app.App) router.Router {
				return router.NewRouter()
			},
		},
	})
}
```
This is the most barebone main.go to run the app. As you can see, `serve.ServeCmdArgs` is a struct that has other fields that you can filled in with things that you need.

Take a look at this sample below on a bit, bit, tiny more complex than the barebone one.
```
package main

import (
	"fmt"

	mw "github.com/superduper/lala-backend/internal/middleware"
	"github.com/superduper/lala-backend/internal/routers"
	"github.com/flazhgrowth/fg-tamagochi/app"
	"github.com/flazhgrowth/fg-tamagochi/cmd"
	"github.com/flazhgrowth/fg-tamagochi/cmd/serve"
	"github.com/flazhgrowth/fg-tamagochi/pkg/config"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/router"
)

func main() {
	if err := config.New(); err != nil {
		panic(fmt.Sprintf("error initializing configs: %s", err.Error()))
	}
	cmd.Conjure(cmd.CmdArgs{
		ServeCmdArgs: serve.ServeCmdArgs{
			GetRoutesFn: func(app *app.App) router.Router {
				return routers.GetRoutes(app)
			},
			CorsOpts:    mw.GetCorsOpts(),
			Middlewares: mw.GetMiddlewares(),
		},
	})
}
```
Here, you can note a few things:
1. Now we filled the Middlewares field with mw.GetMiddlewares(). This is how we register any middleware (by registering it to fields `Middlewares` on `ServeCmdArgs` struct) so later we can use it in the routes. Note that we made Tamagochi with middlewares mapping. So, GetMiddlewares function technically returns a `map[middleware.HTTPMiddleware]func(next http.Handler) http.Handler` type. Please refer to middlewares section to learn more about Tamagochi middlewares.
2. CorsOpts field basically filled out the cors setting for the app. Again, please refer to middlewares section to learn more about Tamagochi Cors. TL;DR, it uses chi cors middleware. But you can definitely modify it and add more validation or anything. Yes, you can add more validation to your liking.

### Routing and HTTP Handlers
The basics of Tamagochi routing and a details on HTTP Handler on Tamagochi. Refer to [Routing and HTTP Handler Section](./projectdocs/routing.md)

### Middleware
The basics of Tamagochi middleware, and how it differs from how chi does it. Refer to [Middleware](./projectdocs/middleware.md)

### CORS
The basic of Tamagochi CORS. Refer to [CORS](./projectdocs/cors.md)

### Cache
The basic of Tamagochi cache. Refer to [Cache](./projectdocs/cache.md)

### Config and Vault
Docs TBA

### Database and SQLator
Docs TBA

### HTTP Handler Transport (what you can do, and how we recommend you to use utilize this)
Docs TBA

### (I WON'T FORCE, BUT PLEASE CONSIDER LMAO) How Tamagochi structure its project
Docs TBA, sorry, lol

## Current Limitations
1. Currently, only postgres database driver supported. Man, I'm sorry lol. 
2. Configurable logger is not yet implemented. Yeah, its gonna be zerolog.
3. I'm thinking on implementing helper on APM. But still considering.
4. So many other things, but yeah, we forgot lol.