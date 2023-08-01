package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

func (e *ECS) AddState(entity Entity, s attribute.State) Entity {
	stateId := e.stateAllocator.Allocate()

	s.Id = stateId
	s.EntityId = entity.Id
	entity.StateId = stateId

	e.state = e.state.Set(stateId, s)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetState(entityId data.GenerationalIndex) (attribute.State, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return attribute.State{}, err
	}

	state := e.state.Get(entity.StateId)
	if !e.stateAllocator.IsLive(state.Id) {
		return state, ErrAttributeNotFound
	}

	return state, nil
}

func (e *ECS) GetAllStates() []attribute.State {
	return e.state.GetAll(e.stateAllocator)
}

func (e *ECS) SetState(state attribute.State) {
	e.state = e.state.Set(state.Id, state)
}
