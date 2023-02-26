package adapter

import (
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/application/command"
	"github.com/maladroitthief/entree/domain/scene"
	"github.com/sirupsen/logrus"
)

func NewSceneService(logger *logrus.Entry, repo scene.Repository) *application.SceneService {
	return &application.SceneService{
		Commands: application.SceneCommands{
			Update: command.NewUpdateSceneHandler(logger, repo),
		},
		Queries: application.SceneQueries{},
	}
}
