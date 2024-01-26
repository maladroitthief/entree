package content

import (
	"context"

	"github.com/maladroitthief/entree/common/data"
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/rs/zerolog/log"
)

type World struct {
	Context context.Context
	ECS     *core.ECS
	AI      bt.Manager
	Grid    *data.SpatialGrid[core.Entity]
}

func NewWorld(ctx context.Context, ecs *core.ECS, ai bt.Manager, grid *data.SpatialGrid[core.Entity]) *World {
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
