package core

import "github.com/maladroitthief/caravan"

const (
	Neutral OrientationX = iota
	West
	East

	South OrientationY = iota
	North

	Moving state = 1 << iota
	Dodging
	Windup
	Attacking
	Cooldown

	DodgeDuration = 40
)

type (
	OrientationX int
	OrientationY int

	Condition struct {
		strength int
		duration int
		decay    int
	}

	state uint
	State struct {
		Id       caravan.GIDX
		EntityId caravan.GIDX

		state state

		OrientationX OrientationX
		OrientationY OrientationY

		DodgeCounter int

		SkillWindUpCounter   int
		SkillActiveCounter   int
		SkillCoolDownCounter int
	}
)

func (ecs *ECS) NewState() State {
	state := State{
		Id:           ecs.states.Allocate(),
		DodgeCounter: 0,
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

func (s State) SetIdle() State {
	return s.Unset(Moving | Dodging)
}

func (s State) Idling() bool {
	return !s.Check(Moving | Dodging)
}

func (s State) Set(state state) State {
	s.state |= state
	return s
}

func (s State) Unset(state state) State {
	s.state &= ^state
	return s
}

func (s State) Check(state state) bool {
	return s.state&state != 0
}
