package log

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	// SetLogFormatter(&logrus.JSONFormatter{})
	// SetLogLevel(logrus.InfoLevel)

	lvl := GetLogLevel()
	log.Printf("log level is: %v", lvl)

	Debug("debug", "a")
	Info("info", "b")
	InfoKV("info", Fields{
		"user": "me",
		"id":   1,
	})
	Warn("warn", "c")
}
