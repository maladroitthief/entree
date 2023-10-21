package driver

import (
	"fmt"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/common/theme"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/ui"
)

type EbitenGame struct {
	log   logs.Logger
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
	log logs.Logger,
	scene *ui.SceneManager,
) (*EbitenGame, error) {
	e := &EbitenGame{
		log:           log,
		scene:         scene,
		width:         0,
		height:        0,
		title:         "",
		theme:         &theme.TokyoNight{},
		scale:         1,
		spriteOptions: &ebiten.DrawImageOptions{},
		spriteSheets:  make(map[string]*ebiten.Image),
		sprites:       make(map[string]*ebiten.Image),
	}
	canvasWidth, canvasHeight := e.scene.Size()
	e.canvas = ebiten.NewImage(canvasWidth, canvasHeight)

	err := e.WindowHandler()

	return e, err
}

func (e *EbitenGame) Update() (err error) {
	// check for screen resizing
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
		Cursor: data.Vector{X: float64(cursorX), Y: float64(cursorY)},
		Keys:   keys,
	}

	err = e.scene.Update(inputState)

	if err == ui.Termination {
		return ebiten.Termination
	}

	return err
}

func (e *EbitenGame) Draw(screen *ebiten.Image) {
	e.canvas.Clear()
	screen.Fill(e.theme.Black())
	e.canvas.Fill(e.scene.BackgroundColor())

	state := e.scene.GetState()
	if state == nil {
		return
	}

	animations := state.GetAllAnimations()
	sort.Slice(
		animations,
		func(i, j int) bool { return animations[i].ZLayer < animations[j].ZLayer },
	)

	for _, animation := range animations {
		err := e.DrawAnimation(screen, state, animation)
		if err != nil {
			e.log.Error("Draw", animation, err)
		}
	}

	e.Render(screen)
	e.DrawDebug(screen)
}

func (e *EbitenGame) DrawAnimation(
	screen *ebiten.Image,
	world *core.ECS,
	animation attribute.Animation,
) (err error) {
	entity, entityErr := world.GetEntity(animation.EntityId)
	state, stateErr := world.GetState(entity.Id)
  position, positionErr := world.GetPosition(entity.Id)
  dimension, dimensionErr := world.GetDimension(entity.Id)

	if entityErr != nil {
		return nil
	}

	if positionErr != nil {
		return nil
	}

	if dimensionErr != nil {
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

	// Scale the sprite
	e.spriteOptions.GeoM.Scale(dimension.Scale, dimension.Scale)

	// Flip the sprite if moving west
	if stateErr == nil && state.OrientationX == attribute.West {
		e.spriteOptions.GeoM.Scale(-1, 1)
	}

	// Position the sprite and draw it
	e.spriteOptions.GeoM.Translate(
		position.Position.X+dimension.Offset.X,
		position.Position.Y+dimension.Offset.Y,
	)
	e.canvas.DrawImage(sprite, e.spriteOptions)

	return nil
}

func (e *EbitenGame) Render(screen *ebiten.Image) {
	m := ebiten.GeoM{}
	c := e.scene.GetCamera()
	position := c.Position
	m.Translate(-position.X, -position.Y)

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
	columns, rows := e.canvas.Bounds().Dx()/cellSize, e.canvas.Bounds().Dy()/cellSize

	for i := 0; i <= rows; i++ {
		vector.StrokeLine(
			e.canvas,
			0,
			float32(i*cellSize),
			float32(columns*cellSize),
			float32(i*cellSize),
			1,
			e.theme.Magenta(),
			false,
		)
	}

	for i := 0; i <= columns; i++ {
		vector.StrokeLine(
			e.canvas,
			float32(i*cellSize),
			0,
			float32(i*cellSize),
			float32(rows*cellSize),
			1,
			e.theme.Magenta(),
			false,
		)
	}
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
