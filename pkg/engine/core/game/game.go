package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/ui/input"
	"github.com/maladroitthief/entree/pkg/ui/scene"
	"github.com/maladroitthief/entree/pkg/ui/window"
)

type Game struct {
	input         *input.Input
	sceneManager  *scene.SceneManager
	windowManager *window.WindowManager
}

func NewGame() *Game {
	g := &Game{}

	return g
}

func (g *Game) SetInput(i *input.Input) {
	g.input = i
}

func (g *Game) SetSceneManager(sm *scene.SceneManager) {
	g.sceneManager = sm
}

func (g *Game) SetWindowManager(wm *window.WindowManager) {
	g.windowManager = wm
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.windowManager.GetWidth(), g.windowManager.GetHeight()
}

func (g *Game) Update() error {
	g.input.Update()
	err := g.sceneManager.Update()

	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(r *ebiten.Image) {
	g.sceneManager.Draw(r)
}
