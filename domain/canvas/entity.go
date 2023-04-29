package canvas

import (
	"math"

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
	DefaultAcceleration = 2
	DefaultDeceleration = 15
	DefaultMaxVelocity  = 10
	DefaultMass         = 10
)

type Entity struct {
	Size              collision.Vector
	Position          collision.Vector
	DeltaPosition     collision.Vector
	Velocity          collision.Vector
	Acceleration      float64
	MaxVelocity       float64
	Mass              float64
	Sheet             string
	Sprite            string
	SpriteSpeed       float64
	SpriteVariant     int
	SpriteMaxVariants int
	State             string
	StateCounter      int
	OrientationX      OrientationX
	OrientationY      OrientationY
	Input             InputComponent
	Physics           PhysicsComponent
	Graphics          GraphicsComponent
}

func (e *Entity) Update(c *Canvas) {
	e.Input.Update(e)
	e.Physics.Update(e, c)
	e.Graphics.Update(e)
}

func (e *Entity) VariantUpdate() {
	speed := float64(e.StateCounter) /
		(e.SpriteSpeed / float64(e.SpriteMaxVariants))
	e.SpriteVariant = int(speed)%e.SpriteMaxVariants + 1
}

func (e *Entity) LimitVelocity() {
	direction := collision.Vector{X: 1, Y: 1}

	if e.Velocity.X < 0 {
		direction.X = -1
	}

	if e.Velocity.Y < 0 {
		direction.Y = -1
	}

	if math.Abs(e.Velocity.X) > e.MaxVelocity {
		e.Velocity.X = e.MaxVelocity
	}

	if math.Abs(e.Velocity.Y) > e.MaxVelocity {
		e.Velocity.Y = e.MaxVelocity
	}

	e.Velocity = e.Velocity.ScaleXY(direction.X, direction.Y)

	// Normalize the diagonals
	m := e.Velocity.Magnitude()
	if m > e.MaxVelocity {
		e.Velocity = e.Velocity.Scale(e.MaxVelocity / m)
	}
}

func (e *Entity) Reset() {
	e.State = "idle"
	dx, dy := e.DeltaPosition.X, e.DeltaPosition.Y
	if dx == 0 && dy != 0 {
		e.OrientationX = Neutral
	}
	e.DeltaPosition = collision.Vector{X: 0, Y: 0}
}
