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
	behaviorAllocator  *data.GenerationalIndexAllocator
	behavior           data.GenerationalIndexArray[attribute.Behavior]
	stateAllocator     *data.GenerationalIndexAllocator
	state              data.GenerationalIndexArray[attribute.State]
	movementAllocator  *data.GenerationalIndexAllocator
	movement           data.GenerationalIndexArray[attribute.Movement]
	positionAllocator  *data.GenerationalIndexAllocator
	position           data.GenerationalIndexArray[attribute.Position]
	dimensionAllocator *data.GenerationalIndexAllocator
	dimension          data.GenerationalIndexArray[attribute.Dimension]
	colliderAllocator  *data.GenerationalIndexAllocator
	collider           data.GenerationalIndexArray[attribute.Collider]
	animationAllocator *data.GenerationalIndexAllocator
	animation          data.GenerationalIndexArray[attribute.Animation]
}

func NewECS() *ECS {
	ecs := &ECS{
		entityAllocator:    data.NewGenerationalIndexAllocator(),
		entities:           data.NewGenerationalIndexArray[Entity](),
		aiAllocator:        data.NewGenerationalIndexAllocator(),
		ai:                 data.NewGenerationalIndexArray[attribute.AI](),
		behaviorAllocator:  data.NewGenerationalIndexAllocator(),
		behavior:           data.NewGenerationalIndexArray[attribute.Behavior](),
		stateAllocator:     data.NewGenerationalIndexAllocator(),
		state:              data.NewGenerationalIndexArray[attribute.State](),
		movementAllocator:  data.NewGenerationalIndexAllocator(),
		movement:           data.NewGenerationalIndexArray[attribute.Movement](),
		positionAllocator:  data.NewGenerationalIndexAllocator(),
		position:           data.NewGenerationalIndexArray[attribute.Position](),
		dimensionAllocator: data.NewGenerationalIndexAllocator(),
		dimension:          data.NewGenerationalIndexArray[attribute.Dimension](),
		colliderAllocator:  data.NewGenerationalIndexAllocator(),
		collider:           data.NewGenerationalIndexArray[attribute.Collider](),
		animationAllocator: data.NewGenerationalIndexAllocator(),
		animation:          data.NewGenerationalIndexArray[attribute.Animation](),
	}

	return ecs
}
