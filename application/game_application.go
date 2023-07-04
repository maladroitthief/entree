package application

import (
	"errors"
	"image"
	"image/color"

	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/scene"
	"github.com/maladroitthief/entree/domain/sprite"
	"github.com/maladroitthief/entree/service"
)

var (
	ErrLoggerNil          = errors.New("nil logger")
	ErrSceneServiceNil    = errors.New("nil scene service")
	ErrGraphicsServiceNil = errors.New("nil graphics service")
	ErrSettingsServiceNil = errors.New("nil settings service")
	Termination           = errors.New("game exited normally")
)

type GameApplication struct {
	log         logs.Logger
	sceneSvc    service.SceneService
	graphicsSvc service.GraphicsService
	settingsSvc service.SettingsService
}

type UpdateArgs struct {
	CursorX int
	CursorY int
	Inputs  []string
}

func NewGameApplication(
	log logs.Logger,
	sceneSvc service.SceneService,
	graphicsSvc service.GraphicsService,
	settingsSvc service.SettingsService,
) (*GameApplication, error) {
	ga := GameApplication{
		log:         log,
		settingsSvc: settingsSvc,
		sceneSvc:    sceneSvc,
		graphicsSvc: graphicsSvc,
	}

	if log == nil {
		return nil, ErrLoggerNil
	}

	if sceneSvc == nil {
		return nil, ErrSceneServiceNil
	}

	if graphicsSvc == nil {
		return nil, ErrGraphicsServiceNil
	}

	if settingsSvc == nil {
		return nil, ErrSettingsServiceNil
	}

	return &ga, nil
}

func (ga *GameApplication) Update(args UpdateArgs) error {
	// Scene Update
	err := ga.sceneSvc.Update(service.Inputs{
		CursorX: args.CursorX,
		CursorY: args.CursorY,
		Inputs:  args.Inputs,
	})

	if err == scene.SceneTermination {
		return Termination
	}

	if err != nil {
		return err
	}

	return nil
}

func (ga *GameApplication) GetCamera() scene.Camera {
	return ga.sceneSvc.GetCamera()
}

func (ga *GameApplication) GetCanvasSize() (width, height int) {
	return ga.sceneSvc.GetCanvasSize()
}

func (ga *GameApplication) GetCanvasCellSize() int {
	return ga.sceneSvc.GetCanvasCellSize()
}

func (ga *GameApplication) GetEntities() []canvas.Entity {
	return ga.sceneSvc.GetEntities()
}

func (ga *GameApplication) GetSpriteSheet(sheet string) (sprite.SpriteSheet, error) {
	return ga.graphicsSvc.GetSpriteSheet(sheet)
}

func (ga *GameApplication) GetSpriteRectangle(
	sheet string,
	sprite string,
) (image.Rectangle, error) {
	return ga.graphicsSvc.GetSprite(sheet, sprite)
}

func (ga *GameApplication) Layout(width, height int) (screenWidth, screenHeight int) {
	return ga.GetWindowSize()
}

func (ga *GameApplication) GetWindowSize() (screenWidth, screenHeight int) {
	return ga.settingsSvc.GetWindowWidth(), ga.settingsSvc.GetWindowHeight()
}

func (ga *GameApplication) GetWindowTitle() string {
	return ga.settingsSvc.GetWindowTitle()
}

func (ga *GameApplication) GetScale() float64 {
	return ga.settingsSvc.GetScale()
}

func (ga *GameApplication) GetBackgroundColor() color.Color {
	return ga.sceneSvc.GetBackgroundColor()
}
