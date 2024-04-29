package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
)

func Initialize(logLevel string) {
	var zl zerolog.Logger
	switch logLevel {
	case "info":
		zl = zerolog.New(os.Stdout)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zl = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zl = zerolog.New(os.Stdout)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	zl = zl.With().Timestamp().Logger()
	logger = zl
	logger.Info().Msg("Logging Initialized")
}

func LogFatalError(err error, message string, messages ...interface{}) {
	if err == nil {
		logger.Error().Fields(messages).Msg(message)
	} else {
		logger.Error().Err(err).Fields(messages).Msg(message)
	}
	os.Exit(1)
}

func LogDebug(message string, messages ...interface{}) {
	logger.Debug().Fields(messages).Msg(message)
}

func LogInfo(message string, messages ...interface{}) {
	logger.Info().Fields(messages).Msg(message)
}

func LogError(err error, message string, messages ...interface{}) {
	if err == nil {
		logger.Error().Fields(messages).Msg(message)
	} else {
		logger.Error().Err(err).Fields(messages).Msg(message)
	}
}
