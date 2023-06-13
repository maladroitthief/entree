package environment

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
)

const (
	WallSize = 16
)

func Wall(x, y float64) canvas.Entity {
	return &environmentEntity{
		size:         physics.Vector{X: WallSize, Y: WallSize},
		position:     physics.Vector{X: x, Y: y},
		scale:        canvas.DefaultScale,
		sheet:        "test",
		sprite:       "wall",
		orientationX: canvas.Neutral,
		orientationY: canvas.South,
		input:        &environmentInput{},
		physics:      &environmentPhysics{},
		graphics:     &environmentGraphics{},
	}
}

func WallFactory() []*canvas.Entity {
	return nil
}
