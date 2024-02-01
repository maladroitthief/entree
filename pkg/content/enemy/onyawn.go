package enemy

import (
	"time"

	"github.com/maladroitthief/entree/common/data"
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/rs/zerolog/log"
)

func NewOnyawn(world *content.World) core.Entity {
	entity := world.ECS.NewEntity()
	state := world.ECS.NewState()
	faction := world.ECS.NewFaction(core.Vegetable)

	duration := time.Millisecond * 200
	frequency := time.Millisecond * 10
	rootNode := onyawnBehaviorTree(world, entity.Id, duration, frequency)
	ai := world.ECS.NewAI(world.Context, rootNode)

	position := world.ECS.NewPosition(0, 0, 1.6)
	movement := world.ECS.NewMovement()
	dimension := world.ECS.NewDimension(
		data.Vector{X: position.X, Y: position.Y},
		data.Vector{X: 16, Y: 16},
	)
	dimension.Offset = data.Vector{X: 0, Y: -6}
	collider := world.ECS.NewCollider(1.0)

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
	id data.GenerationalIndex,
	duration time.Duration,
	frequency time.Duration,
) bt.Node {
	search := func() bt.Tick {
		type index struct {
			x int
			y int
		}
		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

		return func(children []bt.Node) (bt.Status, error) {
			position, err := world.ECS.GetPosition(id)
			if err != nil {
				log.Debug().Err(err).Any("entityId", id).Msg("search error")
				return bt.Failure, nil
			}

			frontier := data.NewQueue[index]()
			x, y := world.Grid.Location(position.X, position.Y)

			start := index{x, y}
			frontier.Enqueue(start)
			cameFrom := map[index]index{}
			cameFrom[start] = start

			for frontier.Len() > 0 {
				current, err := frontier.Dequeue()
				if err != nil {
					return bt.Failure, err
				}

				entities := world.Grid.GetItemsAtLocation(current.x, current.y)
				for _, entity := range entities {
					faction, err := world.ECS.GetFaction(entity.Id)
					if err != nil {
						continue
					}

					if faction.IsArchetype(core.Human) {
						log.Debug().Msg("HUMAN FOUND")
						return bt.Success, nil
					}
				}

				_, ok := cameFrom[current]
				if current.x != start.x && current.y != start.y && ok {
					continue
				}

				for _, direction := range directions {
					next := index{current.x + direction[0], current.y + direction[1]}
					if next.x < 0 || next.x >= world.Grid.SizeX {
						continue
					}
					if next.y < 0 || next.y >= world.Grid.SizeY {
						continue
					}

					_, ok := cameFrom[next]
					if !ok {
						frontier.Enqueue(next)
						cameFrom[next] = current
					}
				}
			}

			return bt.Success, nil
		}
	}

	searching := func() (bt.Tick, []bt.Node) {
		return search(), nil
	}

	moveUp := func() (bt.Tick, []bt.Node) {
		return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
			core.MoveUp(world.ECS)(id)
			return bt.Success, nil
		}), nil
	}
	moveDown := func() (bt.Tick, []bt.Node) {
		return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
			core.MoveDown(world.ECS)(id)
			return bt.Success, nil
		}), nil
	}
	moveLeft := func() (bt.Tick, []bt.Node) {
		return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
			core.MoveLeft(world.ECS)(id)
			return bt.Success, nil
		}), nil
	}
	moveRight := func() (bt.Tick, []bt.Node) {
		return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
			core.MoveRight(world.ECS)(id)
			return bt.Success, nil
		}), nil
	}

	return bt.New(
		bt.Shuffle(bt.Sequence, nil),
		moveUp, moveRight, moveDown, moveLeft, searching,
	)
}
