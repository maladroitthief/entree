package environment

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics/collision"
)

const (
	SpriteSize = 16
)

func StaticEntity(x, y float64, sheet, sprite string) *canvas.Entity {
	return &canvas.Entity{
		Size:         collision.Vector{X: SpriteSize, Y: SpriteSize},
		Position:     collision.Vector{X: x, Y: y},
		Sheet:        sheet,
		Sprite:       sprite,
		OrientationX: canvas.Neutral,
		OrientationY: canvas.South,
		Input:        &environmentInput{},
		Physics:      &environmentPhysics{},
		Graphics:     &environmentGraphics{},
	}
}

type environmentInput struct{}

func (i *environmentInput) Update(e *canvas.Entity)                   {}
func (i *environmentInput) Receive(e *canvas.Entity, msg, val string) {}

type environmentPhysics struct{}

func (p *environmentPhysics) Update(e *canvas.Entity, c *canvas.Canvas) {}
func (p *environmentPhysics) Receive(e *canvas.Entity, msg, val string) {}

type environmentGraphics struct{}

func (g *environmentGraphics) Update(e *canvas.Entity)                   {}
func (g *environmentGraphics) Receive(e *canvas.Entity, msg, val string) {}
