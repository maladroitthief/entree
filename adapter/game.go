package adapter

import (
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/application/command"
	"github.com/maladroitthief/entree/application/query"
	"github.com/sirupsen/logrus"
)

func NewGameService() application.Game {
	// Service Setup
	// TODO Scene Service
	sceneApp := NewSceneService()

	return newGameService(sceneApp)
}

func newGameService(scene *application.Scene) application.Game {
	// Uniform logging setup
	logger := logrus.NewEntry(logrus.StandardLogger())

	return application.Game{
		Commands: application.GameCommands{
			Update: command.NewUpdateGameHandler(logger, scene),
		},
		Queries: application.GameQueries{
			WindowSettings: query.NewWindowSettingsHandler(logger),
		},
	}
}
