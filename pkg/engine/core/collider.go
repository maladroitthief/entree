package core

import "github.com/maladroitthief/caravan"

type ColliderType int

const (
	Immovable ColliderType = iota
	Moveable
	Impeding
	Hitbox

	BaseImpedingRate = 1.0
	BaseResitution   = 0.5
)

type Collider struct {
	Id       caravan.GIDX
	EntityId caravan.GIDX

	ColliderType ColliderType
	ImpedingRate float64
	Restitution  float64
}

func (ecs *ECS) NewCollider(colliderType ColliderType, impedingRate float64) Collider {
	collider := Collider{
		Id:           ecs.colliders.Allocate(),
		ColliderType: colliderType,
		ImpedingRate: impedingRate,
		Restitution:  BaseResitution,
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

	ecs.colliders.Set(collider.Id, collider)
	ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetColliderById(id caravan.GIDX) (Collider, error) {
	ecs.colliderMu.RLock()
	defer ecs.colliderMu.RUnlock()

	collider := ecs.colliders.Get(id)
	if !ecs.colliders.IsLive(collider.Id) {
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

	return ecs.colliders.GetAll()
}

func (ecs *ECS) SetCollider(collider Collider) {
	ecs.colliderMu.Lock()
	defer ecs.colliderMu.Unlock()

	ecs.colliders.Set(collider.Id, collider)
}

func (ecs *ECS) ColliderActive(collider Collider) bool {
	return ecs.colliders.IsLive(collider.Id)
}
