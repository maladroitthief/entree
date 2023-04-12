package infrastructure

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/maladroitthief/entree/adapter"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/canvas"
)

const (
  DefaultScale = 3
)

type EbitenGame struct {
	log      logs.Logger
	gameAdpt *adapter.GameAdapter

	width         int
	height        int
	title         string
	spriteOptions *ebiten.DrawImageOptions
	spriteSheets  map[string]*ebiten.Image
	sprites       map[string]*ebiten.Image
}

func NewEbitenGame(
	log logs.Logger,
	gameAdpt *adapter.GameAdapter,
) (*EbitenGame, error) {
	e := &EbitenGame{
		log:           log,
		gameAdpt:      gameAdpt,
		width:         0,
		height:        0,
		title:         "",
		spriteOptions: &ebiten.DrawImageOptions{},
		spriteSheets:  make(map[string]*ebiten.Image),
		sprites:       make(map[string]*ebiten.Image),
	}

	err := e.WindowHandler()

	return e, err
}

func (e *EbitenGame) Update() (err error) {
	// check for screen resizing
	err = e.WindowHandler()
	if err != nil {
		return err
	}

	// grab cursor coordinates
	cursorX, cursorY := ebiten.CursorPosition()

	// grab current keyboard inputs
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	inputs := []string{}
	for _, key := range pressedKeys {
		inputs = append(inputs, key.String())
	}

	args := adapter.UpdateArgs{
		CursorX: cursorX,
		CursorY: cursorY,
		Inputs:  inputs,
	}

	// update the main game
	return e.gameAdpt.Update(args)
}

func (e *EbitenGame) Draw(screen *ebiten.Image) {
  screen.Fill(e.gameAdpt.GetBackgroundColor())

	entities := e.gameAdpt.GetEntities()
	for _, entity := range entities {
		err := e.DrawEntity(screen, entity)
		if err != nil {
			e.log.Error("Draw", entity, err)
		}
	}
	e.DrawDebug(screen)
}

func (e *EbitenGame) DrawEntity(screen *ebiten.Image, entity *canvas.Entity) (err error) {
	// Load the sprite
	sprite, ok := e.sprites[SpriteKey(entity.Sheet, entity.Sprite)]
	if !ok {
		sprite, err = e.LoadSprite(entity.Sheet, entity.Sprite)
		if err != nil {
			return err
		}
	}

	// Draw the sprite
	e.spriteOptions.GeoM.Reset()
	e.spriteOptions.GeoM.Scale(DefaultScale, DefaultScale)
  if entity.OrientationX == canvas.West {
    e.spriteOptions.GeoM.Scale(-1, 1)
    e.spriteOptions.GeoM.Translate(float64(entity.Width)*DefaultScale, 0)
  }
	e.spriteOptions.GeoM.Translate(
		float64(entity.Width)/2,
		float64(entity.Height)/2,
	)
	e.spriteOptions.GeoM.Translate(float64(entity.X), float64(entity.Y))
	screen.DrawImage(sprite, e.spriteOptions)

	return nil
}

func (e *EbitenGame) LoadSpriteSheet(sheet string) (*ebiten.Image, error) {
	// Get the sprite sheet
	ss, err := e.gameAdpt.GetSpriteSheet(sheet)
	if err != nil {
		return nil, err
	}

	// Create the sprite sheet image and cache it
	e.spriteSheets[sheet] = ebiten.NewImageFromImage(ss.GetImage())

	return e.spriteSheets[sheet], nil
}

func (e *EbitenGame) LoadSprite(
	sheetName string,
	spriteName string,
) (spriteImage *ebiten.Image, err error) {
	// Load the sprite sheet
	spriteSheet, ok := e.spriteSheets[sheetName]
	if !ok {
		spriteSheet, err = e.LoadSpriteSheet(sheetName)
		if err != nil {
			return nil, err
		}
	}

	// Get the sprite rectangle and the sprite sheet sub image
	spriteRectangle, err := e.gameAdpt.GetSpriteRectangle(sheetName, spriteName)
	if err != nil {
		return nil, err
	}
	sprite := spriteSheet.SubImage(spriteRectangle)

	// Create the ebiten image for the sprite and cache it
	spriteImage = ebiten.NewImageFromImage(sprite)
	e.sprites[SpriteKey(sheetName, spriteName)] = spriteImage

	return spriteImage, nil
}

func (e *EbitenGame) Layout(width, height int) (screenWidth, screenHeight int) {
	return e.gameAdpt.Layout(width, height)
}

func (e *EbitenGame) WindowHandler() error {
	windowSettings, err := e.gameAdpt.GetWindowSettings()
	if err != nil {
		return err
	}

	widthChanged := e.width != windowSettings.Width
	heightChanged := e.height != windowSettings.Height

	if widthChanged || heightChanged {
		e.width = windowSettings.Width
		e.height = windowSettings.Height
		ebiten.SetWindowSize(e.width, e.height)
	}

	if e.title != windowSettings.Title {
		e.title = windowSettings.Title
		ebiten.SetWindowTitle(e.title)
	}

	return nil
}

func (e *EbitenGame) DrawDebug(screen *ebiten.Image) {
	msg := fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f",
		ebiten.ActualTPS(),
		ebiten.ActualFPS(),
	)
	ebitenutil.DebugPrint(screen, msg)
}

func SpriteKey(sheet, sprite string) string {
	return fmt.Sprintf("%s_%s", sheet, sprite)
}
