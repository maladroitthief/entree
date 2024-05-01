package enemy

import (
	"errors"
	"math"

	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/rs/zerolog/log"
)

func search(world *content.World, entity core.Entity, depth int) bt.Tick {
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

		_, err = world.ECS.GetEntity(ai.TargetEntityId)
		if err == nil {
			log.Debug().Msg("search end, target found")
			return bt.Success, nil
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

					return errSuccess
				}
			}

			return nil
		}

		err = world.Grid.Search(
			position.X,
			position.Y,
			depth,
			findTarget,
		)

		if errors.Is(err, errSuccess) {
			log.Debug().Msg("search end, target found")
			return bt.Success, nil
		}

		return bt.Failure, nil
	}
}

func follow(world *content.World, entity core.Entity) bt.Tick {
	return func(children []bt.Node) (bt.Status, error) {
		entity, err := world.ECS.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("follow error")
			return bt.Failure, nil
		}

		ai, err := world.ECS.GetAI(entity)
		if err != nil {
			log.Debug().Err(err).Any("ai", ai).Msg("follow error")
			return bt.Failure, nil
		}

		position, err := world.ECS.GetPosition(entity)
		if err != nil {
			log.Debug().Err(err).Any("position", position).Msg("follow error")
			return bt.Failure, nil
		}

		target, err := world.ECS.GetEntity(ai.TargetEntityId)
		if err != nil {
			log.Debug().Err(err).Any("target", target).Msg("follow error")
			return bt.Failure, nil
		}

		targetPosition, err := world.ECS.GetPosition(target)
		ai.TargetLocation = targetPosition.Vector()
		world.ECS.SetAI(ai)

		if err != nil {
			log.Debug().Err(err).Any("target position", targetPosition).Msg("follow error")
			return bt.Failure, nil
		}

		path, err := world.Grid.WeightedSearch(
			position.Vector(),
			targetPosition.Vector(),
		)

		if err != nil {
			log.Debug().Err(err).Msg("follow error")
			return bt.Failure, nil
		}

		if len(path) == 0 {
			log.Debug().Msg("follow error: path length zero")
			return bt.Failure, nil
		}

		ai.TargetLocation = targetPosition.Vector()
		ai.PathToTarget = path[1:]
		world.ECS.SetAI(ai)

		return bt.Success, nil
	}
}

func move(world *content.World, entity core.Entity) bt.Tick {
	return func(children []bt.Node) (bt.Status, error) {
		entity, err := world.ECS.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("move error")
			return bt.Failure, err
		}

		ai, err := world.ECS.GetAI(entity)
		if err != nil {
			log.Debug().Err(err).Any("ai", ai).Msg("move error")
			return bt.Failure, err
		}

		position, err := world.ECS.GetPosition(entity)
		if err != nil {
			log.Debug().Err(err).Any("position", position).Msg("move error")
			return bt.Failure, err
		}

		if len(ai.PathToTarget) <= 0 {
			log.Debug().Any("path to target", ai.PathToTarget).Msg("path empty")
			return bt.Failure, nil
		}

		to := ai.PathToTarget[0]
		for math.Abs(position.Vector().Distance(to)) <= 1 {
			ai.PathToTarget = ai.PathToTarget[1:]
			if len(ai.PathToTarget) <= 0 {
				log.Debug().Any("path to target", ai.PathToTarget).Msg("path empty")
				return bt.Failure, nil
			}

			to = ai.PathToTarget[0]
		}
		world.ECS.SetAI(ai)

		if math.Abs(position.Vector().Distance(to)) <= 1 && len(ai.PathToTarget) > 1 {
			log.Debug().Any("path to target", ai.PathToTarget).Msg("already there")

			ai.PathToTarget = ai.PathToTarget[1:]
			world.ECS.SetAI(ai)

			return bt.Running, nil
		}

		direction := to.Subtract(position.Vector()).Normalize()
		core.MoveX(world.ECS, direction.X)(entity)
		core.MoveY(world.ECS, direction.Y)(entity)

		return bt.Running, nil
	}
}
