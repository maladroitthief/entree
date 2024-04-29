package main

import (
	"context"
	"flag"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/assets/sheets"
	"github.com/maladroitthief/entree/driver"
	"github.com/maladroitthief/entree/pkg/ui"
	"github.com/pkg/profile"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	mode := flag.String("profile", "", "enable profiling mode, one of [cpu, mem, mutex, block]")
	debug := flag.Bool("debug", false, "sets log level to debug")
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

	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: zerolog.TimeFormatUnixMs})

	ctx := context.Background()
	settingsRepo := driver.NewSettingsRepository("settings.json")

	graphicsServer, err := ui.NewGraphicsServer()
	if err != nil {
		log.Fatal().Err(err).Any("graphics", graphicsServer)
	}
	loadSpriteSheets(graphicsServer)

	inputHandler, err := ui.NewInputHandler(settingsRepo)
	if err != nil {
		log.Fatal().Err(err).Any("input", inputHandler)
	}

	windowHandler, err := ui.NewWindowHandler(settingsRepo)
	if err != nil {
		log.Fatal().Err(err).Any("window", windowHandler)
	}

	sceneManager, err := ui.NewSceneManager(
		ctx,
		graphicsServer,
		inputHandler,
		windowHandler,
	)
	if err != nil {
		log.Fatal().Err(err).Any("scene", sceneManager)
	}

	err = runGame(ctx, sceneManager)
	if err != nil {
		log.Fatal().Err(err)
	}
}

func runGame(ctx context.Context, sceneManager *ui.SceneManager) error {
	ebitenDriver, err := driver.NewEbitenDriver(ctx, sceneManager)
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
		log.Fatal().Err(err).Any("federico_sheet", federicoSheet)
	}
	g.LoadSpriteSheet(federicoSheet)

	onyawnSheet, err := sheets.OnyawnSheet()
	if err != nil {
		log.Fatal().Err(err).Any("onyawn_sheet", onyawnSheet)
	}
	g.LoadSpriteSheet(onyawnSheet)

	tilesSheet, err := sheets.TilesSheet()
	if err != nil {
		log.Fatal().Err(err).Any("tiles_sheet", tilesSheet)
	}
	g.LoadSpriteSheet(tilesSheet)
}
