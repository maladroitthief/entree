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
	Player
	Computer
)

var (
	ErrAIContextNil = errors.New("NewAI nil context")
)

type AI struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Node    bt.Node
	Context context.Context
	cancel  context.CancelFunc
}

func (e *ECS) NewAI(ctx context.Context, node bt.Node) AI {
	if ctx == nil {
		log.Panic(ErrAIContextNil)
	}

	ai := AI{
		Id:   e.aiAllocator.Allocate(),
		Node: node,
	}

	ai.Context, ai.cancel = context.WithCancel(ctx)
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
