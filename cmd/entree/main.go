package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/engine/core/game"
)

func main() {
  g := game.NewGame()
  // Content initialization
  // UI initialization
  // Keybinding
  // Push MainMenuScreen

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
