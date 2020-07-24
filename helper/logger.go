package helper

import (
	log "github.com/sirupsen/logrus"
)

const (
	TOPIC = "user-service-log"
	LogTag = "user-service"
)

// LogContext function for logging the context of echo
func LogContext(c string, s string) *log.Entry {
	return log.WithFields(log.Fields{
		"topic":   TOPIC,
		"context": c,
		"scope":   s,
	})
}

func Log(level log.Level, message string, context string, scope string) {
	log.SetFormatter(&log.JSONFormatter{})
	entry := LogContext(context, scope)
	switch level {
	case log.DebugLevel:
		entry.Debug(message)
	case log.InfoLevel:
		entry.Info(message)
	case log.WarnLevel:
		entry.Warn(message)
	case log.ErrorLevel:
		entry.Error(message)
	case log.FatalLevel:
		entry.Fatal(message)
	case log.PanicLevel:
		entry.Panic(message)
	}
}
