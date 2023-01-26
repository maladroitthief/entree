package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/ui/input"
)

type Scene interface {
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}

type GameState struct {
	SceneManager *SceneManager
	Input        *input.Input
}
