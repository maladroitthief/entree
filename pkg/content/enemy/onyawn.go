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

	duration := time.Millisecond * 500
	frequency := time.Millisecond * 20
	rootNode := onyawnBehaviorTree(world, entity, duration, frequency)
	ai := world.ECS.NewAI(world.Context, rootNode)
	ai.Targets = ai.Targets.Set(core.Human)

	position := world.ECS.NewPosition(0, 0, 1.6)
	movement := world.ECS.NewMovement()
	movement.MaxVelocity = 40

	dimension := world.ECS.NewDimension(
		mosaic.Vector{X: position.X, Y: position.Y},
		mosaic.Vector{X: 16, Y: 16},
	)
	dimension.Offset = mosaic.Vector{X: 0, Y: -6}
	collider := world.ECS.NewCollider(1.0)
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

	world.AI.Add(bt.NewTicker(ai.Context, duration, ai.Node))

	return entity
}

func onyawnBehaviorTree(
	world *content.World,
	entity core.Entity,
	duration time.Duration,
	frequency time.Duration,
) bt.Node {
	// idleMovement := func() (bt.Tick, []bt.Node) {
	// 	log.Info().Msg("idling")
	// 	x := rand.Float64()
	// 	if x < 0.5 {
	// 		x -= 1.0
	// 	}
	// 	y := rand.Float64()
	// 	if y < 0.5 {
	// 		y -= 1.0
	// 	}
	// 	return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
	// 		core.MoveX(world.ECS, x)(entity)
	// 		core.MoveY(world.ECS, y)(entity)
	// 		return bt.Success, nil
	// 	}), nil
	// }

	// 	following := func() bt.Tick {
	// 		return follow(world, entity)
	// 	}
	// moving := func() bt.Tick {
	// 	return bt.Retryer(
	// 		time.Millisecond*2000,
	// 		time.Millisecond*10,
	// 		move(world, entity),
	// 	)
	// }

	return bt.New(
		bt.Selector,
		bt.New(
			bt.Retryer(
				time.Millisecond*1000,
				time.Millisecond*20,
				move(world, entity),
			),
		),
		bt.New(follow(world, entity)),
		bt.New(search(world, entity, 5)),
	)

	// return bt.New(
	// 	bt.Selector,
	// 	bt.New(
	// 		bt.Switch,
	// 		bt.New(
	// 			bt.Async(
	// 				search(world, entity, 3),
	// 			),
	// 		),
	// 		bt.New(
	// 			bt.Sequence,
	// 			bt.New(
	// 				follow(world, entity),
	// 			),
	// 			bt.New(
	// 				bt.Retryer(
	// 					time.Millisecond*5000,
	// 					time.Millisecond*20,
	// 					move(world, entity),
	// 				),
	// 			),
	// 		),
	// 	),
	// 	// idleMovement,
	// )
}
