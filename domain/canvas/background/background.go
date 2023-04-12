package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
)

const (
  SpriteSize = 16
)

func StaticEntity(x, y int, sheet, sprite string) *canvas.Entity {
	return &canvas.Entity{
		Width:             SpriteSize,
		Height:            SpriteSize,
		X:                 x,
		Y:                 y,
		Sheet:             sheet,
		Sprite:            sprite,
		SpriteSpeed:       40,
		SpriteVariant:     1,
		SpriteMaxVariants: 1,
		OrientationX:      canvas.Neutral,
		OrientationY:      canvas.South,
		Input:             &backgroundInput{},
		Physics:           &backgroundPhysics{},
		Graphics:          &backgroundGraphics{},
	}
}

type backgroundGraphics struct {
}

func (g *backgroundGraphics) Update(e *canvas.Entity) {
}

type backgroundPhysics struct {
}

func (g *backgroundPhysics) Update(e *canvas.Entity, c *canvas.Canvas) {
}

type backgroundInput struct {
}

func (g *backgroundInput) Update(e *canvas.Entity) {
}
