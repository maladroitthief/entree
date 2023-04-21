package environment

import "github.com/maladroitthief/entree/domain/canvas"

const (
	SpriteSize = 16
)

func StaticEntity(x, y float64, sheet, sprite string) *canvas.Entity {
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
		Input:             &environmentInput{},
		Physics:           &environmentPhysics{},
		Graphics:          &environmentGraphics{},
	}
}

type environmentGraphics struct {
}

func (g *environmentGraphics) Update(e *canvas.Entity) {
}

type environmentPhysics struct {
}

func (g *environmentPhysics) Update(e *canvas.Entity, c *canvas.Canvas) {
}

type environmentInput struct {
}

func (g *environmentInput) Update(e *canvas.Entity) {
}
