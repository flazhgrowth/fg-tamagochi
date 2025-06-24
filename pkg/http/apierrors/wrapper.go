package apierrors

import "fmt"

type LogPath string

func (lw LogPath) With(postfix string) LogPath {
	return LogPath(fmt.Sprintf("%s.%s", lw, postfix))
}

func (lw LogPath) WrapError(postfix string) error {
	return fmt.Errorf("%s.%s", lw, postfix)
}
