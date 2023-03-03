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

	sceneRepository := infrastructure.NewSceneMemoryRepository()
	sceneService := application.NewSceneService(log, sceneRepository)

	settingsRepository := infrastructure.NewSettingsJsonRepository("settings.json")
	gameService := application.NewGameService(
		log,
		settingsRepository,
		sceneService,
	)

	gameAdapter := adapter.NewGameAdapter(log, gameService)

	ebitenGame, err := infrastructure.NewEbitenGame(gameAdapter)
	if err != nil {
		log.Fatal("main", nil, err)
	}

	err = ebiten.RunGame(ebitenGame)
	if err != nil {
		log.Fatal("main", nil, err)
	}
}
