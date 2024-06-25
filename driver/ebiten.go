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
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/ui"
	"github.com/maladroitthief/lattice"
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
		scale:         4,
		spriteOptions: &ebiten.DrawImageOptions{},
		spriteSheets:  make(map[string]*ebiten.Image),
		sprites:       make(map[string]*ebiten.Image),
	}
	canvasWidth, canvasHeight := e.scene.Size()
	e.canvas = ebiten.NewImage(canvasWidth, canvasHeight)

	err := e.WindowHandler()

	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(60)

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

	world := e.scene.GetWorld()
	if world == nil {
		return
	}

	positions := world.ECS.GetAllPositions()
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
		err := e.DrawAnimation(screen, world, position)
		if err != nil {
			log.Warn().Err(err).Any("position", position)
		}
	}

	e.Render(screen)
	e.DebugFPS(screen)
}

func (e *EbitenGame) DrawAnimation(
	screen *ebiten.Image,
	world *content.World,
	position core.Position,
) (err error) {
	entity, entityErr := world.ECS.GetEntity(position.EntityId)
	state, stateErr := world.ECS.GetState(entity)
	animation, animationErr := world.ECS.GetAnimation(entity)
	dimension, dimensionErr := world.ECS.GetDimension(entity)

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

	e.spriteOptions.GeoM.Scale(e.scale, e.scale)
	// Position the sprite and draw it
	e.spriteOptions.GeoM.Translate(
		math.Ceil((position.X+dimension.Offset.X)*e.scale),
		math.Ceil((position.Y+dimension.Offset.Y)*e.scale),
	)
	e.canvas.DrawImage(sprite, e.spriteOptions)

	// ai, err := world.ECS.GetAI(entity)
	// if err != nil {
	// 	return nil
	// }

	// e.DebugWeights(world.Grid, e.theme.White())
	// e.DebugPathfinding(world, ai, e.theme.Cyan())
	// e.DebugNode(world.Grid, e.theme.White())

	return nil
}

func (e *EbitenGame) Render(screen *ebiten.Image) {
	m := ebiten.GeoM{}
	c := e.scene.GetCamera()
	m.Translate(
		math.Ceil(-c.X*e.scale),
		math.Ceil(-c.Y*e.scale),
	)

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

	return nil
}

func (e *EbitenGame) CanvasHandler() {
	canvasWidth, canvasHeight := e.scene.Size()
	canvasHeight *= int(e.scale)
	canvasWidth *= int(e.scale)
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
		float32(bounds.Position.X-bounds.Width()/2),
		float32(bounds.Position.Y-bounds.Height()/2),
		float32(bounds.Width()),
		float32(bounds.Height()),
		1,
		color,
		false,
	)
}

func (e *EbitenGame) DebugPathfinding(world *content.World, ai core.AI, color color.Color) {
	entity, err := world.ECS.GetEntity(ai.EntityId)
	if err != nil {
		return
	}

	from, err := world.ECS.GetPosition(entity)
	if err != nil {
		return
	}

	x := len(world.Grid.Nodes)
	y := len(world.Grid.Nodes[0])
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			to := mosaic.Vector{
				X: ((float64(i) * world.Grid.ChunkSize) + world.Grid.ChunkSize/2) * e.scale,
				Y: ((float64(j) * world.Grid.ChunkSize) + world.Grid.ChunkSize/2) * e.scale,
			}
			weight := world.Grid.GetLocationWeight(i, j)
			heuristic := math.Abs(from.X-to.X) + math.Abs(from.Y-to.Y) + weight
			msg := fmt.Sprintf("%v", heuristic)
			ebitenutil.DebugPrintAt(
				e.canvas,
				msg,
				int(((float64(i)*world.Grid.ChunkSize)+world.Grid.ChunkSize/2)*e.scale),
				int(((float64(j)*world.Grid.ChunkSize)+world.Grid.ChunkSize/2)*e.scale)+32,
			)
		}
	}

	paths := ai.PathToTarget
	for i := 0; i < len(paths)-1; i++ {
		current := paths[i]
		next := paths[i+1]

		vector.StrokeLine(
			e.canvas,
			float32(current.X*e.scale),
			float32(current.Y*e.scale),
			float32(next.X*e.scale),
			float32(next.Y*e.scale),
			5,
			color,
			false,
		)
	}
}

func (e *EbitenGame) DebugWeights(grid *lattice.SpatialGrid[core.Entity], color color.Color) {
	x := len(grid.Nodes)
	y := len(grid.Nodes[0])

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			weight := grid.GetLocationWeight(i, j)
			msg := fmt.Sprintf("%v", weight)
			ebitenutil.DebugPrintAt(
				e.canvas,
				msg,
				int(((float64(i)*grid.ChunkSize)+grid.ChunkSize/2)*e.scale),
				int(((float64(j)*grid.ChunkSize)+grid.ChunkSize/2)*e.scale),
			)
		}
	}
}

func (e *EbitenGame) DebugNode(grid *lattice.SpatialGrid[core.Entity], color color.Color) {
	x := len(grid.Nodes)
	y := len(grid.Nodes[0])

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {

			count := len(grid.GetItemsAtLocation(i, j))
			msg := fmt.Sprintf("%v", count)
			ebitenutil.DebugPrintAt(
				e.canvas,
				msg,
				int(((float64(i)*grid.ChunkSize)+grid.ChunkSize/2)*e.scale),
				int(((float64(j)*grid.ChunkSize)+grid.ChunkSize/2)*e.scale),
			)
		}
	}
}

func SpriteKey(sheet, sprite string) string {
	return fmt.Sprintf("%s_%s", sheet, sprite)
}
