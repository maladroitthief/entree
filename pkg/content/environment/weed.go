package environment

import (
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
	"github.com/rs/zerolog/log"
)

func Weeds(world *content.World, x, y float64) core.Entity {
	item := content.WorldItem{
		Entity:   world.ECS.NewEntity("weeds"),
		State:    world.ECS.NewState(),
		Faction:  world.ECS.NewFaction(core.Plant),
		Position: world.ECS.NewPosition(x, y, 0.25),
		Dimension: world.ECS.NewDimension(
			mosaic.Vector{X: x, Y: y},
			mosaic.Vector{X: 32, Y: 32},
		),
		Collider:  world.ECS.NewCollider(0.6),
		Animation: world.ECS.NewAnimation("tiles", "weeds_1"),
	}
	item.Animation.Static = true
	item.Collider.ColliderType = core.Impeding

	item, err := world.NewItem(item)
	if err != nil {
		log.Error().Err(err).Msg("unable to create weeds")
	}

	return item.Entity
}
