package environment

import (
	"github.com/maladroitthief/entree/domain/canvas"
)

func Wall(x, y float64) canvas.Entity {
	wall := StaticEntity(x, y, "test", "wall")

	return wall
}
