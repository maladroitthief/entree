package core

import (
	"math/rand"

	"github.com/maladroitthief/entree/common/data"
	"github.com/rs/zerolog/log"
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

type BehaviorTree struct {
	Root  data.GenerationalIndex
	stack *data.Stack[data.GenerationalIndex]
}

type Behavior struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Status    Status
	Archetype Archetype
	Parent    data.GenerationalIndex
	Children  []data.GenerationalIndex

	Tick func(*ECS) (Behavior, error)
}

func NewBehaviorTree(root data.GenerationalIndex) *BehaviorTree {
	return &BehaviorTree{
		Root:  root,
		stack: data.NewStack[data.GenerationalIndex](),
	}
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

	b.Tick = func(e *ECS) (Behavior, error) {
		log.Debug().Any("root", b)
		ai := e.ai.Get(b.EntityId)
		if !e.behaviorAllocator.IsLive(ai.Id) {
			b.Status = FAILURE
			return b, ErrAttributeNotFound
		}

		child := e.behaviors.Get(b.Children[0])
		if !e.behaviorAllocator.IsLive(child.Id) {
			b.Status = FAILURE
			return b, ErrAttributeNotFound
		}

		ai.BehaviorStack.Push(child.Id)
		e.SetAI(ai)

		child, err := child.Tick(e)
		if err != nil {
			b.Status = FAILURE
			return b, err
		}
		b.Status = child.Status
		e.SetBehavior(b)

		return b, nil
	}
	e.SetBehavior(b)

	return b
}

func (e *ECS) RandomSequence(parent Behavior) (Behavior, Behavior) {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: COMPOSITE_NODE,
		Parent:    parent.Id,
		Children:  make([]data.GenerationalIndex, 0),
	}

	b.Tick = func(e *ECS) (Behavior, error) {
		log.Debug().Any("randomSequence", b)
		if len(b.Children) <= 0 {
			b.Status = FAILURE
			return b, ErrAttributeNotFound
		}

		child := e.behaviors.Get(b.Children[rand.Intn(len(b.Children))])
		if !e.behaviorAllocator.IsLive(child.Id) {
			b.Status = FAILURE
			return b, ErrAttributeNotFound
		}

		child, err := child.Tick(e)
		if err != nil {
			b.Status = FAILURE
			return b, err
		}

		b.Status = child.Status
		e.SetBehavior(b)
		return b, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.SetBehavior(b)

	return parent, b
}

func (e *ECS) Inverter(parent Behavior) (Behavior, Behavior) {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: DECORATOR_NODE,
		Parent:    parent.Id,
		Children:  make([]data.GenerationalIndex, 1),
	}

	b.Tick = func(e *ECS) (Behavior, error) {
		log.Debug().Any("inverter", b)
		child := e.behaviors.Get(b.Children[0])
		if !e.behaviorAllocator.IsLive(child.Id) {
			b.Status = FAILURE
			return b, ErrAttributeNotFound
		}
		child, err := child.Tick(e)
		if err != nil {
			b.Status = FAILURE
			return b, err
		}

		switch child.Status {
		case SUCCESS:
			child.Status = FAILURE
			return b, nil
		case FAILURE:
			child.Status = SUCCESS
			return child, nil
		}

		return child, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.SetBehavior(b)

	return parent, b
}

func (e *ECS) MovingUp(parent Behavior) (Behavior, Behavior) {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: LEAF_NODE,
		Parent:    parent.Id,
	}

	b.Tick = func(e *ECS) (Behavior, error) {
		log.Debug().Any("moveUp", b)
		state, err := e.GetState(b.EntityId)
		if err == nil {
			state.State = Moving
			state.OrientationY = North
			e.SetState(state)
		}

		movement, err := e.GetMovement(b.EntityId)
		if err != nil {
			b.Status = FAILURE
			return b, err
		}

		movement.Acceleration.Y = -1
		e.SetMovement(movement)
		b.Status = SUCCESS
		return b, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.SetBehavior(b)

	return parent, b
}

func (e *ECS) MovingDown(parent Behavior) (Behavior, Behavior) {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: LEAF_NODE,
		Parent:    parent.Id,
	}

	b.Tick = func(e *ECS) (Behavior, error) {
		log.Debug().Any("moveDown", b)
		state, err := e.GetState(b.EntityId)
		if err == nil {
			state.State = Moving
			state.OrientationY = North
			e.SetState(state)
		}

		movement, err := e.GetMovement(b.EntityId)
		if err != nil {
			b.Status = FAILURE
			return b, err
		}

		movement.Acceleration.Y = 1
		e.SetMovement(movement)
		b.Status = SUCCESS
		return b, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.SetBehavior(b)

	return parent, b
}

func (e *ECS) MovingLeft(parent Behavior) (Behavior, Behavior) {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: LEAF_NODE,
		Parent:    parent.Id,
	}

	b.Tick = func(e *ECS) (Behavior, error) {
		log.Debug().Any("moveLeft", b)
		state, err := e.GetState(b.EntityId)
		if err == nil {
			state.State = Moving
			state.OrientationY = North
			e.SetState(state)
		}

		movement, err := e.GetMovement(b.EntityId)
		if err != nil {
			b.Status = FAILURE
			return b, err
		}

		movement.Acceleration.X = -1
		e.SetMovement(movement)
		b.Status = SUCCESS
		return b, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.SetBehavior(b)

	return parent, b
}

func (e *ECS) MovingRight(parent Behavior) (Behavior, Behavior) {
	b := Behavior{
		Id:        e.behaviorAllocator.Allocate(),
		Status:    SUCCESS,
		Archetype: LEAF_NODE,
		Parent:    parent.Id,
	}

	b.Tick = func(e *ECS) (Behavior, error) {
		log.Debug().Any("moveRight", b)
		state, err := e.GetState(b.EntityId)
		if err == nil {
			state.State = Moving
			state.OrientationY = North
			e.SetState(state)
		}

		movement, err := e.GetMovement(b.EntityId)
		if err != nil {
			b.Status = FAILURE
			return b, err
		}

		movement.Acceleration.X = 1
		e.SetMovement(movement)
		b.Status = SUCCESS
		return b, nil
	}
	parent.Children = append(parent.Children, b.Id)
	e.SetBehavior(parent)
	e.SetBehavior(b)

	return parent, b
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
