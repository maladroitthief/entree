package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/adapter"
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/infrastructure"
)

func main() {
	log := logs.NewLogrusLogger()

  // Scene Service
	sceneRepo := infrastructure.NewSceneMemoryRepository()
	sceneSvc := application.NewSceneService(log, sceneRepo)

  // Game Service
	gameRepo := infrastructure.NewGameMemoryRepository()
	settingsRepo := infrastructure.NewSettingsJsonRepository("settings.json")
	gameSvc := application.NewGameService(
		log,
		gameRepo,
		settingsRepo,
	)

	gameAdpt := adapter.NewGameAdapter(log, gameSvc, sceneSvc)

	ebitenGame, err := infrastructure.NewEbitenGame(gameAdpt)
	if err != nil {
		log.Fatal("main", nil, err)
	}

	err = ebiten.RunGame(ebitenGame)
	if err != nil {
		log.Fatal("main", nil, err)
	}
}
