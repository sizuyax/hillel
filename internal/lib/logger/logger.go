package logger

import (
	sloglogrus "github.com/samber/slog-logrus/v2"
	"github.com/sirupsen/logrus"
	"log/slog"
)

func SetupLogger(logLevel string) *slog.Logger {
	l := logrus.New()

	parsedLevelLogrus, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Error("failed to parse log level, log level will be set [info]")
		parsedLevelLogrus = logrus.InfoLevel
	}

	l.SetLevel(parsedLevelLogrus)

	logger := slog.New(sloglogrus.Option{Level: slog.LevelDebug, Logger: l}.NewLogrusHandler())

	return logger
}
