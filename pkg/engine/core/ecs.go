package core

import (
	"errors"

	"github.com/maladroitthief/entree/common/data"
)

var (
	ErrEnityNotFound     = errors.New("entity not found at id")
	ErrAttributeNotFound = errors.New("attribute not found at id")
)

type ECS struct {
	entityAllocator    *data.GenerationalIndexAllocator
	entities           data.GenerationalIndexArray[Entity]
	aiAllocator        *data.GenerationalIndexAllocator
	ai                 data.GenerationalIndexArray[AI]
	commandAllocator   *data.GenerationalIndexAllocator
	command            data.GenerationalIndexArray[Command]
	stateAllocator     *data.GenerationalIndexAllocator
	states             data.GenerationalIndexArray[State]
	movementAllocator  *data.GenerationalIndexAllocator
	movements          data.GenerationalIndexArray[Movement]
	positionAllocator  *data.GenerationalIndexAllocator
	positions          data.GenerationalIndexArray[Position]
	dimensionAllocator *data.GenerationalIndexAllocator
	dimensions         data.GenerationalIndexArray[Dimension]
	colliderAllocator  *data.GenerationalIndexAllocator
	colliders          data.GenerationalIndexArray[Collider]
	animationAllocator *data.GenerationalIndexAllocator
	animations         data.GenerationalIndexArray[Animation]
}

func NewECS() *ECS {
	ecs := &ECS{
		entityAllocator:    data.NewGenerationalIndexAllocator(),
		entities:           data.NewGenerationalIndexArray[Entity](),
		aiAllocator:        data.NewGenerationalIndexAllocator(),
		ai:                 data.NewGenerationalIndexArray[AI](),
		commandAllocator:   data.NewGenerationalIndexAllocator(),
		command:            data.NewGenerationalIndexArray[Command](),
		stateAllocator:     data.NewGenerationalIndexAllocator(),
		states:             data.NewGenerationalIndexArray[State](),
		movementAllocator:  data.NewGenerationalIndexAllocator(),
		movements:          data.NewGenerationalIndexArray[Movement](),
		positionAllocator:  data.NewGenerationalIndexAllocator(),
		positions:          data.NewGenerationalIndexArray[Position](),
		dimensionAllocator: data.NewGenerationalIndexAllocator(),
		dimensions:         data.NewGenerationalIndexArray[Dimension](),
		colliderAllocator:  data.NewGenerationalIndexAllocator(),
		colliders:          data.NewGenerationalIndexArray[Collider](),
		animationAllocator: data.NewGenerationalIndexAllocator(),
		animations:         data.NewGenerationalIndexArray[Animation](),
	}

	return ecs
}
