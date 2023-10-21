package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

func (e *ECS) AddMovement(entity Entity, p attribute.Movement) Entity {
	movementId := e.movementAllocator.Allocate()

	p.Id = movementId
	p.EntityId = entity.Id
	entity.MovementId = movementId

	e.movement = e.movement.Set(movementId, p)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetMovement(entityId data.GenerationalIndex) (attribute.Movement, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return attribute.Movement{}, err
	}

	movement := e.movement.Get(entity.MovementId)
	if !e.movementAllocator.IsLive(movement.Id) {
		return movement, ErrAttributeNotFound
	}

	return movement, nil
}

func (e *ECS) GetAllMovement() []attribute.Movement {
	return e.movement.GetAll(e.movementAllocator)
}

func (e *ECS) SetMovement(movement attribute.Movement) {
	e.movement = e.movement.Set(movement.Id, movement)
}
