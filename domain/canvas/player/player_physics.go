package player

import (
	"math"
	"strconv"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics/collision"
)

type PlayerPhysicsComponent struct {
	DeltaPosition collision.Vector
	Velocity      collision.Vector
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

func (p *PlayerPhysicsComponent) Update(e *canvas.Entity, c *canvas.Canvas) {
	e.StateCounter++

	// Position Handling
	p.resolveVelocity()
	p.resolvePosition(c, e)
}

func (p *PlayerPhysicsComponent) Receive(e *canvas.Entity, msg, val string) {
	switch msg {
	case "Reset":
		if p.DeltaPosition.X == 0 && p.DeltaPosition.Y != 0 {
			e.OrientationX = canvas.Neutral
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

func (p *PlayerPhysicsComponent) resolvePosition(c *canvas.Canvas, e *canvas.Entity) {
	newPosition := e.Position.Add(p.Velocity)
	newBounds := collision.NewRectangle(
		newPosition.X,
		newPosition.Y,
		newPosition.X+e.Size.X,
		newPosition.Y+e.Size.Y,
	)
	collisions := c.Collisions(e, newBounds)

	if len(collisions) == 0 {
		e.Position = newPosition
		return
	}

	for _, ce := range collisions {
		// Set the X position
		if p.DeltaPosition.X != 0 {
			newBounds = collision.NewRectangle(
				newPosition.X,
				e.Bounds.MinPoint.Y,
				newPosition.X+e.Size.X,
				e.Bounds.MaxPoint.Y,
			)
			if ce.Bounds.Intersects(newBounds) {
				if p.DeltaPosition.X > 0 && newBounds.MaxPoint.X > ce.Bounds.MinPoint.X {
					newPosition.X = ce.Bounds.MinPoint.X - e.Size.X - 1
				} else if p.DeltaPosition.X < 0 && newBounds.MinPoint.X < ce.Bounds.MaxPoint.X {
					newPosition.X = ce.Bounds.MaxPoint.X + 1
				}
			}
		}

		// Set the Y position
		if p.DeltaPosition.Y != 0 {
			newBounds = collision.NewRectangle(
				e.Bounds.MinPoint.X,
				newPosition.Y,
				e.Bounds.MaxPoint.X,
				newPosition.Y+e.Size.Y,
			)
			if ce.Bounds.Intersects(newBounds) {
				if p.DeltaPosition.Y > 0 && newBounds.MaxPoint.Y > ce.Bounds.MinPoint.Y {
					newPosition.Y = ce.Bounds.MinPoint.Y - e.Size.Y - 1
				} else if p.DeltaPosition.Y < 0 && newBounds.MinPoint.Y < ce.Bounds.MaxPoint.Y {
					newPosition.Y = ce.Bounds.MaxPoint.Y + 1
				}
			}
		}
	}

	e.Position = newPosition
}

func (p *PlayerPhysicsComponent) limitVelocity() {
	direction := collision.Vector{X: 1, Y: 1}

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
