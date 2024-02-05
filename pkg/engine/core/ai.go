package core

import (
	"context"
	"errors"
	"log"

	"github.com/maladroitthief/entree/common/data"
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
)

type BehaviorType int

const (
	None BehaviorType = iota
	Computer
)

var (
	ErrAIContextNil = errors.New("NewAI nil context")
)

type AI struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex
	TargetId data.GenerationalIndex

	Node    bt.Node
	Context context.Context
	Cancel  context.CancelFunc
}

func (e *ECS) NewAI(ctx context.Context, node bt.Node) AI {
	if ctx == nil {
		log.Panic(ErrAIContextNil)
	}

	ai := AI{
		Id:   e.aiAllocator.Allocate(),
		Node: node,
	}

	ai.Context, ai.Cancel = context.WithCancel(ctx)
	e.SetAI(ai)

	return ai
}

func (e *ECS) BindAI(entity Entity, ai AI) Entity {
	ai.EntityId = entity.Id
	entity.AIId = ai.Id

	e.ai = e.ai.Set(ai.Id, ai)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetAIById(id data.GenerationalIndex) (AI, error) {
	ai := e.ai.Get(id)
	if !e.aiAllocator.IsLive(ai.Id) {
		return ai, ErrAttributeNotFound
	}

	return ai, nil
}

func (e *ECS) GetAI(entity Entity) (AI, error) {
	return e.GetAIById(entity.AIId)
}

func (e *ECS) GetAllAI() []AI {
	return e.ai.GetAll(e.aiAllocator)
}

func (e *ECS) SetAI(ai AI) {
	e.ai = e.ai.Set(ai.Id, ai)
}
