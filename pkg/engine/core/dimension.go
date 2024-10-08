package core

import (
	"github.com/maladroitthief/caravan"
	"github.com/maladroitthief/mosaic"
)

type Dimension struct {
	Id       caravan.GIDX
	EntityId caravan.GIDX

	Size    mosaic.Vector
	Scale   float64
	Offset  mosaic.Vector
	Polygon mosaic.Polygon
}

func (ecs *ECS) NewDimension(position mosaic.Vector, size mosaic.Vector) Dimension {
	dimension := Dimension{
		Id:      ecs.dimensionAllocator.Allocate(),
		Size:    size,
		Scale:   1,
		Offset:  mosaic.Vector{X: 0, Y: 0},
		Polygon: mosaic.NewRectangle(position, size.X, size.Y).ToPolygon(),
	}
	ecs.dimensions.Set(dimension.Id, dimension)

	return dimension
}

func (d Dimension) Bounds() mosaic.Rectangle {
	return d.Polygon.Bounds.Scale(d.Scale)
}

func (ecs *ECS) BindDimension(entity Entity, dimension Dimension) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.dimensionMu.Lock()
	defer ecs.dimensionMu.Unlock()

	dimension.EntityId = entity.Id
	entity.DimensionId = dimension.Id

	ecs.dimensions = ecs.dimensions.Set(dimension.Id, dimension)
	ecs.entities = ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetDimension(entity Entity) (Dimension, error) {
	return ecs.GetDimensionById(entity.DimensionId)
}

func (ecs *ECS) GetDimensionById(id caravan.GIDX) (Dimension, error) {
	ecs.dimensionMu.RLock()
	defer ecs.dimensionMu.RUnlock()

	dimension := ecs.dimensions.Get(id)
	if !ecs.dimensionAllocator.IsLive(dimension.Id) {
		return dimension, ErrAttributeNotFound
	}

	return dimension, nil
}

func (ecs *ECS) GetAllDimensions() []Dimension {
	ecs.dimensionMu.RLock()
	defer ecs.dimensionMu.RUnlock()

	return ecs.dimensions.GetAll(ecs.dimensionAllocator)
}

func (ecs *ECS) SetDimension(dimension Dimension) {
	ecs.dimensionMu.Lock()
	defer ecs.dimensionMu.Unlock()

	ecs.dimensions = ecs.dimensions.Set(dimension.Id, dimension)
}
