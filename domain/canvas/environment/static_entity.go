package environment

import (
	"github.com/maladroitthief/entree/domain/canvas"
)

func StaticEntity(x, y float64, sheet, sprite string) canvas.Entity {
  e := canvas.NewEntity()
  e.SetX(x)
  e.SetY(y)
  e.SetSheet(sheet)
  e.SetSprite(sprite)
  e.SetInputComponent(&environmentInput{})
  e.SetPhysicsComponent(&environmentPhysics{})
  e.SetGraphicsComponent(&environmentGraphics{})
  
  return e
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
