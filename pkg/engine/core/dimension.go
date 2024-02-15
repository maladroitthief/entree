package core

import (
	"github.com/maladroitthief/entree/common/data"
)

type Dimension struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Size    data.Vector
	Scale   float64
	Offset  data.Vector
	Polygon data.Polygon
}

func (e *ECS) NewDimension(position data.Vector, size data.Vector) Dimension {
	dimension := Dimension{
		Id:      e.dimensionAllocator.Allocate(),
		Size:    size,
		Scale:   1,
		Offset:  data.Vector{X: 0, Y: 0},
		Polygon: data.NewRectangle(position, size.X, size.Y).ToPolygon(),
	}
	e.dimensions.Set(dimension.Id, dimension)

	return dimension
}

func (d Dimension) Bounds() data.Rectangle {
	return d.Polygon.Bounds.Scale(d.Scale)
}

func (e *ECS) BindDimension(entity Entity, dimension Dimension) Entity {
	dimension.EntityId = entity.Id
	entity.DimensionId = dimension.Id

	e.dimensions = e.dimensions.Set(dimension.Id, dimension)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetDimension(entity Entity) (Dimension, error) {
	return e.GetDimensionById(entity.DimensionId)
}

func (e *ECS) GetDimensionById(id data.GenerationalIndex) (Dimension, error) {
	dimension := e.dimensions.Get(id)
	if !e.dimensionAllocator.IsLive(dimension.Id) {
		return dimension, ErrAttributeNotFound
	}

	return dimension, nil
}

func (e *ECS) GetAllDimensions() []Dimension {
	return e.dimensions.GetAll(e.dimensionAllocator)
}

func (e *ECS) SetDimension(dimension Dimension) {
	e.dimensions = e.dimensions.Set(dimension.Id, dimension)
}
