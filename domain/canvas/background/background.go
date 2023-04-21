package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
)

const (
	SpriteSize = 16
)

func StaticTile(x, y float64, sheet, sprite string) *canvas.Entity {
	return &canvas.Entity{
		Size:              canvas.Size{X: SpriteSize, Y: SpriteSize},
		Position:          canvas.Position{X: x, Y: y},
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
