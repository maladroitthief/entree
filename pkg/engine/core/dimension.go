package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

func (e *ECS) AddDimension(entity Entity, p attribute.Dimension) Entity {
	dimensionId := e.dimensionAllocator.Allocate()

	p.Id = dimensionId
	p.EntityId = entity.Id
	entity.DimensionId = dimensionId

	e.dimension = e.dimension.Set(dimensionId, p)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetDimension(entityId data.GenerationalIndex) (attribute.Dimension, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return attribute.Dimension{}, err
	}

	dimension := e.dimension.Get(entity.DimensionId)
	if !e.dimensionAllocator.IsLive(dimension.Id) {
		return dimension, ErrAttributeNotFound
	}

	return dimension, nil
}

func (e *ECS) GetAllDimensions() []attribute.Dimension {
	return e.dimension.GetAll(e.dimensionAllocator)
}

func (e *ECS) SetDimension(dimension attribute.Dimension) {
	e.dimension = e.dimension.Set(dimension.Id, dimension)
}
