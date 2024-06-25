package core

import "github.com/maladroitthief/caravan"

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
		Id       caravan.GIDX
		EntityId caravan.GIDX

		State        string
		Counter      int
		OrientationX OrientationX
		OrientationY OrientationY
	}
)

func (ecs *ECS) NewState() State {
	state := State{
		Id:           ecs.states.Allocate(),
		State:        Idling,
		Counter:      0,
		OrientationX: Neutral,
		OrientationY: South,
	}
	ecs.states.Set(state.Id, state)

	return state
}

func (ecs *ECS) BindState(entity Entity, state State) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.stateMu.Lock()
	defer ecs.stateMu.Unlock()

	state.EntityId = entity.Id
	entity.StateId = state.Id

	ecs.states.Set(state.Id, state)
	ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetState(entity Entity) (State, error) {
	return ecs.GetStateById(entity.StateId)
}
func (ecs *ECS) GetStateById(id caravan.GIDX) (State, error) {
	ecs.stateMu.RLock()
	defer ecs.stateMu.RUnlock()

	state := ecs.states.Get(id)
	if !ecs.states.IsLive(state.Id) {
		return state, ErrAttributeNotFound
	}

	return state, nil
}

func (ecs *ECS) GetAllStates() []State {
	ecs.stateMu.RLock()
	defer ecs.stateMu.RUnlock()

	return ecs.states.GetAll()
}

func (ecs *ECS) SetState(state State) {
	ecs.stateMu.Lock()
	defer ecs.stateMu.Unlock()

	ecs.states.Set(state.Id, state)
}

func (ecs *ECS) StateActive(state State) bool {
	return ecs.states.IsLive(state.Id)
}
