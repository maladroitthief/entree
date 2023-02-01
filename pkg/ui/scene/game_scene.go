package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameScene struct {
}

func NewGameScene() *GameScene {
  gs := &GameScene{}

  return gs
}

func (s *GameScene) Update(state *GameState) error {
	return nil
}

func (s *GameScene) Draw(r *ebiten.Image) {
	ebitenutil.DebugPrint(r, "Game Scene")
}
