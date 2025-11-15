package apierrors

import (
	"context"
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
func (lw LogPath) LogError(ctx context.Context, msg string, err error) {
	log.Error().
		Ctx(ctx).
		Err(err).
		Str("path", string(lw)).
		Msg(msg)
}
