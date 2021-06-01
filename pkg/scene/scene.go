package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update(state *gameState) error
	Draw(screen *ebiten.Image)
}
