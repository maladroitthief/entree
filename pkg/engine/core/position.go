package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

func (e *ECS) AddPosition(entity Entity, p attribute.Position) Entity {
	positionId := e.positionAllocator.Allocate()

	p.Id = positionId
	p.EntityId = entity.Id
	entity.PositionId = positionId

	e.position = e.position.Set(positionId, p)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetPosition(entityId data.GenerationalIndex) (attribute.Position, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return attribute.Position{}, err
	}

	position := e.position.Get(entity.PositionId)
	if !e.positionAllocator.IsLive(position.Id) {
		return position, ErrAttributeNotFound
	}

	return position, nil
}

func (e *ECS) GetAllPosition() []attribute.Position {
	return e.position.GetAll(e.positionAllocator)
}

func (e *ECS) SetPosition(position attribute.Position) {
	e.position = e.position.Set(position.Id, position)
}
