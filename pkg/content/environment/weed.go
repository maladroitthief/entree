package environment

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func Weeds(e *core.ECS, x, y float64) core.Entity {
	state := attribute.NewState()

	position := attribute.NewPosition(data.Vector{X: x, Y: y})
	dimension := attribute.NewDimension(position.Position, data.Vector{X: 16, Y: 16})
	collider := attribute.NewCollider()
	collider.ColliderType = attribute.Impeding
	collider.ImpedingRate = 0.6

	animation := attribute.NewAnimation("test", "weeds", 0.25)
	animation.Static = true

	entity := e.NewEntity()
	entity = e.AddState(entity, state)
	entity = e.AddPosition(entity, position)
	entity = e.AddDimension(entity, dimension)
	entity = e.AddCollider(entity, collider)
	entity = e.AddAnimation(entity, animation)

	return entity
}
