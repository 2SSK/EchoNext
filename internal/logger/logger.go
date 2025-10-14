package logger

import (
	"io"
	"os"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

// NewLogger creates a logger based on environment
func NewLogger(env string) zerolog.Logger {
	var logLevel zerolog.Level
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	switch level {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	default:
		logLevel = zerolog.InfoLevel
	}

	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	var writer io.Writer
	if env == "production" {
		// JSON writer for production
		writer = os.Stdout
	} else {
		// Console writer for development
		writer = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	}

	logger := zerolog.New(writer).
		Level(logLevel).
		With().
		Timestamp().
		Str("service", "EchoNext").
		Str("environment", env).
		Logger()

	// Include stack traces for errors in development
	if env != "production" {
		logger = logger.With().Stack().Logger()
	}

	return logger
}

// NewPgxLogger creates a database logger
func NewPgxLogger(level zerolog.Level) zerolog.Logger {
	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
	}

	return zerolog.New(writer).
		Level(level).
		With().
		Timestamp().
		Str("component", "database").
		Logger()
}

// GetPgxTraceLogLevel converts zerolog level to pgx tracelog level
func GetPgxTraceLogLevel(level zerolog.Level) tracelog.LogLevel {
	switch level {
	case zerolog.DebugLevel:
		return tracelog.LogLevelDebug
	case zerolog.InfoLevel:
		return tracelog.LogLevelInfo
	case zerolog.WarnLevel:
		return tracelog.LogLevelWarn
	case zerolog.ErrorLevel:
		return tracelog.LogLevelError
	default:
		return tracelog.LogLevelNone
	}
}
