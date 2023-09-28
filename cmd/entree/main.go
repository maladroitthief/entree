package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/assets/sheets"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/driver"
	"github.com/maladroitthief/entree/pkg/ui"
)

func main() {
	// defer profile.Start().Stop()

	log := logs.NewLogrusLogger()
	settingsRepo := driver.NewSettingsRepository("settings.json")

	graphicsServer, err := ui.NewGraphicsServer(log)
	if err != nil {
		log.Fatal("main", "graphicsServer", err)
	}
	loadSpriteSheets(graphicsServer)

	inputHandler, err := ui.NewInputHandler(log, settingsRepo)
	if err != nil {
		log.Fatal("main", "inputHandler", err)
	}

	windowHandler, err := ui.NewWindowHandler(log, settingsRepo)
	if err != nil {
		log.Fatal("main", "windowHandler", err)
	}

	sceneManager, err := ui.NewSceneManager(
		log,
		graphicsServer,
		inputHandler,
		windowHandler,
	)
	if err != nil {
		log.Fatal("main", "sceneManager", err)
	}

	ebitenDriver, err := driver.NewEbitenDriver(log, sceneManager)
	if err != nil {
		log.Fatal("main", nil, err)
	}

	err = ebiten.RunGame(ebitenDriver)
	if err != nil {
		log.Fatal("main", nil, err)
	}
}

func loadSpriteSheets(g *ui.GraphicsServer) {
	heroSheet, err := sheets.HeroSheet()
	if err != nil {
		log.Fatal("main", "hero_sheet", err)
	}
	g.LoadSpriteSheet(heroSheet)

	testSheet, err := sheets.TestSheet()
	if err != nil {
		log.Fatal("main", "test_sheet", err)
	}
	g.LoadSpriteSheet(testSheet)
}
