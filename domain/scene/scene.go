package scene

import (
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/settings"
)

const (
	TransitionMaxCount = 20
)

type Scene interface {
	Update(*GameState) error
	GetEntities() []*canvas.Entity
}

type GameState struct {
	Log      logs.Logger
	SceneSvc SceneService
	InputSvc settings.InputService
}

type SceneService interface {
	GoTo(Scene) error
}
