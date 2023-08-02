package server

import (
	"github.com/maladroitthief/entree/pkg/engine/attribute"
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
		switch ai.Behavior {
		case attribute.Input:
			ProcessInput(e, ai, inputs)
		}
	}
}

func ProcessInput(e *core.ECS, ai attribute.AI, inputs []core.Input) {
	state, err := e.GetState(ai.EntityId)
	if err != nil {
		return
	}

	physics, err := e.GetPhysics(ai.EntityId)
	if err != nil {
		return
	}

	if physics.Acceleration.X == 0 && physics.Acceleration.Y != 0 {
		state.OrientationX = attribute.Neutral
	}
	physics.Acceleration.X, physics.Acceleration.Y = 0, 0

	for _, input := range inputs {
		switch input {
		case core.MoveUp:
			state.State = "move"
			state.OrientationY = attribute.North
			physics.Acceleration.Y = -1
		case core.MoveDown:
			state.State = "move"
			state.OrientationY = attribute.South
			physics.Acceleration.Y = 1
		case core.MoveRight:
			state.State = "move"
			state.OrientationX = attribute.East
			physics.Acceleration.X = 1
		case core.MoveLeft:
			state.State = "move"
			state.OrientationX = attribute.West
			physics.Acceleration.X = -1
		}
	}

  e.SetPhysics(physics)
	e.SetState(state)
}
