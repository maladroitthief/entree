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

func (e *ECS) GetMovement(entityId data.GenerationalIndex) (Movement, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return Movement{}, err
	}

	movement := e.movements.Get(entity.MovementId)
	if !e.movementAllocator.IsLive(movement.Id) {
		return movement, ErrAttributeNotFound
	}

	return movement, nil
}

func (e *ECS) GetAllMovement() []Movement {
	return e.movements.GetAll(e.movementAllocator)
}

func (e *ECS) SetMovement(movement Movement) {
	e.movements = e.movements.Set(movement.Id, movement)
}
