package environment

import (
	"math/rand"

	"github.com/maladroitthief/entree/common/data"
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
	state := e.NewState()

	position := e.NewPosition(x, y, 0.25)
	dimension := e.NewDimension(
		data.Vector{X: position.X, Y: position.Y},
		data.Vector{X: 32, Y: 32},
	)

	sprite := grassSprites[rand.Intn(len(grassSprites))]
	animation := e.NewAnimation("tiles", sprite)
	animation.Static = true

	entity := e.NewEntity()
	entity = e.BindState(entity, state)
	entity = e.BindPosition(entity, position)
	entity = e.BindDimension(entity, dimension)
	entity = e.BindAnimation(entity, animation)

	return entity
}
