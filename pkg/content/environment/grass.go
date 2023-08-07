package environment

import (
	"math/rand"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

var (
	grassSprites = []string{
		"grass",
		"flowers",
		"tall_grass",
	}
)

func Grass(e *core.ECS, x, y float64) core.Entity {
	state := attribute.NewState()

	physics := attribute.NewPhysics(
		data.Vector{X: x, Y: y},
		0,
		data.Vector{X: 16, Y: 16},
	)
	physics.CollisionType = attribute.Impeding
  physics.ImpedingRate = 0.2

	sprite := grassSprites[rand.Intn(len(grassSprites))]
	animation := attribute.NewAnimation("test", sprite)
	animation.Static = true

	entity := e.NewEntity()
	entity = e.AddState(entity, state)
	entity = e.AddPhysics(entity, physics)
	entity = e.AddAnimation(entity, animation)

	return entity
}
