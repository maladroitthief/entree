package core

import (
	"math/rand"

	"github.com/maladroitthief/entree/common/data"
)

type Status int
type Archetype int

const (
	RUNNING Status = iota
	SUCCESS
	FAILURE

	COMPOSITE_NODE Archetype = iota
	DECORATOR_NODE
	LEAF_NODE
)

type Behavior struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Status    Status
	Archetype Archetype
	Parent    data.GenerationalIndex
	Children  []data.GenerationalIndex

	Tick func(*ECS, Behavior) (Status, error)
}

func AddChild(parent, child Behavior) (Behavior, Behavior) {
	parent.Children = append(parent.Children, child.Id)
	child.Parent = parent.Id

	return parent, child
}
func (e *ECS) Root() Behavior {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: COMPOSITE_NODE,
		Children:  make([]data.GenerationalIndex, 0),
	}

	b.Tick = func(e *ECS, b Behavior) (Status, error) {
		b.Status = RUNNING
		child := e.behaviors.Get(b.Children[0])
		if !e.behaviorAllocator.IsLive(child.Id) {
			b.Status = FAILURE
			e.SetBehavior(b)
			return b.Status, ErrAttributeNotFound
		}
		status, err := child.Tick(e, b)
		if err != nil {
			b.Status = FAILURE
			e.SetBehavior(b)
			return b.Status, err
		}
		b.Status = status
		e.SetBehavior(b)

		return status, nil
	}

	e.behaviors.Set(b.Id, b)

	return b
}

func (e *ECS) RandomSequence(parent Behavior) Behavior {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: COMPOSITE_NODE,
		Parent:    parent.Id,
		Children:  make([]data.GenerationalIndex, 0),
	}

	b.Tick = func(e *ECS, b Behavior) (Status, error) {
		b.Status = RUNNING
		if len(b.Children) <= 0 {
			b.Status = FAILURE
			e.SetBehavior(b)
			return b.Status, ErrAttributeNotFound
		}
		child := e.behaviors.Get(b.Children[rand.Intn(len(b.Children))])
		if !e.behaviorAllocator.IsLive(child.Id) {
			b.Status = FAILURE
			e.SetBehavior(b)
			return b.Status, ErrAttributeNotFound
		}

		status, err := child.Tick(e, b)
		if err != nil {
			b.Status = FAILURE
			e.SetBehavior(b)
			return b.Status, err
		}

		b.Status = status
		e.SetBehavior(b)
		return b.Status, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.behaviors.Set(b.Id, b)

	return b
}

func (e *ECS) Inverter(parent Behavior) Behavior {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: DECORATOR_NODE,
		Parent:    parent.Id,
		Children:  make([]data.GenerationalIndex, 1),
	}

	b.Tick = func(e *ECS, b Behavior) (Status, error) {
		child := e.behaviors.Get(b.Children[0])
		if !e.behaviorAllocator.IsLive(child.Id) {
			return FAILURE, ErrAttributeNotFound
		}
		status, err := child.Tick(e, b)
		if err != nil {
			return FAILURE, err
		}

		switch status {
		case SUCCESS:
			return FAILURE, nil
		case FAILURE:
			return SUCCESS, nil
		}

		return status, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.behaviors.Set(b.Id, b)

	return b
}

func (e *ECS) MovingUp(parent Behavior) Behavior {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: LEAF_NODE,
		Parent:    parent.Id,
	}

	b.Tick = func(e *ECS, b Behavior) (Status, error) {
		state, err := e.GetState(b.EntityId)
		if err == nil {
			state.State = Moving
			state.OrientationY = North
			e.SetState(state)
		}

		movement, err := e.GetMovement(b.EntityId)
		if err != nil {
			return FAILURE, err
		}

		movement.Acceleration.Y = -1
		e.SetMovement(movement)
		return SUCCESS, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.behaviors.Set(b.Id, b)

	return b
}

func (e *ECS) MovingDown(parent Behavior) Behavior {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: LEAF_NODE,
		Parent:    parent.Id,
	}

	b.Tick = func(e *ECS, b Behavior) (Status, error) {
		state, err := e.GetState(b.EntityId)
		if err == nil {
			state.State = Moving
			state.OrientationY = North
			e.SetState(state)
		}

		movement, err := e.GetMovement(b.EntityId)
		if err != nil {
			return FAILURE, err
		}

		movement.Acceleration.Y = 1
		e.SetMovement(movement)
		return SUCCESS, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.behaviors.Set(b.Id, b)

	return b
}

func (e *ECS) MovingLeft(parent Behavior) Behavior {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: LEAF_NODE,
		Parent:    parent.Id,
	}

	b.Tick = func(e *ECS, b Behavior) (Status, error) {
		state, err := e.GetState(b.EntityId)
		if err == nil {
			state.State = Moving
			state.OrientationY = North
			e.SetState(state)
		}

		movement, err := e.GetMovement(b.EntityId)
		if err != nil {
			return FAILURE, err
		}

		movement.Acceleration.X = -1
		e.SetMovement(movement)
		return SUCCESS, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.behaviors.Set(b.Id, b)

	return b
}

func (e *ECS) MovingRight(parent Behavior) Behavior {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: LEAF_NODE,
		Parent:    parent.Id,
	}

	b.Tick = func(e *ECS, b Behavior) (Status, error) {
		state, err := e.GetState(b.EntityId)
		if err == nil {
			state.State = Moving
			state.OrientationY = North
			e.SetState(state)
		}

		movement, err := e.GetMovement(b.EntityId)
		if err != nil {
			return FAILURE, err
		}

		movement.Acceleration.X = 1
		e.SetMovement(movement)
		return SUCCESS, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.behaviors.Set(b.Id, b)

	return b
}

func (e *ECS) BindBehavior(entity Entity, behavior Behavior) Entity {
	behavior.EntityId = entity.Id
	entity.BehaviorId = behavior.Id

	e.behaviors = e.behaviors.Set(behavior.Id, behavior)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetEntityBehavior(entityId data.GenerationalIndex) (Behavior, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return Behavior{}, err
	}

	behavior := e.behaviors.Get(entity.BehaviorId)
	if !e.behaviorAllocator.IsLive(behavior.Id) {
		return behavior, ErrAttributeNotFound
	}

	return behavior, nil
}

func (e *ECS) GetBehavior(id data.GenerationalIndex) (Behavior, error) {
	behavior := e.behaviors.Get(id)
	if !e.behaviorAllocator.IsLive(behavior.Id) {
		return behavior, ErrAttributeNotFound
	}

	return behavior, nil
}

func (e *ECS) GetAllBehavior() []Behavior {
	return e.behaviors.GetAll(e.behaviorAllocator)
}

func (e *ECS) SetBehavior(behavior Behavior) {
	e.behaviors = e.behaviors.Set(behavior.Id, behavior)
}
