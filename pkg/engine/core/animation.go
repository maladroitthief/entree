package core

import (
	"fmt"

	"github.com/maladroitthief/caravan"
)

const (
	DefaultSize  = 32
	DefaultSpeed = 60
)

type Animation struct {
	Id       caravan.GIDX
	EntityId caravan.GIDX

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
		Id:          ecs.animations.Allocate(),
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

	ecs.animations.Set(animation.Id, animation)
	ecs.entities.Set(entity.Id, entity)

	return entity
}

func (ecs *ECS) GetAnimation(entity Entity) (Animation, error) {
	return ecs.GetAnimationById(entity.AnimationId)
}

func (ecs *ECS) GetAnimationById(id caravan.GIDX) (Animation, error) {
	ecs.animationMu.RLock()
	defer ecs.animationMu.RUnlock()

	animation := ecs.animations.Get(id)
	if !ecs.animations.IsLive(animation.Id) {
		return animation, ErrAttributeNotFound
	}

	return animation, nil
}

func (ecs *ECS) GetAllAnimations() []Animation {
	ecs.animationMu.RLock()
	defer ecs.animationMu.RUnlock()

	return ecs.animations.GetAll()
}

func (ecs *ECS) SetAnimation(animation Animation) {
	ecs.animationMu.Lock()
	defer ecs.animationMu.Unlock()

	ecs.animations.Set(animation.Id, animation)
}

func (ecs *ECS) AnimationActive(animation Animation) bool {
	return ecs.animations.IsLive(animation.Id)
}
