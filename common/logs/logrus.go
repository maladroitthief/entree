package logs

import "github.com/sirupsen/logrus"

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() Logger {
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})

	return &LogrusLogger{logger}
}

func (l *LogrusLogger) Info(message string, args interface{}) {
	log := l.logger.WithField("args", args)
  log.Info(message)
}

func (l *LogrusLogger) Error(methodName string, args interface{}, err error) {
	log := l.logger.WithField("args", args)
  log.WithError(err).Error(methodName + " failed")
}

func (l *LogrusLogger) Fatal(methodName string, args interface{}, err error) {
	log := l.logger.WithField("args", args)
  log.WithError(err).Fatal(methodName + " failed")
}
