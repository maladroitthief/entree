package enemy

import (
	"context"
	"time"

	"github.com/maladroitthief/entree/common/data"
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func NewOnyawn(ecs *core.ECS, x, y float64) core.Entity {
	entity := ecs.NewEntity()
	state := ecs.NewState()

	ai := ecs.NewAI(core.Computer)
	ai.Ticker = OnyawnBehaviorTree(ecs.Context, ecs, entity.Id)

	position := ecs.NewPosition(x, y, 1.6)
	movement := ecs.NewMovement()
	dimension := ecs.NewDimension(
		data.Vector{X: position.X, Y: position.Y},
		data.Vector{X: 16, Y: 16},
	)
	dimension.Offset = data.Vector{X: 0, Y: -6}
	collider := ecs.NewCollider()

	animation := ecs.NewAnimation("onyawn", "idle_front_1")
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

	entity = ecs.BindAI(entity, ai)
	entity = ecs.BindState(entity, state)
	entity = ecs.BindPosition(entity, position)
	entity = ecs.BindMovement(entity, movement)
	entity = ecs.BindDimension(entity, dimension)
	entity = ecs.BindCollider(entity, collider)
	entity = ecs.BindAnimation(entity, animation)

	return entity
}

func OnyawnBehaviorTree(ctx context.Context, ecs *core.ECS, id data.GenerationalIndex) bt.Ticker {
	duration := time.Millisecond * 10

	moveUp := func() (bt.Tick, []bt.Node) {
		return func(children []bt.Node) (bt.Status, error) {
			core.MoveUp(ecs)(id)
			return bt.Success, nil
		}, nil
	}
	moveDown := func() (bt.Tick, []bt.Node) {
		return func(children []bt.Node) (bt.Status, error) {
			core.MoveDown(ecs)(id)
			return bt.Success, nil
		}, nil
	}
	moveLeft := func() (bt.Tick, []bt.Node) {
		return func(children []bt.Node) (bt.Status, error) {
			core.MoveLeft(ecs)(id)
			return bt.Success, nil
		}, nil
	}
	moveRight := func() (bt.Tick, []bt.Node) {
		return func(children []bt.Node) (bt.Status, error) {
			core.MoveRight(ecs)(id)
			return bt.Success, nil
		}, nil
	}

	var root bt.Node = func() (bt.Tick, []bt.Node) {
		return bt.Sequence, []bt.Node{
			moveUp, moveRight, moveDown, moveLeft,
		}
	}

	return bt.NewTicker(ctx, duration, root)
}
