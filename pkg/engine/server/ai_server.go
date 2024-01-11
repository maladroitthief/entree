package server

import (
	"fmt"

	"github.com/maladroitthief/entree/common/data"
	behaviortree "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/rs/zerolog/log"
)

type AIServer struct {
}

func NewAIServer() *AIServer {
	s := &AIServer{}

	return s
}

func (s *AIServer) Update(e *core.ECS, inputs []core.Input) {
	ais := e.GetAllAI()

	for _, ai := range ais {
		switch ai.BehaviorType {
		case core.Player:
			ProcessInput(e, ai, inputs)
		case core.Computer:
			ProcessBehavior(e, ai)
		}
	}
}

type aiAttributes struct {
	movement core.Movement
	state    core.State
}

func ProcessInput(e *core.ECS, ai core.AI, inputs []core.Input) {
	state, err := e.GetState(ai.EntityId)
	if err != nil {
		return
	}

	movement, err := e.GetMovement(ai.EntityId)
	if err != nil {
		return
	}

	attr := aiAttributes{movement: movement, state: state}

	movementInputs := []core.Input{}
	actionInputs := []core.Input{}
	for _, input := range inputs {
		switch input {
		case core.MoveUp:
			movementInputs = append(movementInputs, core.MoveUp)
		case core.MoveDown:
			movementInputs = append(movementInputs, core.MoveDown)
		case core.MoveRight:
			movementInputs = append(movementInputs, core.MoveRight)
		case core.MoveLeft:
			movementInputs = append(movementInputs, core.MoveLeft)
		case core.Dodge:
			actionInputs = append(actionInputs, core.Dodge)
		}
	}

	attr = handleActionInputs(actionInputs, attr)
	attr = handleMovementInputs(movementInputs, attr)

	e.SetMovement(attr.movement)
	e.SetState(attr.state)
}

func handleActionInputs(inputs []core.Input, a aiAttributes) aiAttributes {
	for _, input := range inputs {
		switch input {
		case core.Dodge:
			a.state.State = core.Dodging
		}
	}

	return a
}

func handleMovementInputs(inputs []core.Input, a aiAttributes) aiAttributes {
	if a.movement.Acceleration.X == 0 && a.movement.Acceleration.Y != 0 {
		a.state.OrientationX = core.Neutral
	}

	if a.state.State == core.Dodging &&
		a.state.Counter <= core.DodgeDuration {
		return a
	}

	a.movement.Acceleration.X, a.movement.Acceleration.Y = 0, 0

	if len(inputs) == 0 {
		a.state.State = core.Idling
	}

	for _, input := range inputs {
		switch input {
		case core.MoveUp:
			a.state.State = core.Moving
			a.state.OrientationY = core.North
			a.movement.Acceleration.Y = -1
		case core.MoveDown:
			a.state.State = core.Moving
			a.state.OrientationY = core.South
			a.movement.Acceleration.Y = 1
		case core.MoveRight:
			a.state.State = core.Moving
			a.state.OrientationX = core.East
			a.movement.Acceleration.X = 1
		case core.MoveLeft:
			a.state.State = core.Moving
			a.state.OrientationX = core.West
			a.movement.Acceleration.X = -1
		}
	}

	return a
}

func ProcessBehavior(e *core.ECS, ai core.AI) {
	behavior, err := e.GetBehavior(ai.ActiveBehavior)
	if err != nil {
		log.Warn().Err(err).Any("ActiveBehavior", ai.ActiveBehavior).Msg("Active behavior error")
		behavior, err = e.GetBehavior(ai.RootBehavior)
		if err != nil {
			log.Warn().Err(err).Any("RootBehavior", ai.RootBehavior).Msg("Root behavior error")
			return
		}
	}

	if behavior.Status != core.RUNNING {
		behavior.Status = core.RUNNING
	}

	activeTickBehavior, err := behavior.Tick(e)
	if err != nil {
		log.Warn().Err(err).Any("activeBehavior", ai.ActiveBehavior).Msg("Active behavior tick error")
		return
	}

	if status != core.RUNNING {
		ai.ActiveBehavior = behavior.Parent
		e.SetAI(ai)
	}
}
