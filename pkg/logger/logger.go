package logger

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

func (lw LogPath) LogError(ctx context.Context, msg string, err error, logdata ...map[string]any) {
	e := log.Error().
		Ctx(ctx).
		Err(err).
		Str("path", string(lw))
	for _, logdatum := range logdata {
		for key, val := range logdatum {
			e.Any(key, val)
		}
	}

	e.Msg(msg)
}

func (lw LogPath) LogDebug(ctx context.Context, msg string, logdata ...map[string]any) {
	e := log.Debug().
		Ctx(ctx).
		Str("path", string(lw))
	for _, logdatum := range logdata {
		for key, val := range logdatum {
			e.Any(key, val)
		}
	}

	e.Msg(msg)
}

func (lw LogPath) LogInfo(ctx context.Context, msg string, logdata ...map[string]any) {
	e := log.Info().
		Ctx(ctx).
		Str("path", string(lw))
	for _, logdatum := range logdata {
		for key, val := range logdatum {
			e.Any(key, val)
		}
	}

	e.Msg(msg)
}

func (lw LogPath) LogFatal(ctx context.Context, msg string, err error, logdata ...map[string]any) {
	e := log.Fatal().
		Ctx(ctx).
		Str("path", string(lw))
	if err != nil {
		e.Err(err)
	}
	for _, logdatum := range logdata {
		for key, val := range logdatum {
			e.Any(key, val)
		}
	}
	e.Msg(msg)
}
