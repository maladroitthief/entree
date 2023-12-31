package logs

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() Logger {
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.999999999",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	return &LogrusLogger{logger}
}

func (l *LogrusLogger) SetLevel(level string) {
	switch level {
	case "trace", "Trace":
		l.logger.SetLevel(logrus.TraceLevel)
	case "debug", "Debug":
		l.logger.SetLevel(logrus.DebugLevel)
	case "info", "Info":
		l.logger.SetLevel(logrus.InfoLevel)
	case "warn", "Warn":
		l.logger.SetLevel(logrus.WarnLevel)
	case "error", "Error":
		l.logger.SetLevel(logrus.ErrorLevel)
	case "fatal", "Fatal":
		l.logger.SetLevel(logrus.FatalLevel)
	case "panic", "Panic":
		l.logger.SetLevel(logrus.PanicLevel)
	default:
		l.Fatal("Logrus", "SetLevel()", errors.New("Not a valid level: "+level))
	}
}

func (l *LogrusLogger) Debug(message string, context interface{}) {
	log := l.logger.WithField("context", context)
	log.Debug(message)
}

func (l *LogrusLogger) Info(message string, context interface{}) {
	log := l.logger.WithField("context", context)
	log.Info(message)
}

func (l *LogrusLogger) Error(methodName string, context interface{}, err error) {
	log := l.logger.WithField("context", context)
	log.WithError(err).Error(methodName + " failed")
}

func (l *LogrusLogger) Fatal(methodName string, context interface{}, err error) {
	log := l.logger.WithField("context", context)
	log.WithError(err).Fatal(methodName + " failed")
}
