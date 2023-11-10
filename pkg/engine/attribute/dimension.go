package attribute

import "github.com/maladroitthief/entree/common/data"

type Dimension struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Size   data.Vector
	Scale  float64
	Offset data.Vector
	Bounds data.Polygon
}

func NewDimension(position data.Vector, size data.Vector) Dimension {
	return Dimension{
		Size:   size,
		Scale:  1,
		Offset: data.Vector{X: 0, Y: 0},
		Bounds: data.NewRectangle(position, size.X, size.Y).ToPolygon(),
	}
}
