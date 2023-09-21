package logger

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

//nolint:gochecknoinits
func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}

func GetLogger() *logrus.Logger {
	return logger
}
