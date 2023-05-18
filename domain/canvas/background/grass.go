package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
)

func Grass(x, y float64) canvas.Entity {
	return StaticTile(x, y, "test", "grass")
}
