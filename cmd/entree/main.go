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
	log.SetLevel("Debug")
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
	federicoSheet, err := sheets.FedericoSheet()
	if err != nil {
		log.Fatal("main", "federico_sheet", err)
	}
	g.LoadSpriteSheet(federicoSheet)

	onyawnSheet, err := sheets.OnyawnSheet()
	if err != nil {
		log.Fatal("main", "onyawn_sheet", err)
	}
	g.LoadSpriteSheet(onyawnSheet)

	tilesSheet, err := sheets.TilesSheet()
	if err != nil {
		log.Fatal("main", "tiles_sheet", err)
	}
	g.LoadSpriteSheet(tilesSheet)
}
