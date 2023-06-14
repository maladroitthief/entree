package scene

import (
	"errors"
	"image/color"

	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/common/theme"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/settings"
)

const (
	TransitionMaxCount = 20
)

var (
	SceneTermination = errors.New("scene exited normally")
)

type Scene interface {
	Update(*GameState) error
	GetCamera() Camera
	GetCanvasSize() (width, height int)
	GetCanvasCellSize() int
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
