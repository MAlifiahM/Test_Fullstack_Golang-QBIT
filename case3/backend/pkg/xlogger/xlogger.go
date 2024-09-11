package xlogger

import (
	"case3/internal/config"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var (
	Logger *zerolog.Logger
)

func Setup(cfg config.Config) {
	if cfg.IsDevelopment {
		l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
		l.Level(zerolog.DebugLevel)
		Logger = &l
		return
	}

	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	Logger = &l
}
