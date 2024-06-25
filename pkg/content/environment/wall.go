package environment

import (
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
	"github.com/rs/zerolog/log"
)

func Wall(world *content.World, x, y float64) core.Entity {
	item := content.WorldItem{
		Entity:   world.ECS.NewEntity("wall"),
		State:    world.ECS.NewState(),
		Faction:  world.ECS.NewFaction(core.Stone),
		Position: world.ECS.NewPosition(x, y, 1.4),
		Dimension: world.ECS.NewDimension(
			mosaic.Vector{X: x, Y: y},
			mosaic.Vector{X: 32, Y: 32},
		),
		Collider:  world.ECS.NewCollider(0.0001),
		Animation: world.ECS.NewAnimation("tiles", "rock_1"),
	}
	item.Collider.ColliderType = core.Immovable
	item.Animation.Static = true

	item, err := world.NewItem(item)
	if err != nil {
		log.Error().Err(err).Msg("unable to create wall")
	}

	return item.Entity
}
