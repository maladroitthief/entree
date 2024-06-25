package environment

import (
	"math/rand"

	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
	"github.com/rs/zerolog/log"
)

var (
	grassSprites = []string{
		"grass_1",
		"grass_2",
		"grass_3",
		"grass_4",
	}
)

func Grass(world *content.World, x, y float64) core.Entity {
	sprite := grassSprites[rand.Intn(len(grassSprites))]
	item := content.WorldItem{
		Entity:   world.ECS.NewEntity("grass"),
		State:    world.ECS.NewState(),
		Faction:  world.ECS.NewFaction(core.Plant),
		Position: world.ECS.NewPosition(x, y, 0.25),
		Dimension: world.ECS.NewDimension(
			mosaic.Vector{X: x, Y: y},
			mosaic.Vector{X: 32, Y: 32},
		),
		Animation: world.ECS.NewAnimation("tiles", sprite),
	}
	item.Animation.Static = true

	item, err := world.NewItem(item)
	if err != nil {
		log.Error().Err(err).Msg("unable to create grass")
	}

	return item.Entity
}
