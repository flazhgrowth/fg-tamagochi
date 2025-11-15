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
	SecurityBearerAuth SecAuth = "bearer"
	SecurityAPIKey     SecAuth = "apikey"
)

type (
	Path       string
	SecAuth    string
	RouterDocs struct {
		Security     SecAuth
		Request      any
		Response     any
		IsDeprecated bool
		Tags         string
		Summary      string
		Description  string
	}
)

func (sec SecAuth) IsPublic() bool {
	return string(sec) == ""
}

func (p Path) EndsWith(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", p, path)
}
