package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type LogLevel zerolog.Level

const (
	DEBUG = LogLevel(zerolog.DebugLevel)
	INFO  = LogLevel(zerolog.InfoLevel)
	WARN  = LogLevel(zerolog.WarnLevel)
	ERROR = LogLevel(zerolog.ErrorLevel)
	FATAL = LogLevel(zerolog.FatalLevel)
)

type Logger struct {
	logger zerolog.Logger
}

func New(level LogLevel) *Logger {
	zerolog.TimeFieldFormat = time.RFC3339

	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		NoColor:    false,
	}

	logger := zerolog.New(output).
		Level(zerolog.Level(level)).
		With().
		Timestamp().
		Logger()

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
}
