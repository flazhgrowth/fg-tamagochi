package notification

import (
	"context"

	"github.com/flazhgrowth/fg-tamagotchi/pkg/notification/brevo"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/notification/notiftype"
)

type Notification interface {
	Send(ctx context.Context, args notiftype.SendArgs) (resp *notiftype.SendResponse, err error)
}

func New(cfg notiftype.Config) Notification {
	var notif Notification
	if cfg.Driver.Is(notiftype.DriverBrevo) {
		notif = brevo.New(brevo.BrevoCfg{
			APIKey: cfg.APIKey,
		})
	}

	return notif
}
