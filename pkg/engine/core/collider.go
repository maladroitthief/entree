package core

import "github.com/maladroitthief/caravan"

type ColliderType int

const (
	Immovable ColliderType = iota
	Moveable
	Impeding

	BaseImpedingRate = 1.0
	MaxImpedingRate  = 1.0
)

type Collider struct {
	Id       caravan.GIDX
	EntityId caravan.GIDX

	ColliderType ColliderType
	ImpedingRate float64
}

func (ecs *ECS) NewCollider(impedingRate float64) Collider {
	collider := Collider{
		Id:           ecs.colliderAllocator.Allocate(),
		ColliderType: Moveable,
		ImpedingRate: impedingRate,
	}
	ecs.colliders.Set(collider.Id, collider)

	return collider
}

func (ecs *ECS) BindCollider(entity Entity, collider Collider) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.colliderMu.Lock()
	defer ecs.colliderMu.Unlock()

	collider.EntityId = entity.Id
	entity.ColliderId = collider.Id

	ecs.colliders = ecs.colliders.Set(collider.Id, collider)
	ecs.entities = ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetColliderById(id caravan.GIDX) (Collider, error) {
	ecs.colliderMu.RLock()
	defer ecs.colliderMu.RUnlock()

	collider := ecs.colliders.Get(id)
	if !ecs.colliderAllocator.IsLive(collider.Id) {
		return collider, ErrAttributeNotFound
	}

	return collider, nil
}

func (ecs *ECS) GetCollider(entity Entity) (Collider, error) {
	return ecs.GetColliderById(entity.ColliderId)
}

func (ecs *ECS) GetAllColliders() []Collider {
	ecs.colliderMu.RLock()
	defer ecs.colliderMu.RUnlock()

	return ecs.colliders.GetAll(ecs.colliderAllocator)
}

func (ecs *ECS) SetCollider(collider Collider) {
	ecs.colliderMu.Lock()
	defer ecs.colliderMu.Unlock()

	ecs.colliders = ecs.colliders.Set(collider.Id, collider)
}
