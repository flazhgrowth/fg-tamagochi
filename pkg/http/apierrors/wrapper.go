package apierrors

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type LogPath string

// With extends the LogPath value with given postfix
/*
	eg:
		lw := LogPath("Lala")
		lw.With("Lili") // lw value is "Lala.Lili"
*/
func (lw LogPath) With(postfix string) LogPath {
	return LogPath(fmt.Sprintf("%s.%s", lw, postfix))
}

// LogError logs error using zerolog log.Error.Msgf()
func (lw LogPath) LogError(err error) {
	lw = LogPath(fmt.Sprintf("%s:%s", lw, err.Error()))
	log.Error().Msgf(string(lw))
}
