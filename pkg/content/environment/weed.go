package environment

import (
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
)

func Weeds(e *core.ECS, x, y float64) core.Entity {
	state := e.NewState()

	position := e.NewPosition(x, y, 1.25)
	dimension := e.NewDimension(
		mosaic.Vector{X: position.X, Y: position.Y},
		mosaic.Vector{X: 32, Y: 32},
	)
	collider := e.NewCollider(0.6)
	collider.ColliderType = core.Impeding

	animation := e.NewAnimation("tiles", "weeds_1")
	animation.Static = true

	entity := e.NewEntity()
	entity = e.BindState(entity, state)
	entity = e.BindPosition(entity, position)
	entity = e.BindDimension(entity, dimension)
	entity = e.BindCollider(entity, collider)
	entity = e.BindAnimation(entity, animation)

	return entity
}
