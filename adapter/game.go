package adapter

import (
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/application/command"
	"github.com/maladroitthief/entree/application/query"
	"github.com/sirupsen/logrus"
)

func NewGameService() application.GameService {
	// Uniform logging setup
	logger := logrus.NewEntry(logrus.StandardLogger())

	// Service Setup
	// TODO Scene Service
	sceneService := NewSceneService(logger)

	return newGameService(logger, sceneService)
}

func newGameService(
	logger *logrus.Entry,
	scene *application.SceneService,
) application.GameService {

	return application.GameService{
		Commands: application.GameCommands{
			Update: command.NewUpdateGameHandler(logger, scene),
		},
		Queries: application.GameQueries{
			WindowSettings: query.NewWindowSettingsHandler(logger),
		},
	}
}
