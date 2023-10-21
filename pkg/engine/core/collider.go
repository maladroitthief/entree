package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

func (e *ECS) AddCollider(entity Entity, p attribute.Collider) Entity {
	colliderId := e.colliderAllocator.Allocate()

	p.Id = colliderId
	p.EntityId = entity.Id
	entity.ColliderId = colliderId

	e.collider = e.collider.Set(colliderId, p)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetCollider(entityId data.GenerationalIndex) (attribute.Collider, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return attribute.Collider{}, err
	}

	collider := e.collider.Get(entity.ColliderId)
	if !e.colliderAllocator.IsLive(collider.Id) {
		return collider, ErrAttributeNotFound
	}

	return collider, nil
}

func (e *ECS) GetAllCollider() []attribute.Collider {
	return e.collider.GetAll(e.colliderAllocator)
}

func (e *ECS) SetCollider(collider attribute.Collider) {
	e.collider = e.collider.Set(collider.Id, collider)
}
