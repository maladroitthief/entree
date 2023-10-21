package attribute

import "github.com/maladroitthief/entree/common/data"

type Dimension struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Size   data.Vector
	Scale  float64
	Offset data.Vector
	Bounds data.Rectangle
}

func NewDimension(position data.Vector, size data.Vector) Dimension {
	return Dimension{
		Size:   size,
		Scale:  1,
		Offset: data.Vector{X: 0, Y: 0},
    // TODO support other bound types
		Bounds: data.Bounds(position, size),
	}
}
