package environment

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func Wall(e *core.ECS, x, y float64) core.Entity {
	state := attribute.NewState()

	physics := attribute.NewPhysics(
		data.Vector{X: x, Y: y},
		data.Vector{X: 16, Y: 16},
	)
	physics.CollisionType = attribute.Immovable

	animation := attribute.NewAnimation("test", "wall", 0.4)
	animation.Static = true

	entity := e.NewEntity()
	entity = e.AddState(entity, state)
	entity = e.AddPhysics(entity, physics)
	entity = e.AddAnimation(entity, animation)

	return entity
}
