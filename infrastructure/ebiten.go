package infrastructure

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/maladroitthief/entree/adapter"
)

type EbitenGame struct {
	gameAdpt *adapter.GameAdapter

	width  int
	height int
	title  string
}

func NewEbitenGame(gameAdpt *adapter.GameAdapter) (*EbitenGame, error) {
	e := &EbitenGame{
		gameAdpt: gameAdpt,
		width:    0,
		height:   0,
		title:    "",
	}

	err := e.WindowHandler()

	return e, err
}

func (e *EbitenGame) Update() (err error) {
	// check for screen resizing
	err = e.WindowHandler()
	if err != nil {
		return err
	}

	// grab cursor coordinates
	cursorX, cursorY := ebiten.CursorPosition()

	// grab current keyboard inputs
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	inputs := []string{}
	for _, k := range pressedKeys {
		inputs = append(inputs, k.String())
	}

	args := adapter.UpdateArgs{
		CursorX: cursorX,
		CursorY: cursorY,
		Inputs:  inputs,
	}

	// update the main game
	return e.gameAdpt.Update(args)
}

func (e *EbitenGame) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf(`
    TPS: %0.2f
    FPS: %0.2f`,
		ebiten.ActualTPS(),
		ebiten.ActualFPS(),
	)
	ebitenutil.DebugPrint(screen, msg)
}

func (e *EbitenGame) Layout(width, height int) (screenWidth, screenHeight int) {
	return e.gameAdpt.Layout(width, height)
}

func (e *EbitenGame) WindowHandler() error {
	windowSettings, err := e.gameAdpt.GetWindowSettings()
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
