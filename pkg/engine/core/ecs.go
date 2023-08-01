package core

import (
	"errors"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

var (
	ErrEnityNotFound     = errors.New("entity not found at id")
	ErrAttributeNotFound = errors.New("attribute not found at id")
)

type ECS struct {
	entityAllocator    *data.GenerationalIndexAllocator
	entities           data.GenerationalIndexArray[Entity]
	aiAllocator        *data.GenerationalIndexAllocator
	ai                 data.GenerationalIndexArray[attribute.AI]
	stateAllocator     *data.GenerationalIndexAllocator
	state              data.GenerationalIndexArray[attribute.State]
	physicsAllocator   *data.GenerationalIndexAllocator
	physics            data.GenerationalIndexArray[attribute.Physics]
	animationAllocator *data.GenerationalIndexAllocator
	animation          data.GenerationalIndexArray[attribute.Animation]
}

func NewECS() *ECS {
	ecs := &ECS{
		entityAllocator:    data.NewGenerationalIndexAllocator(),
		entities:           data.NewGenerationalIndexArray[Entity](),
		aiAllocator:        data.NewGenerationalIndexAllocator(),
		ai:                 data.NewGenerationalIndexArray[attribute.AI](),
		stateAllocator:     data.NewGenerationalIndexAllocator(),
		state:              data.NewGenerationalIndexArray[attribute.State](),
		physicsAllocator:   data.NewGenerationalIndexAllocator(),
		physics:            data.NewGenerationalIndexArray[attribute.Physics](),
		animationAllocator: data.NewGenerationalIndexAllocator(),
		animation:          data.NewGenerationalIndexArray[attribute.Animation](),
	}

	return ecs
}
