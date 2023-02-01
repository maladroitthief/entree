package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TitleScene struct {
	count int
}

func (s *TitleScene) Update(state *GameState) error {
	if state.Input.IsAnyAction() {
		state.SceneManager.GoTo(NewGameScene())
	}

	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) {
	ebitenutil.DebugPrint(r, "Title Scene")
}
