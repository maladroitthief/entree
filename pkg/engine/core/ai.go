package core

import (
	"github.com/maladroitthief/entree/common/data"
)

type BehaviorType int

const (
	None BehaviorType = iota
	Player
	Computer
)

type AI struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	BehaviorType BehaviorType

	RootBehavior   data.GenerationalIndex
	ActiveBehavior data.GenerationalIndex
	ActiveSequence bool
}

func (e *ECS) NewAI(b BehaviorType) AI {
	ai := AI{
		Id:           e.aiAllocator.Allocate(),
		BehaviorType: b,
	}
	e.ai.Set(ai.Id, ai)

	return ai
}

func (e *ECS) BindAI(entity Entity, ai AI) Entity {
	ai.EntityId = entity.Id
	entity.AIId = ai.Id

	e.ai = e.ai.Set(ai.Id, ai)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetAI(entityId data.GenerationalIndex) (AI, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return AI{}, err
	}

	ai := e.ai.Get(entity.AIId)
	if !e.aiAllocator.IsLive(ai.Id) {
		return ai, ErrAttributeNotFound
	}

	return ai, nil
}

func (e *ECS) GetAllAI() []AI {
	return e.ai.GetAll(e.aiAllocator)
}

func (e *ECS) SetAI(ai AI) {
	e.ai = e.ai.Set(ai.Id, ai)
}
