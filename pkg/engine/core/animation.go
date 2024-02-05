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

func (e *ECS) NewAnimation(sheet, defaultSprite string) Animation {
	animation := Animation{
		Id:          e.animationAllocator.Allocate(),
		Speed:       DefaultSpeed,
		Counter:     0,
		Variant:     1,
		VariantMax:  1,
		SpriteSheet: sheet,
		Sprite:      defaultSprite,
		Sprites:     map[string][]string{},
	}
	e.animations.Set(animation.Id, animation)

	return animation
}

func SpriteArray(spriteName string, count int) []string {
	results := make([]string, count)

	for i := 0; i < count; i++ {
		results[i] = fmt.Sprintf("%s_%d", spriteName, i+1)
	}

	return results
}

func (e *ECS) BindAnimation(entity Entity, animation Animation) Entity {
	animation.EntityId = entity.Id
	entity.AnimationId = animation.Id

	e.animations = e.animations.Set(animation.Id, animation)
	e.entities = e.entities.Set(entity.Id, entity)

	return entity
}

func (e *ECS) GetAnimation(entity Entity) (Animation, error) {
	return e.GetAnimationById(entity.AnimationId)
}

func (e *ECS) GetAnimationById(id data.GenerationalIndex) (Animation, error) {
	animation := e.animations.Get(id)
	if !e.animationAllocator.IsLive(animation.Id) {
		return animation, ErrAttributeNotFound
	}

	return animation, nil
}

func (e *ECS) GetAllAnimations() []Animation {
	return e.animations.GetAll(e.animationAllocator)
}

func (e *ECS) SetAnimation(animation Animation) {
	e.animations = e.animations.Set(animation.Id, animation)
}
