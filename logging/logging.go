package logging

import (
	"github.com/sirupsen/logrus"
)

func InitializeLogger(logLevel string) Logger {

	log := logrus.New()

	// Set the log level based on an environment variable
	if logLevel == "" {
		logLevel = "info" // Default to info level if not specified
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel) // Default to InfoLevel if parsing fails
		logrus.Warn("Invalid LOG_LEVEL provided. Defaulting to InfoLevel.")
	} else {
		logrus.SetLevel(level)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // Custom time format
		FullTimestamp:   true,                  // Enable full timestamp
	})

	return &logrusLogger{log}
}
