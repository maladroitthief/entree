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
			core.MoveUp(ecs)(player)
		case core.InputMoveDown:
			core.MoveDown(ecs)(player)
		case core.InputMoveRight:
			core.MoveRight(ecs)(player)
		case core.InputMoveLeft:
			core.MoveLeft(ecs)(player)
		case core.InputDodge:
			core.Dodge(ecs)(player)
		}
	}
}
