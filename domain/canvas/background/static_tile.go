package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
)


func StaticTile(x, y float64, sheet, sprite string) canvas.Entity {
  e := canvas.NewEntity()
  e.SetX(x)
  e.SetY(y)
  e.SetSheet(sheet)
  e.SetSprite(sprite)
  e.SetInputComponent(&backgroundInput{})
  e.SetPhysicsComponent(&backgroundPhysics{})
  e.SetGraphicsComponent(&backgroundGraphics{})
  
  return e
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
