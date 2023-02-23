package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/ports"
	"github.com/maladroitthief/entree/service"
)

func main() {
	app := service.NewApplication()

  game, err := ports.NewEbitenGame(app)
	if err != nil {
		log.Fatal(err)
	}

	err = ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
