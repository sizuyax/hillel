package logger

import (
	"github.com/sirupsen/logrus"
	"hillel/models"
)

var Logger *logrus.Logger

func NewLogger(cfg models.Config) {
	l := logrus.New()

	parsedLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.Error("failed to parse log level, log level will be set [info]")
		parsedLevel = logrus.InfoLevel
	}

	l.SetLevel(parsedLevel)

	Logger = l
}
