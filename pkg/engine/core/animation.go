package core

import (
	"fmt"

	"github.com/maladroitthief/entree/common/data"
)

const (
	DefaultSize  = 32
	DefaultSpeed = 60
)

type Animation struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Counter     int
	Static      bool
	Speed       float64
	Variant     int
	VariantMax  int
	SpriteSheet string
	Sprite      string
	Sprites     map[string][]string
}

func (ecs *ECS) NewAnimation(sheet, defaultSprite string) Animation {
	animation := Animation{
		Id:          ecs.animationAllocator.Allocate(),
		Speed:       DefaultSpeed,
		Counter:     0,
		Variant:     1,
		VariantMax:  1,
		SpriteSheet: sheet,
		Sprite:      defaultSprite,
		Sprites:     map[string][]string{},
	}
	ecs.animations.Set(animation.Id, animation)

	return animation
}

func SpriteArray(spriteName string, count int) []string {
	results := make([]string, count)

	for i := 0; i < count; i++ {
		results[i] = fmt.Sprintf("%s_%d", spriteName, i+1)
	}

	return results
}

func (ecs *ECS) BindAnimation(entity Entity, animation Animation) Entity {
	ecs.entityMu.Lock()
	defer ecs.entityMu.Unlock()
	ecs.animationMu.Lock()
	defer ecs.animationMu.Unlock()

	animation.EntityId = entity.Id
	entity.AnimationId = animation.Id

	ecs.animations = ecs.animations.Set(animation.Id, animation)
	ecs.entities = ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetAnimation(entity Entity) (Animation, error) {
	return ecs.GetAnimationById(entity.AnimationId)
}

func (ecs *ECS) GetAnimationById(id data.GenerationalIndex) (Animation, error) {
	ecs.animationMu.RLock()
	defer ecs.animationMu.RUnlock()

	animation := ecs.animations.Get(id)
	if !ecs.animationAllocator.IsLive(animation.Id) {
		return animation, ErrAttributeNotFound
	}

	return animation, nil
}

func (ecs *ECS) GetAllAnimations() []Animation {
	ecs.animationMu.RLock()
	defer ecs.animationMu.RUnlock()

	return ecs.animations.GetAll(ecs.animationAllocator)
}

func (ecs *ECS) SetAnimation(animation Animation) {
	ecs.animationMu.Lock()
	defer ecs.animationMu.Unlock()

	ecs.animations = ecs.animations.Set(animation.Id, animation)
}
