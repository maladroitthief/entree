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
		1,
		data.Vector{X: 16, Y: 16},
	)

	animation := attribute.NewAnimation("test", "wall")
  animation.Static = true

	entity := e.NewEntity()
	entity = e.AddState(entity, state)
	entity = e.AddPhysics(entity, physics)
	entity = e.AddAnimation(entity, animation)

	return entity
}
