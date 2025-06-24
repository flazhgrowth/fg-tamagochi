package app

import (
	"github.com/flazhgrowth/fg-tamagochi/appconfig"
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/router"
)

type Server struct {
	appCfg       *appconfig.AppConfig
	sqlator      sqlator.SQLator
	serverRouter router.Router
}
