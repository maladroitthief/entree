package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/adapter"
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/assets/sheets"
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
	// Update it once to initialize the service
	err := settingsSvc.Update(application.Inputs{})
	if err != nil {
		log.Fatal("main", "settingsSvc", err)
	}

	// Graphics Service
	graphicsSvc := application.NewGraphicsService(log)
	pilotSheet, err := sheets.PilotSheet()
	if err != nil {
		log.Fatal("main", "pilot_sheet", err)
	}
	graphicsSvc.LoadSpriteSheet(pilotSheet)
	testSheet, err := sheets.TestSheet()
	if err != nil {
		log.Fatal("main", "test_sheet", err)
	}
	graphicsSvc.LoadSpriteSheet(testSheet)

	// Scene Service
	sceneSvc := application.NewSceneService(log, settingsSvc)

	// Game adapter
	gameAdpt, err := adapter.NewGameAdapter(log, sceneSvc, graphicsSvc, settingsSvc)
	if err != nil {
		log.Fatal("main", nil, err)
	}

	// Ebiten driver
	ebitenGame, err := infrastructure.NewEbitenGame(log, gameAdpt)
	if err != nil {
		log.Fatal("main", nil, err)
	}

	err = ebiten.RunGame(ebitenGame)
	if err != nil {
		log.Fatal("main", nil, err)
	}
}
