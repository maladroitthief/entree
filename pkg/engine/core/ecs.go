package core

import (
	"errors"
	"sync"

	"github.com/maladroitthief/entree/common/data"
)

var (
	ErrEnityNotFound     = errors.New("entity not found at id")
	ErrAttributeNotFound = errors.New("attribute not found at id")
)

type ECS struct {
	entityMu           sync.RWMutex
	entityAllocator    *data.GenerationalIndexAllocator
	entities           data.GenerationalIndexArray[Entity]
	aiMu               sync.RWMutex
	aiAllocator        *data.GenerationalIndexAllocator
	ai                 data.GenerationalIndexArray[AI]
	commandMu          sync.RWMutex
	commandAllocator   *data.GenerationalIndexAllocator
	command            data.GenerationalIndexArray[Command]
	stateMu            sync.RWMutex
	stateAllocator     *data.GenerationalIndexAllocator
	states             data.GenerationalIndexArray[State]
	movementMu         sync.RWMutex
	movementAllocator  *data.GenerationalIndexAllocator
	movements          data.GenerationalIndexArray[Movement]
	positionMu         sync.RWMutex
	positionAllocator  *data.GenerationalIndexAllocator
	positions          data.GenerationalIndexArray[Position]
	dimensionMu        sync.RWMutex
	dimensionAllocator *data.GenerationalIndexAllocator
	dimensions         data.GenerationalIndexArray[Dimension]
	colliderMu         sync.RWMutex
	colliderAllocator  *data.GenerationalIndexAllocator
	colliders          data.GenerationalIndexArray[Collider]
	animationMu        sync.RWMutex
	animationAllocator *data.GenerationalIndexAllocator
	animations         data.GenerationalIndexArray[Animation]
	factionMu          sync.RWMutex
	factionAllocator   *data.GenerationalIndexAllocator
	factions           data.GenerationalIndexArray[Faction]
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
		factionAllocator:   data.NewGenerationalIndexAllocator(),
		factions:           data.NewGenerationalIndexArray[Faction](),
	}

	return ecs
}
