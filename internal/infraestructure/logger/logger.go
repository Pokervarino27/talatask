package logger

import (
	"encoding/json"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}

func Fatal(msg string, err error) {
	log.Fatal().Err(err).Msg(msg)
}

func Infof(msg string, v interface{}) {
	json, err := json.Marshal(v)
	if err != nil {
		Error("Error marshalling log data", err)
		return
	}
	log.Info().RawJSON("data", json).Msg(msg)
}
