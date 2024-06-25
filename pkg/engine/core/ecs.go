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
	entityMu    sync.RWMutex
	entities    *caravan.GIDXPool[Entity]
	aiMu        sync.RWMutex
	ai          *caravan.GIDXPool[AI]
	commandMu   sync.RWMutex
	commands    *caravan.GIDXPool[Command]
	stateMu     sync.RWMutex
	states      *caravan.GIDXPool[State]
	movementMu  sync.RWMutex
	movements   *caravan.GIDXPool[Movement]
	positionMu  sync.RWMutex
	positions   *caravan.GIDXPool[Position]
	dimensionMu sync.RWMutex
	dimensions  *caravan.GIDXPool[Dimension]
	colliderMu  sync.RWMutex
	colliders   *caravan.GIDXPool[Collider]
	animationMu sync.RWMutex
	animations  *caravan.GIDXPool[Animation]
	factionMu   sync.RWMutex
	factions    *caravan.GIDXPool[Faction]
}

func NewECS() *ECS {
	ecs := &ECS{
		entities:   caravan.NewGIDXPool[Entity](),
		ai:         caravan.NewGIDXPool[AI](),
		commands:   caravan.NewGIDXPool[Command](),
		states:     caravan.NewGIDXPool[State](),
		movements:  caravan.NewGIDXPool[Movement](),
		positions:  caravan.NewGIDXPool[Position](),
		dimensions: caravan.NewGIDXPool[Dimension](),
		colliders:  caravan.NewGIDXPool[Collider](),
		animations: caravan.NewGIDXPool[Animation](),
		factions:   caravan.NewGIDXPool[Faction](),
	}

	return ecs
}
