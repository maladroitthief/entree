package ports

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/app"
	"github.com/maladroitthief/entree/app/command"
	"github.com/maladroitthief/entree/app/query"
)

type EbitenGame struct {
	app app.Application

	width  int
	height int
	title  string
}

func NewEbitenGame(app app.Application) (*EbitenGame, error) {
	e := &EbitenGame{
		app:    app,
		width:  0,
		height: 0,
		title:  "",
	}

	err := e.WindowHandler()

	return e, err
}

func (e *EbitenGame) Update() (err error) {
	err = e.WindowHandler()
	if err != nil {
		return err
	}

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

func (e *EbitenGame) Draw(screen *ebiten.Image) {
	return
}

func (e *EbitenGame) Layout(width, height int) (screenWidth, screenHeight int) {
	return e.width, e.height
}

func (e *EbitenGame) WindowHandler() error {
	windowSettings, err := e.app.Queries.WindowSettings.Handle(query.WindowSettings{})
	if err != nil {
		return err
	}

	widthChanged := e.width != windowSettings.Width
	heightChanged := e.height != windowSettings.Height

	if widthChanged || heightChanged {
		e.width = windowSettings.Width
		e.height = windowSettings.Height
		ebiten.SetWindowSize(e.width, e.height)
	}

	if e.title != windowSettings.Title {
		e.title = windowSettings.Title
		ebiten.SetWindowTitle(e.title)
	}

	return nil
}
