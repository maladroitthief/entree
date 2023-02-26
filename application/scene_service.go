package application

import (
	"github.com/maladroitthief/entree/application/command"
)

type SceneService struct {
	Commands SceneCommands
	Queries  SceneQueries
}

type SceneCommands struct {
	Update    command.UpdateSceneHandler
	GoToScene command.GoToSceneHandler
}

type SceneQueries struct {
}
