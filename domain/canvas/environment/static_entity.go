package environment

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics/collision"
)

const (
	SpriteSize = 16
)

func StaticEntity(x, y float64, sheet, sprite string) canvas.Entity {
	return &environmentEntity{
		size:         collision.Vector{X: SpriteSize, Y: SpriteSize},
		position:     collision.Vector{X: x, Y: y},
		sheet:        sheet,
		sprite:       sprite,
		orientationX: canvas.Neutral,
		orientationY: canvas.South,
		input:        &environmentInput{},
		physics:      &environmentPhysics{},
		graphics:     &environmentGraphics{},
	}
}

type environmentInput struct{}

func (i *environmentInput) Update(e canvas.Entity)                   {}
func (i *environmentInput) Receive(e canvas.Entity, msg, val string) {}

type environmentPhysics struct{}

func (p *environmentPhysics) Update(e canvas.Entity, c *canvas.Canvas) {}
func (p *environmentPhysics) Receive(e canvas.Entity, msg, val string) {}

type environmentGraphics struct{}

func (g *environmentGraphics) Update(e canvas.Entity)                   {}
func (g *environmentGraphics) Receive(e canvas.Entity, msg, val string) {}
