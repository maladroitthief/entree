package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/engine/core/game"
	"github.com/maladroitthief/entree/pkg/ui/window"
)

func main() {
	wm := window.NewWindowManager("Entree", 800, 600)
	wm.Update()

	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
