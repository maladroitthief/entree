package core

import (
	"github.com/maladroitthief/entree/common/data"
)

const (
	Neutral OrientationX = iota
	West
	East
	South OrientationY = iota
	North

	Idling  = "idle"
	Moving  = "move"
	Dodging = "dodge"

	DodgeDuration = 40
)

type (
	OrientationX int
	OrientationY int

	State struct {
		Id       data.GenerationalIndex
		EntityId data.GenerationalIndex

		State        string
		Counter      int
		OrientationX OrientationX
		OrientationY OrientationY
	}
)

func (e *ECS) NewState() State {
	state := State{
		Id:           e.stateAllocator.Allocate(),
		State:        Idling,
		Counter:      0,
		OrientationX: Neutral,
		OrientationY: South,
	}
	e.states.Set(state.Id, state)

	return state
}

func (e *ECS) BindState(entity Entity, state State) Entity {
	state.EntityId = entity.Id
	entity.StateId = state.Id

	e.states = e.states.Set(state.Id, state)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetState(entity Entity) (State, error) {
	return e.GetStateById(entity.StateId)
}
func (e *ECS) GetStateById(id data.GenerationalIndex) (State, error) {
	state := e.states.Get(id)
	if !e.stateAllocator.IsLive(state.Id) {
		return state, ErrAttributeNotFound
	}

	return state, nil
}

func (e *ECS) GetAllStates() []State {
	return e.states.GetAll(e.stateAllocator)
}

func (e *ECS) SetState(state State) {
	e.states = e.states.Set(state.Id, state)
}
