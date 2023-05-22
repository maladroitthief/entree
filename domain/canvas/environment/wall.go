package environment

import "github.com/maladroitthief/entree/domain/canvas"

const (
	WallSize = 15
)

func Wall(x, y float64) canvas.Entity {
	return StaticEntity(x, y, WallSize, WallSize, "test", "wall")
}

func WallFactory() []*canvas.Entity {
	return nil
}
