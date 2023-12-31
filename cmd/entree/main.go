package main

import (
	"context"
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/assets/sheets"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/driver"
	"github.com/maladroitthief/entree/pkg/ui"
	"github.com/pkg/profile"
)

func main() {
	mode := flag.String("profile", "", "enable profiling mode, one of [cpu, mem, mutex, block]")
	flag.Parse()

	switch *mode {
	case "cpu":
		defer profile.Start(profile.ProfilePath("."), profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.ProfilePath("."), profile.MemProfile).Stop()
	case "mutex":
		defer profile.Start(profile.ProfilePath("."), profile.MutexProfile).Stop()
	case "block":
		defer profile.Start(profile.ProfilePath("."), profile.BlockProfile).Stop()
	default:
		// do nothing
	}

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

	ctx := context.Background()
	err = runGame(ctx, log, sceneManager)
	if err != nil {
		log.Fatal("main", nil, err)
	}
}

func runGame(ctx context.Context, log logs.Logger, sceneManager *ui.SceneManager) error {
	ebitenDriver, err := driver.NewEbitenDriver(ctx, log, sceneManager)
	if err != nil {
		return err
	}

	err = ebiten.RunGame(ebitenDriver)
	if err != nil {
		return err
	}
	return nil
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
