package environment

import (
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
)

func Wall(world *content.World, x, y float64) core.Entity {
	state := world.ECS.NewState()

	faction := world.ECS.NewFaction(core.Stone)

	position := world.ECS.NewPosition(x, y, 1.4)
	dimension := world.ECS.NewDimension(
		mosaic.Vector{X: position.X, Y: position.Y},
		mosaic.Vector{X: 32, Y: 32},
	)
	collider := world.ECS.NewCollider(0.0001)
	collider.ColliderType = core.Immovable

	animation := world.ECS.NewAnimation("tiles", "rock_1")
	animation.Static = true

	entity := world.ECS.NewEntity("wall")
	entity = world.ECS.BindFaction(entity, faction)
	entity = world.ECS.BindState(entity, state)
	entity = world.ECS.BindPosition(entity, position)
	entity = world.ECS.BindDimension(entity, dimension)
	entity = world.ECS.BindCollider(entity, collider)
	entity = world.ECS.BindAnimation(entity, animation)

	return entity
}
