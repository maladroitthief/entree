package enemy

import (
	"time"

	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
)

func NewOnyawn(world *content.World) core.Entity {
	entity := world.ECS.NewEntity("onyawn")

	state := world.ECS.NewState()
	faction := world.ECS.NewFaction(core.Vegetable)
	faction.Archetype.Set(core.Plant)

	rootNode := onyawnBehaviorTree(world, entity)
	ai := world.ECS.NewAI(world.Context, rootNode)
	ai.Targets = ai.Targets.Set(core.Human)

	position := world.ECS.NewPosition(0, 0, 1.6)
	movement := world.ECS.NewMovement()

	dimension := world.ECS.NewDimension(
		mosaic.Vector{X: position.X, Y: position.Y},
		mosaic.Vector{X: 16, Y: 16},
	)
	dimension.Offset = mosaic.Vector{X: 0, Y: -6}
	collider := world.ECS.NewCollider(5.0)
	collider.ColliderType = core.Moveable

	animation := world.ECS.NewAnimation("onyawn", "idle_front_1")
	animation.VariantMax = 6
	animation.Speed = 50
	animation.Sprites = map[string][]string{
		"idle_front":      core.SpriteArray("idle_front", 2),
		"idle_front_side": core.SpriteArray("idle_front_side", 2),
		"idle_back":       core.SpriteArray("idle_back", 2),
		"idle_back_side":  core.SpriteArray("idle_front_side", 2),
		"move_front":      core.SpriteArray("move_front", 6),
		"move_front_side": core.SpriteArray("move_front_side", 6),
		"move_back":       core.SpriteArray("move_back", 6),
		"move_back_side":  core.SpriteArray("move_front_side", 6),
	}

	entity = world.ECS.BindAI(entity, ai)
	entity = world.ECS.BindState(entity, state)
	entity = world.ECS.BindFaction(entity, faction)
	entity = world.ECS.BindPosition(entity, position)
	entity = world.ECS.BindMovement(entity, movement)
	entity = world.ECS.BindDimension(entity, dimension)
	entity = world.ECS.BindCollider(entity, collider)
	entity = world.ECS.BindAnimation(entity, animation)

	world.AI.Add(bt.NewTicker(ai.Context, time.Millisecond*500, ai.Node))

	return entity
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
