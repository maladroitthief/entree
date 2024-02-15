package core

import (
	"github.com/maladroitthief/entree/common/data"
)

const (
	BaseMaxVelocity = 3
	BaseMass        = 10
)

type Movement struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Velocity     data.Vector
	MaxVelocity  float64
	Mass         float64
	Acceleration data.Vector
}

func (e *ECS) NewMovement() Movement {
	movement := Movement{
		Id:          e.movementAllocator.Allocate(),
		Velocity:    data.Vector{X: 0, Y: 0},
		MaxVelocity: BaseMaxVelocity,
		Mass:        BaseMass,
	}
	e.movements.Set(movement.Id, movement)

	return movement
}

func (e *ECS) BindMovement(entity Entity, movement Movement) Entity {
	movement.EntityId = entity.Id
	entity.MovementId = movement.Id

	e.movements = e.movements.Set(movement.Id, movement)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetMovement(entity Entity) (Movement, error) {
	return e.GetMovementById(entity.MovementId)
}
func (e *ECS) GetMovementById(id data.GenerationalIndex) (Movement, error) {
	movement := e.movements.Get(id)
	if !e.movementAllocator.IsLive(movement.Id) {
		return movement, ErrAttributeNotFound
	}

	return movement, nil
}

func (e *ECS) GetAllMovements() []Movement {
	return e.movements.GetAll(e.movementAllocator)
}

func (e *ECS) SetMovement(movement Movement) {
	e.movements = e.movements.Set(movement.Id, movement)
}
