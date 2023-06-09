package infrastructure

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/maladroitthief/entree/adapter"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/common/theme"
	"github.com/maladroitthief/entree/domain/canvas"
)

const (
	DefaultScale = 1
)

type EbitenGame struct {
	log      logs.Logger
	gameAdpt *adapter.GameAdapter

	width         int
	height        int
	title         string
	theme         theme.Colors
	scale         float64
	world         *ebiten.Image
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
		theme:         &theme.TokyoNight{},
		scale:         1,
		spriteOptions: &ebiten.DrawImageOptions{},
		spriteSheets:  make(map[string]*ebiten.Image),
		sprites:       make(map[string]*ebiten.Image),
	}
	canvasWidth, canvasHeight := e.gameAdpt.GetCanvasSize()
	e.world = ebiten.NewImage(canvasWidth, canvasHeight)

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
	e.world.Clear()
	screen.Fill(e.theme.Black())
	e.world.Fill(e.gameAdpt.GetBackgroundColor())

	entities := e.gameAdpt.GetEntities()
	for _, entity := range entities {
		err := e.DrawEntity(screen, entity)
		if err != nil {
			e.log.Error("Draw", entity, err)
		}
	}
	e.Render(screen)
	e.DrawDebug(screen)
}

func (e *EbitenGame) DrawEntity(screen *ebiten.Image, entity canvas.Entity) (err error) {
	// Load the sprite
	sprite, ok := e.sprites[SpriteKey(entity.Sheet(), entity.Sprite())]
	if !ok {
		sprite, err = e.LoadSprite(entity.Sheet(), entity.Sprite())
		if err != nil {
			return err
		}
	}

	// Reset the sprite options
	e.spriteOptions.GeoM.Reset()

	// Move the anchor point to the sprites center
	e.spriteOptions.GeoM.Translate(
		-float64(sprite.Bounds().Size().X)/2,
		-float64(sprite.Bounds().Size().Y)/2,
	)

	// Scale the sprite
	e.spriteOptions.GeoM.Scale(entity.Scale(), entity.Scale())

	// Flip the sprite if moving west
	if entity.OrientationX() == canvas.West {
		e.spriteOptions.GeoM.Scale(-1, 1)
	}

	// Position the sprite and draw it
	e.spriteOptions.GeoM.Translate(
		entity.Position().X+entity.Offset().X,
		entity.Position().Y+entity.Offset().Y,
	)
	e.world.DrawImage(sprite, e.spriteOptions)

	// Draw the debug rectangle
	vector.StrokeRect(
		e.world,
		float32(entity.Bounds().MinPoint.X),
		float32(entity.Bounds().MinPoint.Y),
		float32(entity.Bounds().Width()),
		float32(entity.Bounds().Height()),
		1,
		e.theme.Red(),
		false,
	)
	vector.DrawFilledCircle(
		e.world,
		float32(entity.Position().X),
		float32(entity.Position().Y),
		3,
		e.theme.BrightRed(),
		false,
	)

	return nil
}

func (e *EbitenGame) Render(screen *ebiten.Image) {
	m := ebiten.GeoM{}
	c := e.gameAdpt.GetCamera()
	position := c.Position()
	m.Translate(-position.X, -position.Y)

	// Scale around the center
	m.Translate(float64(e.width/2), float64(e.height/2))

	screen.DrawImage(e.world, &ebiten.DrawImageOptions{GeoM: m})
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
	w, h := e.gameAdpt.GetWindowSize()

	if e.width != w || e.height != h {
		e.width = w
		e.height = h
		ebiten.SetWindowSize(e.width, e.height)
	}

	if e.title != e.gameAdpt.GetWindowTitle() {
		e.title = e.gameAdpt.GetWindowTitle()
		ebiten.SetWindowTitle(e.title)
	}

	if e.scale != e.gameAdpt.GetScale() {
		e.scale = e.gameAdpt.GetScale()
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
