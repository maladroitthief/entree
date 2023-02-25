package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/adapter"
	"github.com/maladroitthief/entree/infrastructure"
)

func main() {
	game := adapter.NewGameService()

	ebitenGame, err := infrastructure.NewEbitenGame(game)
	if err != nil {
		log.Fatal(err)
	}

	err = ebiten.RunGame(ebitenGame)
	if err != nil {
		log.Fatal(err)
	}
}
