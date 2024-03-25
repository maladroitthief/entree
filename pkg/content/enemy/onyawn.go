package enemy

import (
	"errors"
	"time"

	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
	"github.com/rs/zerolog/log"
)

func NewOnyawn(world *content.World) core.Entity {
	entity := world.ECS.NewEntity()
	state := world.ECS.NewState()
	faction := world.ECS.NewFaction(core.Vegetable)

	duration := time.Millisecond * 200
	frequency := time.Millisecond * 10
	rootNode := onyawnBehaviorTree(world, entity, duration, frequency)
	ai := world.ECS.NewAI(world.Context, rootNode)
	ai.Targets = ai.Targets.Set(core.Human)

	position := world.ECS.NewPosition(0, 0, 1.6)
	movement := world.ECS.NewMovement()
	dimension := world.ECS.NewDimension(
		mosaic.Vector{X: position.X, Y: position.Y},
		mosaic.Vector{X: 16, Y: 16},
	)
	dimension.Offset = mosaic.Vector{X: 0, Y: -6}
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
	entity core.Entity,
	duration time.Duration,
	frequency time.Duration,
) bt.Node {
	search := func() bt.Tick {
		return func(children []bt.Node) (bt.Status, error) {
			entity, err := world.ECS.GetEntity(entity.Id)
			if err != nil {
				log.Debug().Err(err).Any("entity", entity).Msg("search error")
				return bt.Failure, nil
			}

			ai, err := world.ECS.GetAI(entity)
			if err != nil {
				log.Debug().Err(err).Any("ai", ai).Msg("search error")
				return bt.Failure, nil
			}

			position, err := world.ECS.GetPosition(entity)
			if err != nil {
				log.Debug().Err(err).Any("position", position).Msg("search error")
				return bt.Failure, nil
			}

			errSuccess := errors.New("target found")
			findTarget := func(entities []core.Entity) error {
				for _, e := range entities {
					faction, err := world.ECS.GetFaction(e)
					if err != nil {
						continue
					}

					if ai.Targets.Check(faction.Archetype) {
						ai.TargetEntityId = e.Id
						world.ECS.SetAI(ai)

						log.Debug().Msg("Target acquired")
						return errSuccess
					}
				}

				return nil
			}

			err = world.Grid.Search(
				position.X,
				position.Y,
				3,
				findTarget,
			)

			if errors.Is(err, errSuccess) {
				return bt.Success, nil
			}

			// if err != nil {
			// 	return bt.Failure, err
			// }

			return bt.Failure, nil
		}
	}

	searching := func() (bt.Tick, []bt.Node) {
		return search(), nil
	}

	// moveUp := func() (bt.Tick, []bt.Node) {
	// 	return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
	// 		core.MoveUp(world.ECS)(entity)
	// 		return bt.Success, nil
	// 	}), nil
	// }
	// moveDown := func() (bt.Tick, []bt.Node) {
	// 	return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
	// 		core.MoveDown(world.ECS)(entity)
	// 		return bt.Success, nil
	// 	}), nil
	// }
	// moveLeft := func() (bt.Tick, []bt.Node) {
	// 	return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
	// 		core.MoveLeft(world.ECS)(entity)
	// 		return bt.Success, nil
	// 	}), nil
	// }
	// moveRight := func() (bt.Tick, []bt.Node) {
	// 	return bt.Repeater(duration, frequency, func(children []bt.Node) (bt.Status, error) {
	// 		core.MoveRight(world.ECS)(entity)
	// 		return bt.Success, nil
	// 	}), nil
	// }

	return bt.New(
		bt.Shuffle(bt.Sequence, nil),
		// moveUp, moveRight, moveDown, moveLeft, searching,
		searching,
	)
}
