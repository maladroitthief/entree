package player

import (
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
	"github.com/rs/zerolog/log"
)

func NewFederico(world *content.World, x, y float64) core.Entity {
	item := content.WorldItem{
		Entity:   world.ECS.NewEntity("federico"),
		State:    world.ECS.NewState(),
		Faction:  world.ECS.NewFaction(core.Human),
		Position: world.ECS.NewPosition(x, y, 1.6),
		Movement: world.ECS.NewMovement(),
		Dimension: world.ECS.NewDimension(
			mosaic.Vector{X: x, Y: y},
			mosaic.Vector{X: 16, Y: 16},
		),
		Collider:  world.ECS.NewCollider(core.Moveable, 1.0),
		Animation: world.ECS.NewAnimation("federico", "idle_front_1"),
	}

	item.Dimension.Offset = mosaic.Vector{X: 0, Y: -6}
	item.Animation.VariantMax = 6
	item.Animation.Speed = 50
	item.Animation.Sprites = map[string][]string{
		"idle_front":      core.SpriteArray("idle_front", 6),
		"idle_front_side": core.SpriteArray("idle_front_side", 6),
		"idle_back":       core.SpriteArray("idle_back", 6),
		"idle_back_side":  core.SpriteArray("idle_back_side", 6),
		"move_front":      core.SpriteArray("move_front", 6),
		"move_front_side": core.SpriteArray("move_front_side", 6),
		"move_back":       core.SpriteArray("move_back", 6),
		"move_back_side":  core.SpriteArray("move_back_side", 6),
	}

	item, err := world.NewItem(item)
	if err != nil {
		log.Error().Err(err).Msg("unable to create federico")
	}

	return item.Entity
}
