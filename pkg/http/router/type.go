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
	SecAuths   []SecAuth
	RouterDocs struct {
		Security     SecAuths
		Request      any
		Response     any
		IsDeprecated bool
		Tags         string
		Title        string
		Description  string
	}
)

func (sec SecAuth) isPublic() bool {
	return string(sec) == ""
}

func (sec SecAuths) isPublic() bool {
	return len(sec) == 0
}

func (p Path) EndsWith(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", p, path)
}
