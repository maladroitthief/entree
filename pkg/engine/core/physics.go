package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

func (e *ECS) AddPhysics(entity Entity, p attribute.Physics) Entity {
	physicsId := e.physicsAllocator.Allocate()

	p.Id = physicsId
	p.EntityId = entity.Id
	entity.PhysicsId = physicsId

	e.physics = e.physics.Set(physicsId, p)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetPhysics(entityId data.GenerationalIndex) (attribute.Physics, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return attribute.Physics{}, err
	}

	physics := e.physics.Get(entity.PhysicsId)
	if !e.physicsAllocator.IsLive(physics.Id) {
		return physics, ErrAttributeNotFound
	}

	return physics, nil
}

func (e *ECS) GetAllPhysics() []attribute.Physics {
	return e.physics.GetAll(e.physicsAllocator)
}

func (e *ECS) SetPhysics(physics attribute.Physics) {
	e.physics = e.physics.Set(physics.Id, physics)
}
