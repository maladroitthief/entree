package canvas

import (
	"github.com/maladroitthief/entree/domain/physics/collision"
)

type OrientationX int
type OrientationY int

const (
	Neutral OrientationX = iota
	West
	East
	South OrientationY = iota
	North
	DefaultAcceleration = 1.5
	DefaultMaxVelocity  = 5
	DefaultMass         = 10
	DefaultSpriteSpeed  = 40
)

type Entity struct {
	Position     collision.Vector
	Size         collision.Vector
	Bounds       collision.Rectangle
	Sheet        string
	Sprite       string
	State        string
	StateCounter int
	OrientationX OrientationX
	OrientationY OrientationY
	Components   []Component
	Input        InputComponent
	Physics      PhysicsComponent
	Graphics     GraphicsComponent
}

func (e *Entity) Update(c *Canvas) {
	e.setBounds()
	e.Input.Update(e)
	e.Physics.Update(e, c)
	e.Graphics.Update(e)
}

func (e *Entity) Send(msg, val string) {
	for _, c := range e.Components {
		c.Receive(e, msg, val)
	}
}

func (e *Entity) setBounds() {
	e.Bounds = collision.NewRectangle(
		e.Position.X,
		e.Position.Y,
		e.Position.X+e.Size.X,
		e.Position.Y+e.Size.Y,
	)
}

func (e *Entity) Reset() {
	e.State = "idle"
	e.Send("Reset", "")
}
