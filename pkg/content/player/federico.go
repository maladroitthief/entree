package player

import (
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
)

func NewFederico(world *content.World) core.Entity {
	state := world.ECS.NewState()
	faction := world.ECS.NewFaction(core.Human)

	position := world.ECS.NewPosition(0, 0, 1.6)
	movement := world.ECS.NewMovement()
	dimension := world.ECS.NewDimension(
		mosaic.Vector{X: position.X, Y: position.Y},
		mosaic.Vector{X: 16, Y: 16},
	)
	dimension.Offset = mosaic.Vector{X: 0, Y: -6}
	collider := world.ECS.NewCollider(1.0)

	animation := world.ECS.NewAnimation("federico", "idle_front_1")
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

	entity := world.ECS.NewEntity()
	entity = world.ECS.BindState(entity, state)
	entity = world.ECS.BindFaction(entity, faction)
	entity = world.ECS.BindPosition(entity, position)
	entity = world.ECS.BindMovement(entity, movement)
	entity = world.ECS.BindDimension(entity, dimension)
	entity = world.ECS.BindCollider(entity, collider)
	entity = world.ECS.BindAnimation(entity, animation)

	return entity
}
