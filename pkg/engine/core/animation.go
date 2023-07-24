package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
)

func (e *ECS) AddAnimation(entity Entity, a attribute.Animation) Entity {
	animationId := e.animationAllocator.Allocate()

	a.Id = animationId
	a.EntityId = entity.Id
	entity.AnimationId = animationId

	e.animation = e.animation.Set(animationId, a)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetAnimation(entityId data.GenerationalIndex) (attribute.Animation, error) {
	entity, err := e.GetEntity(entityId)
	if err != nil {
		return attribute.Animation{}, err
	}

	animation := e.animation.Get(entity.AnimationId)
	if !e.animationAllocator.IsLive(animation.Id) {
		return animation, ErrAttributeNotFound
	}

	return animation, nil
}

func (e *ECS) GetAllAnimations() []attribute.Animation {
	return e.animation.GetAll(e.animationAllocator)
}

func (e *ECS) SetAnimation(animation attribute.Animation) {
	e.animation = e.animation.Set(animation.Id, animation)
}

