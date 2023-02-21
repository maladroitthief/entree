package decorator

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger *logrus.Entry
}

func (d commandLoggingDecorator[C]) Handle(cmd C) (err error) {
	handlerType := generateActionName(cmd)

	logger := d.logger.WithFields(logrus.Fields{
		"command":      handlerType,
		"command_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("executing command")
	defer func() {
		if err == nil {
			logger.Info("command executed successfully")
		} else {
			logger.WithError(err).Error("failed to execute command")
		}
	}()

	return d.base.Handle(cmd)
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger *logrus.Entry
}

func (d queryLoggingDecorator[Q, R]) Handle(q Q) (result R, err error) {
	logger := d.logger.WithFields(logrus.Fields{
		"query":      generateActionName(q),
		"query_body": fmt.Sprintf("%#v", q),
	})

	logger.Debug("executing query")
	defer func() {
		if err == nil {
			logger.Info("query executed successfully")
		} else {
			logger.WithError(err).Error("failed to execute query")
		}
	}()

	return d.base.Handle(q)
}
