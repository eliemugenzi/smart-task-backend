package utils

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger() *Logger {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	return &Logger{
		logger: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger(),
	}
}

func (this_ *Logger) Info(msg string) {
	this_.logger.Info().Msg(msg)
}
func (this_ *Logger) Error(msg string) {
	this_.logger.Error().Msg(msg)
}

func (this_ *Logger) Debug(msg string) {
	this_.logger.Debug().Msg(msg)
}

func (this_ *Logger) Warn(msg string) {
	this_.logger.Warn().Msg(msg)
}

func (this_ *Logger) Fatal(msg string) {
	this_.logger.Fatal().Msg(msg)
}

func (this_ *Logger) Panic(msg string) {
	this_.logger.Panic().Msg(msg)
}

func (this_ *Logger) FatalErr(err error) {
	this_.logger.Fatal().Err(err).Msg(err.Error())
}
