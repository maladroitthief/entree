package ui

import (
	"github.com/maladroitthief/entree/pkg/engine/core"
)

func ProcessPlayerGameInputs(ecs *core.ECS, player core.Entity, inputs []core.Input) {
	if len(inputs) == 0 {
		core.Idle(ecs)(player)
		return
	}

	for _, input := range inputs {
		switch input {
		case core.InputMoveUp:
			core.MoveY(ecs, -1)(player)
		case core.InputMoveDown:
			core.MoveY(ecs, 1)(player)
		case core.InputMoveLeft:
			core.MoveX(ecs, -1)(player)
		case core.InputMoveRight:
			core.MoveX(ecs, 1)(player)
		case core.InputDodge:
			core.Dodge(ecs)(player)
		}
	}
}
