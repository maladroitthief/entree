package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/game"
)

const (
	gameTitle = "Entree"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle(gameTitle)

	game, err := game.New()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}

}
