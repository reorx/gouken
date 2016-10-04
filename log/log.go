package log

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

var logger = logrus.New()

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields

// use ~ as prefix to make the kv show at the last sequence
const fileKey = "~file"

var logFilename bool

// Setup logging config
func Setup(lvl string, lfn bool) {
	lvlmap := map[string]logrus.Level{
		"PANIC": logrus.PanicLevel,
		"FATAL": logrus.FatalLevel,
		"ERROR": logrus.ErrorLevel,
		"WARN":  logrus.WarnLevel,
		"INFO":  logrus.InfoLevel,
		"DEBUG": logrus.DebugLevel,
	}
	lvlv, ok := lvlmap[lvl]
	if !ok {
		log.Fatalf("log level %v not exist\n", lvl)
	}
	SetLogLevel(lvlv)

	// set logFilename
	logFilename = lfn

	SetLogFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		// TimestampFormat: "Z07T2006-01-02 15:04:05",  // with timezone info
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// SetLogLevel ...
func SetLogLevel(level logrus.Level) {
	log.Printf("set log level = %v\n", level)
	logger.Level = level
}

// GetLogLevel ...
func GetLogLevel() logrus.Level {
	return logger.Level
}

// SetLogFormatter ...
func SetLogFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		fillentry(entry)
		entry.Debug(args)
	}
}

// DebugKV logs a message with fields at level Debug on the standard logger.
func DebugKV(l interface{}, f Fields) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields(f))
		fillentry(entry)
		entry.Debug(l)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		fillentry(entry)
		entry.Info(args...)
	}
}

// InfoKV Debug logs a message with fields at level Debug on the standard logger.
func InfoKV(l interface{}, f Fields) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields(f))
		fillentry(entry)
		entry.Info(l)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		fillentry(entry)
		entry.Warn(args...)
	}
}

// WarnKV Debug logs a message with fields at level Debug on the standard logger.
func WarnKV(l interface{}, f Fields) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields(f))
		fillentry(entry)
		entry.Warn(l)
	}
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		fillentry(entry)
		entry.Error(args...)
	}
}

// ErrorKV Debug logs a message with fields at level Debug on the standard logger.
func ErrorKV(l interface{}, f Fields) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields(f))
		fillentry(entry)
		entry.Error(l)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{})
		fillentry(entry)
		entry.Fatal(args...)
	}
}

// FatalKV Debug logs a message with fields at level Debug on the standard logger.
func FatalKV(l interface{}, f Fields) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields(f))
		fillentry(entry)
		entry.Fatal(l)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields{})
		fillentry(entry)
		entry.Panic(args...)
	}
}

// PanicKV Debug logs a message with fields at level Debug on the standard logger.
func PanicKV(l interface{}, f Fields) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields(f))
		fillentry(entry)
		entry.Panic(l)
	}
}

func fillentry(e *logrus.Entry) {
	if logFilename {
		e.Data[fileKey] = fileInfo(2)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
