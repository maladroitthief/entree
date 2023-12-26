package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

func (e *ECS) AddBehavior(entity Entity, b attribute.Behavior) Entity {
	behaviorId := e.behaviorAllocator.Allocate()

	b.Id = behaviorId
	b.EntityId = entity.Id
	entity.BehaviorId = behaviorId

	e.behavior = e.behavior.Set(behaviorId, b)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetBehavior(entityId data.GenerationalIndex) (attribute.Behavior, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return attribute.Behavior{}, err
	}

	behavior := e.behavior.Get(entity.BehaviorId)
	if !e.behaviorAllocator.IsLive(behavior.Id) {
		return behavior, ErrAttributeNotFound
	}

	return behavior, nil
}

func (e *ECS) GetAllBehavior() []attribute.Behavior {
	return e.behavior.GetAll(e.behaviorAllocator)
}

func (e *ECS) SetBehavior(behavior attribute.Behavior) {
	e.behavior = e.behavior.Set(behavior.Id, behavior)
}
