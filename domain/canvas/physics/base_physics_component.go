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
	if e.DeltaX == 0 {
		e.VelocityX = 0
		return
	}

	if math.Signbit(e.DeltaX) != math.Signbit(e.VelocityX) {
		e.VelocityX = 0
	}

	e.VelocityX += e.DeltaX * e.Acceleration * e.Mass
	if e.VelocityX < 0 {
		e.X += int(math.Max(e.VelocityX, -e.MaxVelocity))
	} else {
		e.X += int(math.Min(e.VelocityX, e.MaxVelocity))
	}
}

func (p *BasePhysicsComponent) resolveY(e *canvas.Entity) {
	if e.DeltaY == 0 {
		e.VelocityY = 0
		return
	}

	if math.Signbit(e.DeltaY) != math.Signbit(e.VelocityY) {
		e.VelocityY = 0
	}

	e.VelocityY += e.DeltaY * e.Acceleration * e.Mass
	if e.VelocityY < 0 {
		e.Y += int(math.Max(e.VelocityY, -e.MaxVelocity))
	} else {
		e.Y += int(math.Min(e.VelocityY, e.MaxVelocity))
	}
}
