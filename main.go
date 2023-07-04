package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/assets/sheets"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/infrastructure"
	"github.com/maladroitthief/entree/service"
)

func main() {
	log := logs.NewLogrusLogger()

	// Settings Service
	settingsRepo := infrastructure.NewSettingsJsonRepository("settings.json")
	settingsSvc, err := service.NewSettingsService(
		log,
		settingsRepo,
	)
	if err != nil {
		log.Fatal("main", "settingsSvc", err)
	}

	// Update it once to initialize the service
	err = settingsSvc.Update(service.Inputs{})
	if err != nil {
		log.Fatal("main", "settingsSvc", err)
	}

	// Graphics Service
	graphicsSvc, err := service.NewGraphicsService(log)
	if err != nil {
		log.Fatal("main", "graphicsSvc", err)
	}
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
	sceneSvc, err := service.NewSceneService(log, settingsSvc)
	if err != nil {
		log.Fatal("main", "sceneSvc", err)
	}

	// Game application
	gameApp, err := application.NewGameApplication(log, sceneSvc, graphicsSvc, settingsSvc)
	if err != nil {
		log.Fatal("main", nil, err)
	}

	// Ebiten driver
	ebitenGame, err := infrastructure.NewEbitenGame(log, gameApp)
	if err != nil {
		log.Fatal("main", nil, err)
	}

	err = ebiten.RunGame(ebitenGame)
	if err != nil {
		log.Fatal("main", nil, err)
	}
}
