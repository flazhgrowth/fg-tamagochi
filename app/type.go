package app

import (
	"github.com/flazhgrowth/fg-tamagotchi/appconfig"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/db/sqlator"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/router"
)

type Server struct {
	appCfg       *appconfig.AppConfig
	sqlator      sqlator.SQLator
	serverRouter router.Router
}
