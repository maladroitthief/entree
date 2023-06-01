package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
)

const (
	SpriteSize = 16
)

func StaticTile(x, y float64, sheet, sprite string) canvas.Entity {
	return &backgroundEntity{
		size:         physics.Vector{X: SpriteSize, Y: SpriteSize},
		position:     physics.Vector{X: x, Y: y},
		offset:       physics.Vector{X: 0, Y: 0},
		scale:        canvas.DefaultScale,
		sheet:        sheet,
		sprite:       sprite,
		orientationX: canvas.Neutral,
		orientationY: canvas.South,
		input:        &backgroundInput{},
		physics:      &backgroundPhysics{},
		graphics:     &backgroundGraphics{},
	}
}

type backgroundInput struct{}

func (i *backgroundInput) Update(e canvas.Entity)                   {}
func (i *backgroundInput) Receive(e canvas.Entity, msg, val string) {}

type backgroundPhysics struct{}

func (p *backgroundPhysics) Update(e canvas.Entity, c *canvas.Canvas) {}
func (p *backgroundPhysics) Receive(e canvas.Entity, msg, val string) {}

type backgroundGraphics struct{}

func (g *backgroundGraphics) Update(e canvas.Entity)                   {}
func (g *backgroundGraphics) Receive(e canvas.Entity, msg, val string) {}
