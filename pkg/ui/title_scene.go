package ui

import (
	"image/color"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

type TitleScene struct {
	log             logs.Logger
	camera          *Camera
	width           int
	height          int
	cellSize        int
	backgroundColor color.Color
}

func NewTitleScene(state *SceneState) *TitleScene {
	ts := &TitleScene{
		width:           800,
		height:          800,
		cellSize:        16,
		log:             state.log,
		backgroundColor: state.theme.Black(),
	}

	ts.camera = NewCamera(
		data.Vector{
			X: 0,
			Y: 0,
		},
		data.Vector{X: float64(ts.width), Y: float64(ts.height)},
	)

	return ts
}

func (s *TitleScene) Update(state *SceneState) error {
	for _, input := range state.input.CurrentInputs() {
		switch input {
		case core.Menu:
			return Termination
		}
	}

	if state.input.IsAny() {
		return state.mgr.GoTo(NewGameScene(state))
	}

	return nil
}

func (s *TitleScene) GetState() *core.ECS {
	return nil
}

func (s *TitleScene) Size() (width, height int) {
	return s.width, s.height
}

func (s *TitleScene) CellSize() int {
	return s.cellSize
}

func (s *TitleScene) BackgroundColor() color.Color {
	return s.backgroundColor
}

func (s *TitleScene) GetCamera() *Camera {
	return s.camera
}
