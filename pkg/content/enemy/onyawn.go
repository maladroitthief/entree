package enemy

import (
	"time"

	"github.com/maladroitthief/entree/common/data"
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func NewOnyawn(world *content.World) core.Entity {
	entity := world.ECS.NewEntity()
	state := world.ECS.NewState()

	duration := time.Millisecond * 200
	frequency := time.Millisecond * 10
	rootNode := onyawnBehaviorTree(world.ECS, entity.Id, duration, frequency)
	ai := world.ECS.NewAI(world.Context, rootNode)

	position := world.ECS.NewPosition(0, 0, 1.6)
	movement := world.ECS.NewMovement()
	dimension := world.ECS.NewDimension(
		data.Vector{X: position.X, Y: position.Y},
		data.Vector{X: 16, Y: 16},
	)
	dimension.Offset = data.Vector{X: 0, Y: -6}
	collider := world.ECS.NewCollider()

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
	entity = world.ECS.BindPosition(entity, position)
	entity = world.ECS.BindMovement(entity, movement)
	entity = world.ECS.BindDimension(entity, dimension)
	entity = world.ECS.BindCollider(entity, collider)
	entity = world.ECS.BindAnimation(entity, animation)

	world.AI.Add(bt.NewTicker(ai.Context, duration, ai.Node))

	return entity
}

func onyawnBehaviorTree(ecs *core.ECS, id data.GenerationalIndex, duration, frequency time.Duration) bt.Node {
	moveUp := func() (bt.Tick, []bt.Node) {
		return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
			core.MoveUp(ecs)(id)
			return bt.Success, nil
		}), nil
	}
	moveDown := func() (bt.Tick, []bt.Node) {
		return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
			core.MoveDown(ecs)(id)
			return bt.Success, nil
		}), nil
	}
	moveLeft := func() (bt.Tick, []bt.Node) {
		return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
			core.MoveLeft(ecs)(id)
			return bt.Success, nil
		}), nil
	}
	moveRight := func() (bt.Tick, []bt.Node) {
		return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
			core.MoveRight(ecs)(id)
			return bt.Success, nil
		}), nil
	}

	return bt.New(
		bt.Shuffle(bt.Sequence, nil),
		moveUp, moveRight, moveDown, moveLeft,
	)
}
