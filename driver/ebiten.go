package driver

import (
	"context"
	"fmt"
	"image/color"
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/maladroitthief/entree/common/theme"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/ui"
	"github.com/maladroitthief/mosaic"
	"github.com/rs/zerolog/log"
)

type EbitenGame struct {
	ctx   context.Context
	scene *ui.SceneManager

	width         int
	height        int
	title         string
	theme         theme.Colors
	scale         float64
	canvas        *ebiten.Image
	spriteOptions *ebiten.DrawImageOptions
	spriteSheets  map[string]*ebiten.Image
	sprites       map[string]*ebiten.Image
}

func NewEbitenDriver(
	ctx context.Context,
	scene *ui.SceneManager,
) (*EbitenGame, error) {
	e := &EbitenGame{
		ctx:           ctx,
		scene:         scene,
		width:         0,
		height:        0,
		title:         "",
		theme:         &theme.Endesga32{},
		scale:         1,
		spriteOptions: &ebiten.DrawImageOptions{},
		spriteSheets:  make(map[string]*ebiten.Image),
		sprites:       make(map[string]*ebiten.Image),
	}
	canvasWidth, canvasHeight := e.scene.Size()
	e.canvas = ebiten.NewImage(canvasWidth, canvasHeight)

	err := e.WindowHandler()

	ebiten.SetVsyncEnabled(false)
	// ebiten.SetTPS(ebiten.SyncWithFPS)

	return e, err
}

func (e *EbitenGame) Update() (err error) {
	// check for screen resizing
	select {
	case <-e.ctx.Done():
		return ebiten.Termination
	default:
		err = e.WindowHandler()
		if err != nil {
			return err
		}
		e.CanvasHandler()

		cursorX, cursorY := ebiten.CursorPosition()

		pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
		keys := []string{}
		for _, key := range pressedKeys {
			keys = append(keys, key.String())
		}

		inputState := ui.InputState{
			Cursor: mosaic.Vector{X: float64(cursorX), Y: float64(cursorY)},
			Keys:   keys,
		}

		err = e.scene.Update(inputState)

		if err == ui.Termination {
			return ebiten.Termination
		}

		return err
	}
}

func (e *EbitenGame) Draw(screen *ebiten.Image) {
	e.canvas.Clear()
	screen.Fill(e.theme.Black())
	e.canvas.Fill(e.scene.BackgroundColor())

	state := e.scene.GetState()
	if state == nil {
		return
	}

	positions := state.GetAllPositions()
	sort.Slice(
		positions,
		func(i, j int) bool {
			if math.Ceil(positions[i].Z) != math.Ceil(positions[j].Z) {
				return positions[i].Z < positions[j].Z
			}
			return positions[i].Y < positions[j].Y
		},
	)

	for _, position := range positions {
		err := e.DrawAnimation(screen, state, position)
		if err != nil {
			log.Warn().Err(err).Any("position", position)
		}
	}

	e.Render(screen)
	e.DebugFPS(screen)
}

