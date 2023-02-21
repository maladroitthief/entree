package ports

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/app"
	"github.com/maladroitthief/entree/app/command"
)

type EbitenGame struct {
	app app.Application
}

func NewEbitenGame(app app.Application) EbitenGame {
	return EbitenGame{app}
}

func (e EbitenGame) Update() error {
	// Grab cursor coordinates
	cursorX, cursorY := ebiten.CursorPosition()

	// Grab current keyboard inputs
	inputs := []string{}
	cmd := command.Update{
		CursorX: cursorX,
		CursorY: cursorY,
		Inputs:  inputs,
	}

	return e.app.Commands.Update.Handle(cmd)
}

func (e EbitenGame) Draw(screen *ebiten.Image) {
	return
}

func (e EbitenGame) Layout(width, height int) (screenWidth, screenHeight int) {
	return 320, 240
}
