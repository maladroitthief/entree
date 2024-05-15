package environment

import (
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
)

func Wall(e *core.ECS, x, y float64) core.Entity {
	state := e.NewState()

	faction := e.NewFaction(core.Stone)

	position := e.NewPosition(x, y, 1.4)
	dimension := e.NewDimension(
		mosaic.Vector{X: position.X, Y: position.Y},
		mosaic.Vector{X: 32, Y: 32},
	)
	collider := e.NewCollider(0.001)
	collider.ColliderType = core.Immovable

	animation := e.NewAnimation("tiles", "rock_1")
	animation.Static = true

	entity := e.NewEntity("wall")
	entity = e.BindFaction(entity, faction)
	entity = e.BindState(entity, state)
	entity = e.BindPosition(entity, position)
	entity = e.BindDimension(entity, dimension)
	entity = e.BindCollider(entity, collider)
	entity = e.BindAnimation(entity, animation)

	return entity
}