func (e *EbitenGame) DrawAnimation(
	screen *ebiten.Image,
	world *core.ECS,
	position core.Position,
) (err error) {
	entity, entityErr := world.GetEntity(position.EntityId)
	state, stateErr := world.GetState(entity)
	animation, animationErr := world.GetAnimation(entity)
	dimension, dimensionErr := world.GetDimension(entity)

	if entityErr != nil {
		return nil
	}

	if dimensionErr != nil {
		return nil
	}

	if animationErr != nil {
		bounds := dimension.Bounds()
		e.DebugBounds(bounds, e.theme.Blue())
		return nil
	}

	sprite, ok := e.sprites[SpriteKey(animation.SpriteSheet, animation.Sprite)]
	if !ok {
		sprite, err = e.LoadSprite(animation.SpriteSheet, animation.Sprite)
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

	// Flip the sprite if moving west
	if stateErr == nil && state.OrientationX == core.West {
		e.spriteOptions.GeoM.Scale(-1, 1)
	}

	// Position the sprite and draw it
	e.spriteOptions.GeoM.Translate(
		position.X+dimension.Offset.X,
		position.Y+dimension.Offset.Y,
	)
	e.canvas.DrawImage(sprite, e.spriteOptions)

	// bounds := dimension.Bounds()
	// vector.StrokeRect(
	// 	e.canvas,
	// 	float32(bounds.Position.X-bounds.Width/2),
	// 	float32(bounds.Position.Y-bounds.Height/2),
	// 	float32(bounds.Width),
	// 	float32(bounds.Height),
	// 	1,
	// 	e.theme.Red(),
	// 	false,
	// )

	// msg := fmt.Sprintf(
	// 	"[%v]",
	// 	entity.Id,
	// )
	// ebitenutil.DebugPrintAt(e.canvas, msg, int(position.X), int(position.Y))

	// msg := fmt.Sprintf(
	// 	"[%0.2f, %0.2f]",
	// 	position.X,
	// 	position.Y,
	// )
	// ebitenutil.DebugPrintAt(e.canvas, msg, int(position.X), int(position.Y))

	return nil
}

func (e *EbitenGame) Render(screen *ebiten.Image) {
	m := ebiten.GeoM{}
	c := e.scene.GetCamera()
	m.Translate(-c.X, -c.Y)

	// Scale around the center
	zoom := c.Zoom
	m.Scale(zoom, zoom)
	m.Translate(float64(e.width/2), float64(e.height/2))

	screen.DrawImage(e.canvas, &ebiten.DrawImageOptions{GeoM: m})
}

func (e *EbitenGame) LoadSpriteSheet(sheet string) (*ebiten.Image, error) {
	// Get the sprite sheet
	ss, err := e.scene.SpriteSheet(sheet)
	if err != nil {
		return nil, err
	}

	// Create the sprite sheet image and cache it
	e.spriteSheets[sheet] = ebiten.NewImageFromImage(ss.Image())

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
	spriteRectangle, err := e.scene.SpriteRectangle(sheetName, spriteName)
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
	return e.scene.Layout(width, height)
}

func (e *EbitenGame) WindowHandler() error {
	w, h := e.scene.WindowSize()

	if e.width != w || e.height != h {
		e.width = w
		e.height = h
		ebiten.SetWindowSize(e.width, e.height)
	}

	if e.title != e.scene.WindowTitle() {
		e.title = e.scene.WindowTitle()
		ebiten.SetWindowTitle(e.title)
	}

	if e.scale != e.scene.Scale() {
		e.scale = e.scene.Scale()
	}

	return nil
}

func (e *EbitenGame) CanvasHandler() {
	canvasWidth, canvasHeight := e.scene.Size()
	x, y := e.canvas.Bounds().Dx(), e.canvas.Bounds().Dy()
	if canvasWidth != x || canvasHeight != y {
		e.canvas = ebiten.NewImage(canvasWidth, canvasHeight)
	}
}

func (e *EbitenGame) DrawGrid() {
	cellSize := e.scene.CellSize()
	gridX, gridY := e.canvas.Bounds().Dx()/cellSize, e.canvas.Bounds().Dy()/cellSize

	for i := 0; i <= gridY; i++ {
		vector.StrokeLine(
			e.canvas,
			0,
			float32(i*cellSize),
			float32(gridX*cellSize),
			float32(i*cellSize),
			1,
			e.theme.Magenta(),
			false,
		)
	}

	for i := 0; i <= gridX; i++ {
		vector.StrokeLine(
			e.canvas,
			float32(i*cellSize),
			0,
			float32(i*cellSize),
			float32(gridY*cellSize),
			1,
			e.theme.Magenta(),
			false,
		)
	}
}

func (e *EbitenGame) DebugFPS(screen *ebiten.Image) {
	msg := fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f",
		ebiten.ActualTPS(),
		ebiten.ActualFPS(),
	)
	ebitenutil.DebugPrint(screen, msg)
}

func (e *EbitenGame) DebugEntity(entity core.Entity, x, y int) {
	msg := fmt.Sprintf(
		"%v",
		entity.Id,
	)
	ebitenutil.DebugPrintAt(e.canvas, msg, x, y)
}

func (e *EbitenGame) DebugBounds(bounds mosaic.Rectangle, color color.Color) {
	vector.StrokeRect(
		e.canvas,
		float32(bounds.Position.X-bounds.Width/2),
		float32(bounds.Position.Y-bounds.Height/2),
		float32(bounds.Width),
		float32(bounds.Height),
		1,
		color,
		false,
	)
}

func SpriteKey(sheet, sprite string) string {
	return fmt.Sprintf("%s_%s", sheet, sprite)
}
