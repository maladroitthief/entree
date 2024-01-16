package ui

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func ProcessPlayerGameInputs(ecs *core.ECS, playerId data.GenerationalIndex, inputs []core.Input) {
	if len(inputs) == 0 {
		core.Idle(ecs)(playerId)
		return
	}

	for _, input := range inputs {
		switch input {
		case core.InputMoveUp:
			core.MoveUp(ecs)(playerId)
		case core.InputMoveDown:
			core.MoveDown(ecs)(playerId)
		case core.InputMoveRight:
			core.MoveRight(ecs)(playerId)
		case core.InputMoveLeft:
			core.MoveLeft(ecs)(playerId)
		case core.InputDodge:
			core.Dodge(ecs)(playerId)
		}
	}
}
