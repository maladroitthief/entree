package player

import (
	"math"
	"strconv"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
)

type PlayerPhysicsComponent struct {
	DeltaPosition physics.Vector
	Velocity      physics.Vector
	Acceleration  float64
	MaxVelocity   float64
	Mass          float64
}

func NewPlayerPhysicsComponent() *PlayerPhysicsComponent {
	bpc := &PlayerPhysicsComponent{
		Acceleration: canvas.DefaultAcceleration,
		MaxVelocity:  canvas.DefaultMaxVelocity,
		Mass:         canvas.DefaultMass,
	}

	return bpc
}

func (p *PlayerPhysicsComponent) Update(e canvas.Entity, c *canvas.Canvas) {
	e.IncrementStateCounter()

	// BoundsHandling

	// Position Handling
	p.resolveVelocity()
	p.resolvePosition(c, e)
}

func (p *PlayerPhysicsComponent) Receive(e canvas.Entity, msg, val string) {
	switch msg {
	case "Reset":
		if p.DeltaPosition.X == 0 && p.DeltaPosition.Y != 0 {
			e.SetOrientationX(canvas.Neutral)
		}
		p.DeltaPosition.X, p.DeltaPosition.Y = 0, 0
	case "DeltaPositionX":
		p.DeltaPosition.X, _ = strconv.ParseFloat(val, 64)
	case "DeltaPositionY":
		p.DeltaPosition.Y, _ = strconv.ParseFloat(val, 64)
	}

}

func (p *PlayerPhysicsComponent) resolveVelocity() {
	// reset velocity if delta position has stopped
	p.Velocity = p.Velocity.ScaleXY(p.DeltaPosition.X, p.DeltaPosition.Y)

	// reset velocity if we are changing directions
	if math.Signbit(p.DeltaPosition.X) != math.Signbit(p.Velocity.X) {
		p.Velocity.X = 0
	}

	if math.Signbit(p.DeltaPosition.Y) != math.Signbit(p.Velocity.Y) {
		p.Velocity.Y = 0
	}

	p.Velocity = p.Velocity.Add(p.DeltaPosition.Scale(p.Acceleration * p.Mass))
	p.limitVelocity()
}

func (p *PlayerPhysicsComponent) resolvePosition(c *canvas.Canvas, e canvas.Entity) {
	newPosition := e.Position().Add(p.Velocity)
	newBounds := physics.Bounds(newPosition, e.Size())
	collisions := c.Collisions(e, newBounds)

	for _, oob := range c.Bounds() {
		newPosition = canvas.CollisionVector(e, oob, p.DeltaPosition, newPosition)
	}

	if len(collisions) == 0 {
		e.SetPosition(newPosition)
		return
	}

	for _, ce := range collisions {
		newPosition = canvas.CollisionVector(e, ce.Bounds(), p.DeltaPosition, newPosition)
	}

	e.SetPosition(newPosition)
}

func (p *PlayerPhysicsComponent) limitVelocity() {
	direction := physics.Vector{X: 1, Y: 1}

	if p.Velocity.X < 0 {
		direction.X = -1
	}

	if p.Velocity.Y < 0 {
		direction.Y = -1
	}

	if math.Abs(p.Velocity.X) > p.MaxVelocity {
		p.Velocity.X = p.MaxVelocity
	}

	if math.Abs(p.Velocity.Y) > p.MaxVelocity {
		p.Velocity.Y = p.MaxVelocity
	}

	p.Velocity = p.Velocity.ScaleXY(direction.X, direction.Y)

	// Normalize the diagonals
	m := p.Velocity.Magnitude()
	if m > p.MaxVelocity {
		p.Velocity = p.Velocity.Scale(p.MaxVelocity / m)
	}
}
