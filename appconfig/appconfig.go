package appconfig

import (
	"net/http"
	"time"

	"github.com/flazhgrowth/fg-tamagotchi/pkg/http/middleware"
)

type (
	AppConfig struct {
		HTTPServer  HTTPServerConfig
		Middlewares map[middleware.HTTPMiddleware]func(next http.Handler) http.Handler
		CorsOpt     *middleware.CorsOpt
	}

	HTTPServerConfig struct {
		Timeout HTTPServerTimeoutConfig
		Server  string
	}
)

type (
	HTTPServerTimeoutConfig struct {
		WriteTimeout Timeout
		ReadTimeout  Timeout
		IdleTimeout  Timeout
		Unit         string
	}
)

type Timeout int64

func (t Timeout) getTimeoutUnitHandler() map[string]func() time.Duration {
	return map[string]func() time.Duration{
		"second": t.ValInSecond,
		"minute": t.ValInMinute,
		"hour":   t.ValInHour,
	}
}

func (t Timeout) Val(unit string) time.Duration {
	handler, ok := t.getTimeoutUnitHandler()[unit]
	if !ok {
		return t.ValInSecond()
	}
	return handler()
}

func (t Timeout) ValInSecond() time.Duration {
	return time.Second * time.Duration(t)
}

func (t Timeout) ValInMinute() time.Duration {
	return time.Minute * time.Duration(t)
}

func (t Timeout) ValInHour() time.Duration {
	return time.Hour * time.Duration(t)
}
