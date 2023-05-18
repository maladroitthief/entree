package scene

import (
	"image/color"

	"github.com/maladroitthief/entree/domain/canvas"
)

type TitleScene struct {
	backgroundColor color.Color
}

func NewTitleScene(state *GameState) *TitleScene {
	return &TitleScene{
		backgroundColor: state.Theme.Black(),
	}
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

func (s *TitleScene) GetBackgroundColor() color.Color {
	return s.backgroundColor
}
