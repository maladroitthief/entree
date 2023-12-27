package player

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func NewFederico(e *core.ECS) core.Entity {
	ai := e.NewAI(core.Player)
	state := e.NewState()

	position := e.NewPosition(100, 100, 1.6)
	movement := e.NewMovement()
	dimension := e.NewDimension(
		data.Vector{X: position.X, Y: position.Y},
		data.Vector{X: 16, Y: 16},
	)
	dimension.Offset = data.Vector{X: 0, Y: -6}
	collider := e.NewCollider()

	animation := e.NewAnimation("federico", "idle_front_1")
	animation.VariantMax = 6
	animation.Speed = 50
	animation.Sprites = map[string][]string{
		"idle_front":      core.SpriteArray("idle_front", 6),
		"idle_front_side": core.SpriteArray("idle_front_side", 6),
		"idle_back":       core.SpriteArray("idle_back", 6),
		"idle_back_side":  core.SpriteArray("idle_back_side", 6),
		"move_front":      core.SpriteArray("move_front", 6),
		"move_front_side": core.SpriteArray("move_front_side", 6),
		"move_back":       core.SpriteArray("move_back", 6),
		"move_back_side":  core.SpriteArray("move_back_side", 6),
	}

	entity := e.NewEntity()
	entity = e.BindAI(entity, ai)
	entity = e.BindState(entity, state)
	entity = e.BindPosition(entity, position)
	entity = e.BindMovement(entity, movement)
	entity = e.BindDimension(entity, dimension)
	entity = e.BindCollider(entity, collider)
	entity = e.BindAnimation(entity, animation)

	return entity
}
