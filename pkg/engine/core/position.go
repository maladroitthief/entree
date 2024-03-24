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

func (ecs *ECS) NewPosition(x, y, z float64) Position {
	position := Position{
		Id: ecs.positionAllocator.Allocate(),
		X:  x,
		Y:  y,
		Z:  z,
	}
	ecs.positions.Set(position.Id, position)

	return position
}

func (ecs *ECS) BindPosition(entity Entity, position Position) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.positionMu.Lock()
	defer ecs.positionMu.Unlock()

	position.EntityId = entity.Id
	entity.PositionId = position.Id

	ecs.positions = ecs.positions.Set(position.Id, position)
	ecs.entities = ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetPosition(entity Entity) (Position, error) {
	return ecs.GetPositionById(entity.PositionId)
}
func (ecs *ECS) GetPositionById(id data.GenerationalIndex) (Position, error) {
	ecs.positionMu.RLock()
	defer ecs.positionMu.RUnlock()

	position := ecs.positions.Get(id)
	if !ecs.positionAllocator.IsLive(position.Id) {
		return position, ErrAttributeNotFound
	}

	return position, nil
}

func (ecs *ECS) GetAllPositions() []Position {
	ecs.positionMu.RLock()
	defer ecs.positionMu.RUnlock()

	return ecs.positions.GetAll(ecs.positionAllocator)
}

func (ecs *ECS) SetPosition(position Position) {
	ecs.positionMu.Lock()
	defer ecs.positionMu.Unlock()

	ecs.positions = ecs.positions.Set(position.Id, position)
}
