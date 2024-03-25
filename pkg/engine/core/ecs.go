package core

import (
	"errors"
	"sync"

	"github.com/maladroitthief/caravan"
)

var (
	ErrEnityNotFound     = errors.New("entity not found at id")
	ErrAttributeNotFound = errors.New("attribute not found at id")
)

type ECS struct {
	entityMu           sync.RWMutex
	entityAllocator    *caravan.GIDXAllocator
	entities           caravan.GIDXArray[Entity]
	aiMu               sync.RWMutex
	aiAllocator        *caravan.GIDXAllocator
	ai                 caravan.GIDXArray[AI]
	commandMu          sync.RWMutex
	commandAllocator   *caravan.GIDXAllocator
	command            caravan.GIDXArray[Command]
	stateMu            sync.RWMutex
	stateAllocator     *caravan.GIDXAllocator
	states             caravan.GIDXArray[State]
	movementMu         sync.RWMutex
	movementAllocator  *caravan.GIDXAllocator
	movements          caravan.GIDXArray[Movement]
	positionMu         sync.RWMutex
	positionAllocator  *caravan.GIDXAllocator
	positions          caravan.GIDXArray[Position]
	dimensionMu        sync.RWMutex
	dimensionAllocator *caravan.GIDXAllocator
	dimensions         caravan.GIDXArray[Dimension]
	colliderMu         sync.RWMutex
	colliderAllocator  *caravan.GIDXAllocator
	colliders          caravan.GIDXArray[Collider]
	animationMu        sync.RWMutex
	animationAllocator *caravan.GIDXAllocator
	animations         caravan.GIDXArray[Animation]
	factionMu          sync.RWMutex
	factionAllocator   *caravan.GIDXAllocator
	factions           caravan.GIDXArray[Faction]
}

func NewECS() *ECS {
	ecs := &ECS{
		entityAllocator:    caravan.NewGIDXAllocator(),
		entities:           caravan.NewGIDXArray[Entity](),
		aiAllocator:        caravan.NewGIDXAllocator(),
		ai:                 caravan.NewGIDXArray[AI](),
		commandAllocator:   caravan.NewGIDXAllocator(),
		command:            caravan.NewGIDXArray[Command](),
		stateAllocator:     caravan.NewGIDXAllocator(),
		states:             caravan.NewGIDXArray[State](),
		movementAllocator:  caravan.NewGIDXAllocator(),
		movements:          caravan.NewGIDXArray[Movement](),
		positionAllocator:  caravan.NewGIDXAllocator(),
		positions:          caravan.NewGIDXArray[Position](),
		dimensionAllocator: caravan.NewGIDXAllocator(),
		dimensions:         caravan.NewGIDXArray[Dimension](),
		colliderAllocator:  caravan.NewGIDXAllocator(),
		colliders:          caravan.NewGIDXArray[Collider](),
		animationAllocator: caravan.NewGIDXAllocator(),
		animations:         caravan.NewGIDXArray[Animation](),
		factionAllocator:   caravan.NewGIDXAllocator(),
		factions:           caravan.NewGIDXArray[Faction](),
	}

	return ecs
}
