package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/ui/input"
	"github.com/maladroitthief/entree/pkg/ui/scene"
	"github.com/maladroitthief/entree/pkg/ui/window"
)

type Game struct {
	input         input.InputHandler
	sceneManager  *scene.SceneManager
	windowManager *window.WindowManager
}

func main() {
	// Content initialization
	// UserInterface Setup
	i := input.NewInputHandler()

	wm := window.NewWindowManager(1280, 720, "Entree")

	sm := scene.NewSceneManager()
	sm.SetWindowManager(wm)
	sm.SetInput(i)

	// Push MainMenuScreen
	sm.GoTo(&scene.TitleScene{})

	g := &Game{
		input:         i,
		sceneManager:  sm,
		windowManager: wm,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
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
