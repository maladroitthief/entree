package core

import (
	"github.com/maladroitthief/caravan"
	"github.com/maladroitthief/mosaic"
)

type Position struct {
	Id       caravan.GIDX
	EntityId caravan.GIDX

	X float64
	Y float64
	Z float64
}

func (ecs *ECS) NewPosition(x, y, z float64) Position {
	position := Position{
		Id: ecs.positions.Allocate(),
		X:  x,
		Y:  y,
		Z:  z,
	}
	ecs.positions.Set(position.Id, position)

	return position
}

func (p Position) Vector() mosaic.Vector {
	return mosaic.NewVector(p.X, p.Y)
}

func (ecs *ECS) BindPosition(entity Entity, position Position) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.positionMu.Lock()
	defer ecs.positionMu.Unlock()

	position.EntityId = entity.Id
	entity.PositionId = position.Id

	ecs.positions.Set(position.Id, position)
	ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetPosition(entity Entity) (Position, error) {
	return ecs.GetPositionById(entity.PositionId)
}
func (ecs *ECS) GetPositionById(id caravan.GIDX) (Position, error) {
	ecs.positionMu.RLock()
	defer ecs.positionMu.RUnlock()

	position := ecs.positions.Get(id)
	if !ecs.positions.IsLive(position.Id) {
		return position, ErrAttributeNotFound
	}

	return position, nil
}

func (ecs *ECS) GetAllPositions() []Position {
	ecs.positionMu.RLock()
	defer ecs.positionMu.RUnlock()

	return ecs.positions.GetAll()
}

func (ecs *ECS) SetPosition(position Position) {
	ecs.positionMu.Lock()
	defer ecs.positionMu.Unlock()

	ecs.positions.Set(position.Id, position)
}

func (ecs *ECS) PositionActive(position Position) bool {
	return ecs.positions.IsLive(position.Id)
}
