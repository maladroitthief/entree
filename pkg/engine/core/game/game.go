package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/maladroitthief/entree/pkg/ui/input"
	"github.com/maladroitthief/entree/pkg/ui/scene"
)

const (
  ScreenWidth = 256
  ScreenHeight = 240
)

type Game struct {
	input        input.Input
	sceneManager *scene.SceneManager
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
  if g.sceneManager == nil {
    g.sceneManager = &scene.SceneManager{}
  }

  g.input.Update()
  err := g.sceneManager.Update(&g.input)
  if err != nil {
    return err
  }
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}
