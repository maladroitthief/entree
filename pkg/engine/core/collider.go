package core

import (
	"github.com/maladroitthief/entree/common/data"
)

type ColliderType int

const (
	Immovable ColliderType = iota
	Moveable
	Impeding

	BaseImpedingRate = 1.0
	MaxImpedingRate  = 1.0
)

type Collider struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	ColliderType ColliderType
	ImpedingRate float64
}

func (e *ECS) NewCollider(impedingRate float64) Collider {
	collider := Collider{
		Id:           e.colliderAllocator.Allocate(),
		ColliderType: Moveable,
		ImpedingRate: impedingRate,
	}
	e.colliders.Set(collider.Id, collider)

	return collider
}

func (e *ECS) BindCollider(entity Entity, collider Collider) Entity {
	collider.EntityId = entity.Id
	entity.ColliderId = collider.Id

	e.colliders = e.colliders.Set(collider.Id, collider)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetColliderById(id data.GenerationalIndex) (Collider, error) {
	collider := e.colliders.Get(id)
	if !e.colliderAllocator.IsLive(collider.Id) {
		return collider, ErrAttributeNotFound
	}

	return collider, nil
}

func (e *ECS) GetCollider(entity Entity) (Collider, error) {
	return e.GetColliderById(entity.ColliderId)
}

func (e *ECS) GetAllColliders() []Collider {
	return e.colliders.GetAll(e.colliderAllocator)
}

func (e *ECS) SetCollider(collider Collider) {
	e.colliders = e.colliders.Set(collider.Id, collider)
}
