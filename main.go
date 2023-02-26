package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/adapter"
	"github.com/maladroitthief/entree/infrastructure"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.NewEntry(logrus.StandardLogger())
	sceneRepository := infrastructure.NewSceneMemoryRepository()
	sceneService := adapter.NewSceneService(logger, sceneRepository)
	game := adapter.NewGameService(logger, sceneService)

	ebitenGame, err := infrastructure.NewEbitenGame(game)
	if err != nil {
		log.Fatal(err)
	}

	err = ebiten.RunGame(ebitenGame)
	if err != nil {
		log.Fatal(err)
	}
}
