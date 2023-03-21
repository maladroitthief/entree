package main

import (
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/adapter"
	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/infrastructure"
)

var (
	spriteSheetPaths = []string{
		filepath.Join("assets", "sprite_sheets", "test.json"),
	}
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
	spriteRepo := infrastructure.NewSpriteJsonRepository()
	graphicsSvc := application.NewGraphicsService(log, spriteRepo)
	for _, path := range spriteSheetPaths {
		err := graphicsSvc.LoadSpriteSheet(path)
		if err != nil {
			log.Fatal("main", path, err)
		}
	}

	// Scene Service
	sceneSvc := application.NewSceneService(log, settingsSvc, graphicsSvc)

	// Game adapter
	gameAdpt := adapter.NewGameAdapter(log, sceneSvc, settingsSvc)
	ebitenGame, err := infrastructure.NewEbitenGame(gameAdpt)
	if err != nil {
		log.Fatal("main", nil, err)
	}

	err = ebiten.RunGame(ebitenGame)
	if err != nil {
		log.Fatal("main", nil, err)
	}
}
