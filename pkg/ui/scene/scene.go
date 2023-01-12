package scene

import(
  "github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}
