package attribute

import "github.com/maladroitthief/entree/common/data"

type Position struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	X float64
	Y float64
	Z float64
}

func NewPosition(x, y, z float64) Position {
	return Position{
		X: x,
		Y: y,
		Z: z,
	}
}
