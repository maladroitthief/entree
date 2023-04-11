package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
)

type backgroundGraphics struct {
}

func (g *backgroundGraphics) Update(e *canvas.Entity) {
}

type backgroundPhysics struct {
}

func (g *backgroundPhysics) Update(e *canvas.Entity, c *canvas.Canvas) {
}

type backgroundInput struct {
}

func (g *backgroundInput) Update(e *canvas.Entity) {
}
