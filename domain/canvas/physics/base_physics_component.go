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
  p.ResolveX(e)
  p.ResolveY(e)
}

func (p *BasePhysicsComponent) ResolveX(e *canvas.Entity) {
	if e.DeltaX == 0 {
		return
	}

	x := e.DeltaX * e.Acceleration * e.Mass
	e.X += int(math.Min(float64(x), float64(e.MaxVelocity)))
}

func (p *BasePhysicsComponent) ResolveY(e *canvas.Entity) {
	if e.DeltaY == 0 {
		return
	}

	y := e.DeltaY * e.Acceleration * e.Mass
	e.Y += int(math.Min(float64(y), float64(e.MaxVelocity)))
}
