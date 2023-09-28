package environment

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func Weeds(e *core.ECS, x, y float64) core.Entity {
	state := attribute.NewState()

	physics := attribute.NewPhysics(
		data.Vector{X: x, Y: y},
		data.Vector{X: 16, Y: 16},
	)
	physics.CollisionType = attribute.Impeding
	physics.ImpedingRate = 0.6

	animation := attribute.NewAnimation("test", "weeds", 0.25)
	animation.Static = true

	entity := e.NewEntity()
	entity = e.AddState(entity, state)
	entity = e.AddPhysics(entity, physics)
	entity = e.AddAnimation(entity, animation)

	return entity
}
