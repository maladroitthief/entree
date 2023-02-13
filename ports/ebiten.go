package ports

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/app"
)

type EbitenGame struct {
	app app.Application
}

func NewEbitenGame(app app.Application) EbitenGame {
	return EbitenGame{app}
}

func (e EbitenGame) Update() error {
	return nil
}

func (e EbitenGame) Draw(screen *ebiten.Image) {
	return
}

func (e EbitenGame) Layout(width, height int) (screenWidth, screenHeight int) {
	return 320, 240
}
