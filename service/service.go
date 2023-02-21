package service

import (
	"github.com/maladroitthief/entree/app"
	"github.com/maladroitthief/entree/app/command"
	"github.com/sirupsen/logrus"
)

func NewApplication() app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	return app.Application{
		Commands: app.Commands{
			Update: command.NewUpdateHandler(logger),
		},
	}
}
