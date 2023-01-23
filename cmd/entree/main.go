package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/engine/core/game"
	"github.com/maladroitthief/entree/pkg/ui"
	"github.com/maladroitthief/entree/pkg/ui/input"
	"github.com/maladroitthief/entree/pkg/ui/scene"
)

func main() {
	// Content initialization

  // UserInterface Setup
	i := input.NewInput()
	sm := scene.NewSceneManager()
	ui := ui.NewUserInterface()
	ui.SetSceneManager(sm)
  ui.SetInput(i)

	// Push MainMenuScreen

	g := game.NewGame()
	g.SetUserInterface(ui)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
