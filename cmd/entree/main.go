package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/engine/core/game"
	"github.com/maladroitthief/entree/pkg/ui/screen"
)

func main() {
	sm := screen.NewScreenManager("Entree", 800, 600)
	sm.Update()

	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
