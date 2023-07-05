package game

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

type Game struct {
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

func NewGame(
	log logs.Logger,
	sceneSvc service.SceneService,
	graphicsSvc service.GraphicsService,
	settingsSvc service.SettingsService,
) (*Game, error) {
	g := Game{
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

	return &g, nil
}

func (g *Game) Update(args UpdateArgs) error {
	// Scene Update
	err := g.sceneSvc.Update(service.Inputs{
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

func (g *Game) GetCamera() scene.Camera {
	return g.sceneSvc.GetCamera()
}

func (g *Game) GetCanvasSize() (width, height int) {
	return g.sceneSvc.GetCanvasSize()
}

func (g *Game) GetCanvasCellSize() int {
	return g.sceneSvc.GetCanvasCellSize()
}

func (g *Game) GetEntities() []canvas.Entity {
	return g.sceneSvc.GetEntities()
}

func (g *Game) GetSpriteSheet(sheet string) (sprite.SpriteSheet, error) {
	return g.graphicsSvc.GetSpriteSheet(sheet)
}

func (g *Game) GetSpriteRectangle(
	sheet string,
	sprite string,
) (image.Rectangle, error) {
	return g.graphicsSvc.GetSprite(sheet, sprite)
}

func (g *Game) Layout(width, height int) (screenWidth, screenHeight int) {
	return g.GetWindowSize()
}

func (g *Game) GetWindowSize() (screenWidth, screenHeight int) {
	return g.settingsSvc.GetWindowWidth(), g.settingsSvc.GetWindowHeight()
}

func (g *Game) GetWindowTitle() string {
	return g.settingsSvc.GetWindowTitle()
}

func (g *Game) GetScale() float64 {
	return g.settingsSvc.GetScale()
}

func (g *Game) GetBackgroundColor() color.Color {
	return g.sceneSvc.GetBackgroundColor()
}
