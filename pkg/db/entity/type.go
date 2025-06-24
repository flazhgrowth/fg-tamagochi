package entity

import (
	"fmt"
	"strings"
)

type DirName string

func (dir DirName) EndWith(val string) DirName {
	if strings.HasPrefix(val, "/") {
		return DirName(fmt.Sprintf("%s%s", dir, val))
	}
	return DirName(fmt.Sprintf("%s/%s", dir, val))
}

func (dir DirName) Val() string {
	if strings.HasPrefix(string(dir), "/") {
		return fmt.Sprintf(".%s", dir)
	}
	if strings.HasPrefix(string(dir), "./") {
		return string(dir)
	}
	return fmt.Sprintf("./%s", dir)
}
