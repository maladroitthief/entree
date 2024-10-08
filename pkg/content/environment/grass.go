package environment

import (
	"math/rand"

	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
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
	faction := e.NewFaction(core.Plant)

	position := e.NewPosition(x, y, 0.25)
	dimension := e.NewDimension(
		mosaic.Vector{X: position.X, Y: position.Y},
		mosaic.Vector{X: 32, Y: 32},
	)

	sprite := grassSprites[rand.Intn(len(grassSprites))]
	animation := e.NewAnimation("tiles", sprite)
	animation.Static = true

	entity := e.NewEntity("grass")
	entity = e.BindFaction(entity, faction)
	entity = e.BindState(entity, state)
	entity = e.BindPosition(entity, position)
	entity = e.BindDimension(entity, dimension)
	entity = e.BindAnimation(entity, animation)

	return entity
}
