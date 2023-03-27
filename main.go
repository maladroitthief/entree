package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/adapter"
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/assets"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/infrastructure"
)

func main() {
	log := logs.NewLogrusLogger()

	// Settings Service
	settingsRepo := infrastructure.NewSettingsJsonRepository("settings.json")
	settingsSvc := application.NewSettingsService(
		log,
		settingsRepo,
	)

	// Graphics Service
	graphicsSvc := application.NewGraphicsService(log)
	testSheet, err := assets.TestSheet()
	if err != nil {
		log.Fatal("main", "test_sheet", err)
	}
	graphicsSvc.LoadSpriteSheet(testSheet)

	// Scene Service
	sceneSvc := application.NewSceneService(log, settingsSvc)

	// Game adapter
	gameAdpt := adapter.NewGameAdapter(log, sceneSvc, graphicsSvc, settingsSvc)
	ebitenGame, err := infrastructure.NewEbitenGame(log, gameAdpt)
	if err != nil {
		log.Fatal("main", nil, err)
	}

	err = ebiten.RunGame(ebitenGame)
	if err != nil {
		log.Fatal("main", nil, err)
	}
}
