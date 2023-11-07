package environment

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func Weeds(e *core.ECS, x, y float64) core.Entity {
	state := attribute.NewState()

	position := attribute.NewPosition(x, y, 1.25)
	dimension := attribute.NewDimension(
		data.Vector{X: position.X, Y: position.Y},
		data.Vector{X: 32, Y: 32},
	)
	collider := attribute.NewCollider()
	collider.ColliderType = attribute.Impeding
	collider.ImpedingRate = 0.6

	animation := attribute.NewAnimation("tiles", "weeds_1")
	animation.Static = true

	entity := e.NewEntity()
	entity = e.AddState(entity, state)
	entity = e.AddPosition(entity, position)
	entity = e.AddDimension(entity, dimension)
	entity = e.AddCollider(entity, collider)
	entity = e.AddAnimation(entity, animation)

	return entity
}
