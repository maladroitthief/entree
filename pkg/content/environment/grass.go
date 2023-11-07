package environment

import (
	"math/rand"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

var (
	grassSprites = []string{
		"grass_1",
		"grass_2",
		"grass_3",
		"grass_4",
	}
)

func Grass(e *core.ECS, x, y float64) core.Entity {
	state := attribute.NewState()

	position := attribute.NewPosition(data.Vector{X: x, Y: y})
	dimension := attribute.NewDimension(position.Position, data.Vector{X: 32, Y: 32})

	sprite := grassSprites[rand.Intn(len(grassSprites))]
	animation := attribute.NewAnimation("tiles", sprite, 0.25)
	animation.Static = true

	entity := e.NewEntity()
	entity = e.AddState(entity, state)
	entity = e.AddPosition(entity, position)
	entity = e.AddDimension(entity, dimension)
	entity = e.AddAnimation(entity, animation)

	return entity
}
