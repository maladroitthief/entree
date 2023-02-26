package adapter

import (
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/application/command"
	"github.com/sirupsen/logrus"
)

func NewSceneService(logger *logrus.Entry) *application.SceneService {
	return &application.SceneService{
		Commands: application.SceneCommands{
			Update: command.NewUpdateSceneHandler(logger),
		},
		Queries: application.SceneQueries{},
	}
}
