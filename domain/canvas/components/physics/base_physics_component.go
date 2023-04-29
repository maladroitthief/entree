package physics

import (
	"math"

	"github.com/maladroitthief/entree/domain/canvas"
)

type BasePhysicsComponent struct {
}

func NewBasePhysicsComponent() *BasePhysicsComponent {
	bpc := &BasePhysicsComponent{}

	return bpc
}

func (p *BasePhysicsComponent) Update(e *canvas.Entity, c *canvas.Canvas) {
	e.StateCounter++

	// Position Handling
	p.resolveVelocity(e)
	p.resolvePosition(e)
}

func (p *BasePhysicsComponent) resolveVelocity(e *canvas.Entity) {
	// reset velocity if delta position has stopped
	e.Velocity = e.Velocity.ScaleXY(e.DeltaPosition.X, e.DeltaPosition.Y)

	// reset velocity if we are changing directions
	if math.Signbit(e.DeltaPosition.X) != math.Signbit(e.Velocity.X) {
		e.Velocity.X = 0
	}

	if math.Signbit(e.DeltaPosition.Y) != math.Signbit(e.Velocity.Y) {
		e.Velocity.Y = 0
	}

	e.Velocity = e.Velocity.Add(e.DeltaPosition.Scale(e.Acceleration * e.Mass))
	e.LimitVelocity()
}

func (p *BasePhysicsComponent) resolvePosition(e *canvas.Entity) {
	e.Position = e.Position.Add(e.Velocity)
}
