package enemy

import (
	"time"

	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
	"github.com/rs/zerolog/log"
)

func NewOnyawn(world *content.World, x, y float64) core.Entity {
	entity := world.ECS.NewEntity("onyawn")
	rootNode := onyawnBehaviorTree(world, entity)

	item := content.WorldItem{
		Entity:   entity,
		AI:       world.ECS.NewAI(world.Context, rootNode),
		State:    world.ECS.NewState(),
		Faction:  world.ECS.NewFaction(core.Vegetable),
		Position: world.ECS.NewPosition(x, y, 1.6),
		Movement: world.ECS.NewMovement(),
		Dimension: world.ECS.NewDimension(
			mosaic.Vector{X: x, Y: y},
			mosaic.Vector{X: 16, Y: 16},
		),
		Collider:  world.ECS.NewCollider(5.0),
		Animation: world.ECS.NewAnimation("onyawn", "idle_front_1"),
	}

	item.AI.Targets = item.AI.Targets.Set(core.Human)
	item.Faction.Archetype = item.Faction.Archetype.Set(core.Plant)
	item.Dimension.Offset = mosaic.Vector{X: 0, Y: -6}
	item.Collider.ColliderType = core.Moveable
	item.Animation.VariantMax = 6
	item.Animation.Speed = 50
	item.Animation.Sprites = map[string][]string{
		"idle_front":      core.SpriteArray("idle_front", 2),
		"idle_front_side": core.SpriteArray("idle_front_side", 2),
		"idle_back":       core.SpriteArray("idle_back", 2),
		"idle_back_side":  core.SpriteArray("idle_front_side", 2),
		"move_front":      core.SpriteArray("move_front", 6),
		"move_front_side": core.SpriteArray("move_front_side", 6),
		"move_back":       core.SpriteArray("move_back", 6),
		"move_back_side":  core.SpriteArray("move_front_side", 6),
	}

	item, err := world.NewItem(item)
	if err != nil {
		log.Error().Err(err).Msg("unable to create onyawn")
	}

	return item.Entity
}

func onyawnBehaviorTree(
	world *content.World,
	entity core.Entity,
) bt.Node {
	return bt.New(
		bt.Selector,
		bt.New(
			bt.Repeater(
				time.Millisecond*1000,
				time.Millisecond*17,
				move(world, entity),
			),
		),
		bt.New(follow(world, entity, 32)),
		bt.New(search(world, entity, 5)),
	)
}
