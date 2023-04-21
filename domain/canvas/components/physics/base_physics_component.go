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
	p.resolveX(e)
	p.resolveY(e)
}

func (p *BasePhysicsComponent) resolveX(e *canvas.Entity) {
	if e.DeltaPosition.X == 0 {
		e.VelocityX = 0
		return
	}

	if math.Signbit(e.DeltaPosition.X) != math.Signbit(e.VelocityX) {
		e.VelocityX = 0
	}

	e.VelocityX += e.DeltaPosition.X * e.Acceleration * e.Mass
	if e.VelocityX < 0 {
		e.Position.X += math.Max(e.VelocityX, -e.MaxVelocity)
	} else {
		e.Position.X += math.Min(e.VelocityX, e.MaxVelocity)
	}
}

func (p *BasePhysicsComponent) resolveY(e *canvas.Entity) {
	if e.DeltaPosition.Y == 0 {
		e.VelocityY = 0
		return
	}

	if math.Signbit(e.DeltaPosition.Y) != math.Signbit(e.VelocityY) {
		e.VelocityY = 0
	}

	e.VelocityY += e.DeltaPosition.Y * e.Acceleration * e.Mass
	if e.VelocityY < 0 {
		e.Position.Y += math.Max(e.VelocityY, -e.MaxVelocity)
	} else {
		e.Position.Y += math.Min(e.VelocityY, e.MaxVelocity)
	}
}
