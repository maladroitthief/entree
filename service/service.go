package service

import (
	"github.com/maladroitthief/entree/app"
	"github.com/maladroitthief/entree/app/command"
	"github.com/maladroitthief/entree/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication() app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	return app.Application{
		Commands: app.Commands{
			Update: command.NewUpdateHandler(logger),
		},
		Queries: app.Queries{
			WindowSettings: query.NewWindowSettingsHandler(logger),
		},
	}
}
