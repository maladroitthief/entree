package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

func (e *ECS) AddAI(entity Entity, a attribute.AI) Entity {
	aiId := e.aiAllocator.Allocate()

	a.Id = aiId
	a.EntityId = entity.Id
	entity.AIId = aiId

	e.ai = e.ai.Set(aiId, a)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetAI(entityId data.GenerationalIndex) (attribute.AI, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return attribute.AI{}, err
	}

	ai := e.ai.Get(entity.AIId)
	if !e.aiAllocator.IsLive(ai.Id) {
		return ai, ErrAttributeNotFound
	}

	return ai, nil
}

func (e *ECS) GetAllAI() []attribute.AI {
	return e.ai.GetAll(e.aiAllocator)
}

func (e *ECS) SetAI(ai attribute.AI) {
	e.ai = e.ai.Set(ai.Id, ai)
}
