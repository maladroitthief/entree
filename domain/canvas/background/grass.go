package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
)

func Grass(x, y int) *canvas.Entity {
	return StaticEntity(x, y, "test", "grass")
}
