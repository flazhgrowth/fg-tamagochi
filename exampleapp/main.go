package main

import (
	"context"

	"github.com/flazhgrowth/fg-tamagochi/app"
	"github.com/flazhgrowth/fg-tamagochi/cmd"
	"github.com/flazhgrowth/fg-tamagochi/cmd/serve"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/apierrors"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/middleware"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/request"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/response"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/router"
	"github.com/flazhgrowth/fg-tamagochi/pkg/notification"
	"github.com/flazhgrowth/fg-tamagochi/pkg/notification/notiftype"
)

func getRoutes(app *app.App) router.Router {
	notif := notification.New(notiftype.Config{
		Driver: notiftype.DriverBrevo,
		APIKey: "",
	})
	v1 := router.NewRouter()
	v1.Use(middleware.MIDDLEWARE_CORS)
	v1.Get("/test", func(r request.Request, w response.Response) {
		resp, err := notif.Send(context.Background(), notiftype.SendArgs{
			Type:        notiftype.NotificationEmail,
			TemplateID:  1,
			Subject:     "TESTING Verification OTP",
			SenderEmail: "flazhgrowth@gmail.com",
			SenderName:  "no-reply",
			To:          []string{"henggana7@gmail.com"},
			DataBinding: map[string]any{
				"otp_code": "lalaland",
			},
		})
		if err != nil {
			w.Respond(nil, err)
			return
		}
		w.Respond(resp, nil)
	}, nil)
	v1.Group("/resources", func(usergroup router.Router) {
		usergroup.Get("/", func(r request.Request, w response.Response) {
			w.Respond(nil, nil)
		}, nil)
	})

	return v1
}

func main() {
	cmd.Conjure(cmd.CmdArgs{
		ServeCmdArgs: serve.ServeCmdArgs{
			GetRoutesFn: func(app *app.App) router.Router {
				return getRoutes(app)
			},
			CorsOpts: &middleware.CorsOpt{
				ValidatorHandlers: []middleware.FnCorsAdditionValidator{
					middleware.FnCorsAdditionValidator(func(r request.Request) error {
						return apierrors.ErrorForbidden()
					}),
					middleware.FnCorsAdditionValidator(func(r request.Request) error {
						return nil
					}),
				},
			},
		},
	})
}
