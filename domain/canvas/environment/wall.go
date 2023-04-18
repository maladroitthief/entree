package environment

import "github.com/maladroitthief/entree/domain/canvas"

func Wall(x, y float64) *canvas.Entity {
	return StaticEntity(x, y, "test", "wall")
}

func WallFactory() []*canvas.Entity {
	return nil
}
