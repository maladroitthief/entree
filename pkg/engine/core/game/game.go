package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/maladroitthief/entree/pkg/ui"
)

const (
	ScreenWidth  = 256
	ScreenHeight = 240
)

type Game struct {
	ui *ui.UserInterface
}

func NewGame() *Game {
  g := &Game{}

  return g
}

func (g *Game) SetUserInterface(userInterface *ui.UserInterface) {
  g.ui = userInterface
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	err := g.ui.Update()

	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}
