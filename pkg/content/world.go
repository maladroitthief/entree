package content

import (
	"context"
	"errors"
	"time"

	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/lattice"
	"github.com/rs/zerolog/log"
)

type (
	World struct {
		Context context.Context
		ECS     *core.ECS
		AI      bt.Manager
		Grid    *lattice.SpatialGrid[core.Entity]
	}

	WorldItem struct {
		Entity    core.Entity
		AI        core.AI
		Command   core.Command
		State     core.State
		Movement  core.Movement
		Position  core.Position
		Dimension core.Dimension
		Collider  core.Collider
		Animation core.Animation
		Faction   core.Faction
	}
)

func NewWorld(ctx context.Context, ecs *core.ECS, ai bt.Manager, grid *lattice.SpatialGrid[core.Entity]) *World {
	if ctx == nil {
		log.Fatal().Msg("NewWorld Context is nil")
	}

	if ecs == nil {
		log.Fatal().Msg("NewWorld ECS is nil")
	}

	if ai == nil {
		log.Fatal().Msg("NewWorld AI is nil")
	}

	if grid == nil {
		log.Fatal().Msg("NewWorld Grid is nil")
	}

	return &World{
		Context: ctx,
		ECS:     ecs,
		AI:      ai,
		Grid:    grid,
	}
}

func (w *World) NewItem(item WorldItem) (WorldItem, error) {
	if !w.ECS.EntityActive(item.Entity) {
		return item, errors.New("entity not active")
	}

	if w.ECS.AIActive(item.AI) {
		item.Entity = w.ECS.BindAI(item.Entity, item.AI)
		w.AI.Add(bt.NewTicker(item.AI.Context, time.Millisecond*500, item.AI.Node))
	}

	if w.ECS.StateActive(item.State) {
		item.Entity = w.ECS.BindState(item.Entity, item.State)
	}

	if w.ECS.MovementActive(item.Movement) {
		item.Entity = w.ECS.BindMovement(item.Entity, item.Movement)
	}

	if w.ECS.PositionActive(item.Position) {
		item.Entity = w.ECS.BindPosition(item.Entity, item.Position)
	}

	if w.ECS.DimensionActive(item.Dimension) {
		item.Entity = w.ECS.BindDimension(item.Entity, item.Dimension)
	}

	if w.ECS.ColliderActive(item.Collider) {
		item.Entity = w.ECS.BindCollider(item.Entity, item.Collider)
	}

	if w.ECS.AnimationActive(item.Animation) {
		item.Entity = w.ECS.BindAnimation(item.Entity, item.Animation)
	}

	if w.ECS.FactionActive(item.Faction) {
		item.Entity = w.ECS.BindFaction(item.Entity, item.Faction)
	}

	if w.ECS.ColliderActive(item.Collider) && w.ECS.DimensionActive(item.Dimension) {
		w.Grid.Insert(
			lattice.Item[core.Entity]{
				Value:      item.Entity,
				Bounds:     item.Dimension.Bounds(),
				Multiplier: 1 / item.Collider.ImpedingRate,
			},
		)
	}

	return item, nil
}
