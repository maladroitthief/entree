package core

import (
	"context"
	"errors"
	"log"

	"github.com/maladroitthief/caravan"
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/mosaic"
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
	Id             caravan.GIDX
	EntityId       caravan.GIDX
	TargetEntityId caravan.GIDX
	TargetLocation mosaic.Vector
	PathToTarget   []mosaic.Vector
	Targets        Archetype

	Node    bt.Node
	Context context.Context
	Cancel  context.CancelFunc
}

func (ecs *ECS) NewAI(ctx context.Context, node bt.Node) AI {
	if ctx == nil {
		log.Panic(ErrAIContextNil)
	}

	ai := AI{
		Id:   ecs.ai.Allocate(),
		Node: node,
	}

	ai.Context, ai.Cancel = context.WithCancel(ctx)
	ecs.SetAI(ai)

	return ai
}

func (ecs *ECS) BindAI(entity Entity, ai AI) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.aiMu.Lock()
	defer ecs.aiMu.Unlock()

	ai.EntityId = entity.Id
	entity.AIId = ai.Id

	ecs.ai.Set(ai.Id, ai)
	ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetAIById(id caravan.GIDX) (AI, error) {
	ecs.aiMu.RLock()
	defer ecs.aiMu.RUnlock()

	ai := ecs.ai.Get(id)
	if !ecs.ai.IsLive(ai.Id) {
		return ai, ErrAttributeNotFound
	}

	return ai, nil
}

func (ecs *ECS) GetAI(entity Entity) (AI, error) {
	return ecs.GetAIById(entity.AIId)
}

func (ecs *ECS) GetAllAI() []AI {
	ecs.aiMu.RLock()
	defer ecs.aiMu.RUnlock()

	return ecs.ai.GetAll()
}

func (ecs *ECS) SetAI(ai AI) {
	ecs.aiMu.Lock()
	defer ecs.aiMu.Unlock()

	ecs.ai.Set(ai.Id, ai)
}

func (ecs *ECS) AIActive(ai AI) bool {
	return ecs.ai.IsLive(ai.Id)
}
