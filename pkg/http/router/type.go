package router

import (
	"fmt"
	"strings"

	"github.com/flazhgrowth/fg-tamagochi/pkg/http/handler"
)

var (
	handle = handler.HandleHTTPHandler
)

const (
	SecBearerAuth SecAuth = "bearer"
	SecAPIKey     SecAuth = "apikey"
)

type (
	Path       string
	SecAuth    string
	RouterOpts struct {
		Security SecAuth
	}
)

func (p Path) EndsWith(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", p, path)
}
