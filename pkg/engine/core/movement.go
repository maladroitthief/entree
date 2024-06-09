package core

import (
	"github.com/maladroitthief/caravan"
	"github.com/maladroitthief/mosaic"
)

const (
	BaseMaxVelocity = 30
	BaseMass        = 10
)

type Movement struct {
	Id       caravan.GIDX
	EntityId caravan.GIDX

	MaxVelocity float64
	Mass        float64

	Velocity mosaic.Vector
	Force    mosaic.Vector
}

func (ecs *ECS) NewMovement() Movement {
	movement := Movement{
		Id:          ecs.movementAllocator.Allocate(),
		Velocity:    mosaic.Vector{X: 0, Y: 0},
		MaxVelocity: BaseMaxVelocity,
		Mass:        BaseMass,
	}
	ecs.movements.Set(movement.Id, movement)

	return movement
}

func (ecs *ECS) BindMovement(entity Entity, movement Movement) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.movementMu.Lock()
	defer ecs.movementMu.Unlock()

	movement.EntityId = entity.Id
	entity.MovementId = movement.Id

	ecs.movements = ecs.movements.Set(movement.Id, movement)
	ecs.entities = ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetMovement(entity Entity) (Movement, error) {
	return ecs.GetMovementById(entity.MovementId)
}
func (ecs *ECS) GetMovementById(id caravan.GIDX) (Movement, error) {
	ecs.movementMu.RLock()
	defer ecs.movementMu.RUnlock()

	movement := ecs.movements.Get(id)
	if !ecs.movementAllocator.IsLive(movement.Id) {
		return movement, ErrAttributeNotFound
	}

	return movement, nil
}

func (ecs *ECS) GetAllMovements() []Movement {
	ecs.movementMu.RLock()
	defer ecs.movementMu.RUnlock()

	return ecs.movements.GetAll(ecs.movementAllocator)
}

func (ecs *ECS) SetMovement(movement Movement) {
	ecs.movementMu.Lock()
	defer ecs.movementMu.Unlock()

	ecs.movements = ecs.movements.Set(movement.Id, movement)
}
