package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics/collision"
)

const (
	SpriteSize = 16
)

func StaticTile(x, y float64, sheet, sprite string) *canvas.Entity {
	return &canvas.Entity{
		Size:         collision.Vector{X: SpriteSize, Y: SpriteSize},
		Position:     collision.Vector{X: x, Y: y},
		Sheet:        sheet,
		Sprite:       sprite,
		OrientationX: canvas.Neutral,
		OrientationY: canvas.South,
		Input:        &backgroundInput{},
		Physics:      &backgroundPhysics{},
		Graphics:     &backgroundGraphics{},
	}
}

type backgroundInput struct{}

func (i *backgroundInput) Update(e *canvas.Entity)                   {}
func (i *backgroundInput) Receive(e *canvas.Entity, msg, val string) {}

type backgroundPhysics struct{}

func (p *backgroundPhysics) Update(e *canvas.Entity, c *canvas.Canvas) {}
func (p *backgroundPhysics) Receive(e *canvas.Entity, msg, val string) {}

type backgroundGraphics struct{}

func (g *backgroundGraphics) Update(e *canvas.Entity)                   {}
func (g *backgroundGraphics) Receive(e *canvas.Entity, msg, val string) {}
