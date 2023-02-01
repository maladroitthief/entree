package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/engine/core/game"
	"github.com/maladroitthief/entree/pkg/ui/input"
	"github.com/maladroitthief/entree/pkg/ui/scene"
	"github.com/maladroitthief/entree/pkg/ui/window"
)

func main() {
	// Content initialization
	// UserInterface Setup
	i := input.NewInput()

	wm := window.NewWindowManager(1280, 720, "Entree")

	sm := scene.NewSceneManager()
  sm.SetWindowManager(wm)
	sm.SetInput(i)

	// Push MainMenuScreen
	sm.GoTo(&scene.TitleScene{})

	g := game.NewGame()
	g.SetSceneManager(sm)
	g.SetInput(i)
  g.SetWindowManager(wm)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
