package logger

import (
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

func SetupLogger(logLevel string) *logrus.Logger {
	l := logrus.New()

	parsedLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Debug("failed to parse log level, log level will be set [info]")
	}

	l.SetLevel(parsedLevel)

	return l
}
