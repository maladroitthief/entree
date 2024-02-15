package core

import (
	"github.com/maladroitthief/entree/common/data"
)

type Position struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	X float64
	Y float64
	Z float64
}

func (e *ECS) NewPosition(x, y, z float64) Position {
	position := Position{
		Id: e.positionAllocator.Allocate(),
		X:  x,
		Y:  y,
		Z:  z,
	}
	e.positions.Set(position.Id, position)

	return position
}

func (e *ECS) BindPosition(entity Entity, position Position) Entity {
	position.EntityId = entity.Id
	entity.PositionId = position.Id

	e.positions = e.positions.Set(position.Id, position)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetPosition(entity Entity) (Position, error) {
	return e.GetPositionById(entity.PositionId)
}
func (e *ECS) GetPositionById(id data.GenerationalIndex) (Position, error) {
	position := e.positions.Get(id)
	if !e.positionAllocator.IsLive(position.Id) {
		return position, ErrAttributeNotFound
	}

	return position, nil
}

func (e *ECS) GetAllPositions() []Position {
	return e.positions.GetAll(e.positionAllocator)
}

func (e *ECS) SetPosition(position Position) {
	e.positions = e.positions.Set(position.Id, position)
}
