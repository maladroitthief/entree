package adapter

import (
	"image"
	"image/color"

	"github.com/maladroitthief/entree/application"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/sprite"
)

type GameAdapter struct {
	log         logs.Logger
	sceneSvc    *application.SceneService
	graphicsSvc *application.GraphicsService
	settingsSvc *application.SettingsService
}

type UpdateArgs struct {
	CursorX int
	CursorY int
	Inputs  []string
}

func NewGameAdapter(
	log logs.Logger,
	sceneSvc *application.SceneService,
	graphicsSvc *application.GraphicsService,
	settingsSvc *application.SettingsService,
) *GameAdapter {
	ga := GameAdapter{
		log:         log,
		settingsSvc: settingsSvc,
		sceneSvc:    sceneSvc,
		graphicsSvc: graphicsSvc,
	}

	if sceneSvc == nil {
		panic("nil scene service")
	}

	if graphicsSvc == nil {
		panic("nil graphics service")
	}

	if settingsSvc == nil {
		panic("nil settings service")
	}

	return &ga
}

func (ga *GameAdapter) Update(args UpdateArgs) error {
	// Scene Update
	err := ga.sceneSvc.Update(application.Inputs{
		CursorX: args.CursorX,
		CursorY: args.CursorY,
		Inputs:  args.Inputs,
	})

	if err != nil {
		return err
	}

	return nil
}

func (ga *GameAdapter) GetEntities() []canvas.Entity {
	e := ga.sceneSvc.GetEntities()

	return e
}

func (ga *GameAdapter) GetSpriteSheet(
	sheet string,
) (sprite.SpriteSheet, error) {
	return ga.graphicsSvc.GetSpriteSheet(sheet)
}

func (ga *GameAdapter) GetSpriteRectangle(
	sheet string,
	sprite string,
) (image.Rectangle, error) {
	return ga.graphicsSvc.GetSprite(sheet, sprite)
}

func (ga *GameAdapter) Layout(width, height int) (screenWidth, screenHeight int) {
	return ga.GetWindowSize()
}

func (ga *GameAdapter) GetWindowSize() (screenWidth, screenHeight int) {
	return ga.settingsSvc.GetWindowWidth(), ga.settingsSvc.GetWindowHeight()
}

func (ga *GameAdapter) GetWindowTitle() string {
	return ga.settingsSvc.GetWindowTitle()
}

func (ga *GameAdapter) GetScale() float64 {
	return ga.settingsSvc.GetScale()
}

func (ga *GameAdapter) GetBackgroundColor() color.Color {
	return ga.sceneSvc.GetBackgroundColor()
}
