package application

import (
	"github.com/maladroitthief/entree/application/command"
	"github.com/maladroitthief/entree/entity/scene"
)

type Scene struct {
	current         scene.Scene
	next            scene.Scene
	transitionCount int

	Commands SceneCommands
	Queries  SceneQueries
}

type SceneCommands struct {
	Update command.UpdateSceneHandler
}

type SceneQueries struct {
}
