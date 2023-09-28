package player

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func NewHero(e *core.ECS) core.Entity {
	ai := attribute.NewAI(attribute.Input)
	state := attribute.NewState()

	physics := attribute.NewPhysics(
		data.Vector{X: 100, Y: 100},
		data.Vector{X: 12, Y: 18},
	)
	physics.Offset = data.Vector{X: 0, Y: -6}

	animation := attribute.NewAnimation("hero", "idle_front_1", 0.5)
	animation.VariantMax = 6

	entity := e.NewEntity()
	entity = e.AddAI(entity, ai)
	entity = e.AddState(entity, state)
	entity = e.AddPhysics(entity, physics)
	entity = e.AddAnimation(entity, animation)

	return entity
}
