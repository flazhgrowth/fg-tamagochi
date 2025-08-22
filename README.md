# FG-TAMAGOCHI
Honestly, a rather simple wrapper around the amazing [go-chi](https://github.com/go-chi/chi).

The main idea would be, how to get started, fast, using go-chi, with a few commands up on its sleeves.

## Requirements
1. Golang (version 1.21 or later)

## Install Conjurer CLI
Think of Conjurer as the main tools to do anything that related to a project, whether it is to initiate a project, make docs, and run migrations (well, currently the available commands on Conjurer are the three mentioned before).

Hence, to install Conjurer CLI, use the command below:
```
go get github.com/flazhgrowth/fg-tamagochi/cmd/conjurer@ecb138ea7a2da5cb7214958190bf753466aa2021
```
This will install conjurer cli. You can type `conjurer conjure` in your terminal to see all the available commands on `conjurer`. Note that I use direct commit hash, instead of version. Well, Tamagochi has not yet reach a stable version. Hence, the commit, instead of version.

As of now, there are 3 available commands in `conjurer`:
1. `conjurer conjure init` => use this command to instantiate a new project inside your project directory. (eg: `conjurer conjure init --packagename="github.com/flazhgrowth/supertesterapp")
2. `conjurer conjure docs` => use this command if you want to generate swagger. Conjurer CLI uses [swaggo](https://github.com/swaggo/swag). Swaggo has its CLI tools called `swag`. When you use `init` command, it will try to detect if your system has `swag` already. If you have not, it will try to install `swag` to your system. (eg: `conjurer conjure docs`)
3. `conjurer conjure migration` => use this command to do anything related to migration. Conjurer CLI uses [go-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate). For go-migrate, Conjurer CLI will only tells you if you don't have the `migrate` CLI tool installed. Please refer to the github page on how to install the CLI tool.

NB: command `init` also accept argument `with-wire` (eg: `conjurer conjure init --packagename="github.com/flazhgrowth/supertesterapp" --with-wire`). This will also install [wire cli](github.com/google/wire/cmd/wire@latest). Wire can be use to ease up the process of wiring dependencies across your project (please refer to the github/docs om how to use wire if you prefer to use wire). For now, Tamagochi only provide `wire`. No other DI wiring tools planned for Tamagochi (like UberFx). I'm sorry, lol.

Under the hood, command `init` prepares the basic or the bare minimum to run an app. The provided argument `packagename` when you run `init` command is used to `go mod init` the project. Command `init` will also do `go mod tidy` due to the nature of templating is not directly usable, because of Tamagochi in general also use several packages (like chi, etc)

## Running the app
To run the app, simply run this in your terminal
```
go run main.go conjure serve
```

Notice that after main.go, we provide another command `conjure` followed by `serve`.

The base command is `conjure`, which practically does nothing. I mean, `conjure` sounds cool tho, so I put it in there lol before anything else.

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
				return router.NewRouter("")
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
1. Now we filled the Middlewares field with mw.GetMiddlewares(). This is how we register any middleware so later we can use it in the routes. Note that I made Tamagochi with middlewares mapping. So, GetMiddlewares function technically returns a `map[middleware.HTTPMiddleware]func(next http.Handler) http.Handler` type. Please refer to middlewares section to learn more about Tamagochi middlewares.
2. CorsOpts field basically filled out the cors setting for the app. Again, please refer to middlewares section to learn more about Tamagochi Cors. TL;DR, it uses chi cors middleware. But you can definitely modify it and add more validation or anything. Yes, you can add more validation to your liking.

### Routing
Docs TBA

### Middleware
Docs TBA

### CORS
Docs TBA

### Notification (Email)
Docs TBA

### Cache
Docs TBA

### Config and Vault
Docs TBA

### Database and SQLator
Docs TBA

### HTTP Handler Transport (what you can do, and how I recommend you to use utilize this)
Docs TBA

### (I WON'T FORCE, BUT PLEASE CONSIDER LMAO) How Tamagochi structure its project
Docs TBA, sorry, lol

## Current Limitations
1. Currently, only postgres database driver supported. Man, I'm sorry lol. 
2. Docs still use method annotation like how swaggo mentioned this. Still trying to figure out how I can simplify this using router opts (will probably rename this later to something that is more intent clear on docs).
3. Configurable logger is not yet implemented. Yeah, its gonna be zerolog.
4. I'm thinking on implementing helper on APM. But still considering.
5. So many other things, but yeah, I forgot lol.