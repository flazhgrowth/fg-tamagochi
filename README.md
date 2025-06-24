# FG-TAMAGOCHI
Honestly, a rather simple wrapper around the amazing [go-chi](https://github.com/go-chi/chi).

The main idea would be, how to get started, fast, using go-chi, with a few commands up on its sleeves.

## Install
```
go get github.com/flazhgrowth/fg-tamagochi
```

## Examples
To get started, make a main.go in root app project with a content like so:
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
The snippet above gives the general idea on how to run an fg-tamagotchi based go app.

## Running the app
To run the app, simply run this in your terminal
```
go run main.go conjure serve
```

Notice that after main.go, we provide another command `conjure` followed by `serve`.

The base command is `conjure`, which practically does nothing. I mean, `conjure` sounds cool tho, so I put it in there lol before anything else.

## Available commands
You can run 
```
go run main.go conjure -h
```
It will then list all the available commands right of the bats. Every subcommand arguments can also referred using the `-h` argument
