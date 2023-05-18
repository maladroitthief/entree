package scene

import (
	"image/color"

	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/common/theme"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/settings"
)

const (
	TransitionMaxCount = 20
)

type Scene interface {
	Update(*GameState) error
	GetEntities() []canvas.Entity
	GetBackgroundColor() color.Color
}

type GameState struct {
	Log      logs.Logger
	SceneSvc SceneService
	InputSvc settings.InputService
	Theme    theme.Colors
}

type SceneService interface {
	GoTo(Scene) error
}
