package ui

import (
	"context"
	"image/color"

	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
)

type TitleScene struct {
	ctx             context.Context
	camera          *Camera
	width           int
	height          int
	cellSize        int
	backgroundColor color.Color
}

func NewTitleScene(ctx context.Context, state *SceneState) *TitleScene {
	ts := &TitleScene{
		ctx:             ctx,
		width:           800,
		height:          800,
		cellSize:        32,
		backgroundColor: state.theme.Black(),
	}

	ts.camera = NewCamera(
		0,
		0,
		mosaic.Vector{X: float64(ts.width), Y: float64(ts.height)},
	)

	return ts
}

func (s *TitleScene) Update(state *SceneState) error {
	for _, input := range state.input.CurrentInputs() {
		switch input {
		case core.InputMenu:
			return Termination
		}
	}

	if state.input.IsAny() {
		return state.mgr.GoTo(NewGameScene(s.ctx, state))
	}

	return nil
}

func (s *TitleScene) GetWorld() *content.World {
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
