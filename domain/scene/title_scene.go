package scene

import (
	"image/color"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
	"github.com/maladroitthief/entree/domain/settings"
)

type TitleScene struct {
	backgroundColor color.Color
	camera          Camera
	width           int
	height          int
	cellSize        int
}

func NewTitleScene(state *GameState) *TitleScene {
	ts := &TitleScene{
		width:           800,
		height:          800,
		cellSize:        16,
		backgroundColor: state.Theme.Black(),
	}
	ts.camera = NewCamera(
		&focalPoint{
      x: 0,
      y: 0,
		},
		physics.Vector{X: float64(ts.width), Y: float64(ts.height)},
	)

	return ts
}

func (s *TitleScene) Update(state *GameState) error {
	for _, input := range state.InputSvc.CurrentInputs() {
		switch input {
		case settings.Menu:
			return SceneTermination
		}
	}

	if state.InputSvc.IsAny() {
		return state.SceneSvc.GoTo(NewGameScene(state))
	}

	return nil
}

func (s *TitleScene) GetEntities() []canvas.Entity {
	return nil
}

func (s *TitleScene) GetCanvasSize() (width, height int) {
	return s.width, s.height
}

func (s *TitleScene) GetCanvasCellSize() int {
	return s.cellSize
}

func (s *TitleScene) GetBackgroundColor() color.Color {
	return s.backgroundColor
}

func (s *TitleScene) GetCamera() Camera {
	return s.camera
}
