package server

import (
	"github.com/maladroitthief/entree/pkg/engine/core"
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
		case core.InputMoveUp:
			movementInputs = append(movementInputs, core.InputMoveUp)
		case core.InputMoveDown:
			movementInputs = append(movementInputs, core.InputMoveDown)
		case core.InputMoveRight:
			movementInputs = append(movementInputs, core.InputMoveRight)
		case core.InputMoveLeft:
			movementInputs = append(movementInputs, core.InputMoveLeft)
		case core.InputDodge:
			actionInputs = append(actionInputs, core.InputDodge)
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
		case core.InputDodge:
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
		case core.InputMoveUp:
			a.state.State = core.Moving
			a.state.OrientationY = core.North
			a.movement.Acceleration.Y = -1
		case core.InputMoveDown:
			a.state.State = core.Moving
			a.state.OrientationY = core.South
			a.movement.Acceleration.Y = 1
		case core.InputMoveRight:
			a.state.State = core.Moving
			a.state.OrientationX = core.East
			a.movement.Acceleration.X = 1
		case core.InputMoveLeft:
			a.state.State = core.Moving
			a.state.OrientationX = core.West
			a.movement.Acceleration.X = -1
		}
	}

	return a
}
