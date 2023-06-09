package scene

import (
	"image/color"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
)

type TitleScene struct {
	backgroundColor color.Color
	camera          Camera
	width           int
	height          int
}

func NewTitleScene(state *GameState) *TitleScene {
	ts := &TitleScene{
		width:           800,
		height:          800,
		backgroundColor: state.Theme.Black(),
	}
	ts.camera = NewCamera(
		&focalPoint{
			position: physics.Vector{X: 0, Y: 0},
		},
		physics.Vector{X: float64(ts.width), Y: float64(ts.height)},
	)

	return ts
}

func (s *TitleScene) Update(state *GameState) error {
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

func (s *TitleScene) GetBackgroundColor() color.Color {
	return s.backgroundColor
}

func (s *TitleScene) GetCamera() Camera {
	return s.camera
}
