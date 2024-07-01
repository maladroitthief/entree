package weapons

import (
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
	"github.com/rs/zerolog/log"
)

func NewCleaver(world *content.World, x, y float64) core.Entity {
	item := content.WorldItem{
		Entity:   world.ECS.NewEntity("cleaver"),
		State:    world.ECS.NewState(),
		Faction:  world.ECS.NewFaction(core.Metal),
		Position: world.ECS.NewPosition(x, y, 1.6),
		Movement: world.ECS.NewMovement(),
		Dimension: world.ECS.NewDimension(
			mosaic.Vector{X: x, Y: y},
			mosaic.Vector{X: 16, Y: 16},
		),
		Collider:  world.ECS.NewCollider(core.Hitbox, 1.0),
		Animation: world.ECS.NewAnimation("weapons", "cleaver_1"),
	}

	item.Dimension.Offset = mosaic.Vector{X: 0, Y: -6}
	item.Animation.VariantMax = 6
	item.Animation.Speed = 50
	item.Animation.Sprites = map[string][]string{
		"cleaver_swing": core.SpriteArray("cleaver_swing", 6),
	}

	item, err := world.NewItem(item)
	if err != nil {
		log.Error().Err(err).Msg("unable to create cleaver")
	}

	return item.Entity
}
