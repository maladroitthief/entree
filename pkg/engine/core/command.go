package core

import (
	"github.com/maladroitthief/entree/common/data"
	"github.com/rs/zerolog/log"
)

type (
	Command func(data.GenerationalIndex)
)

func Idle(ecs *ECS) Command {
	return func(entityId data.GenerationalIndex) {
		state, err := ecs.GetState(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("Idle state error")
			return
		}

		state.State = Idling
		ecs.SetState(state)
	}
}

func MoveUp(ecs *ECS) Command {
	return func(entityId data.GenerationalIndex) {
		state, err := ecs.GetState(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("MoveUp state error")
			return
		}

		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("MoveUp movement error")
			return
		}

		state.State = Moving
		state.OrientationY = North
		movement.Acceleration.Y = -1
		ecs.SetState(state)
		ecs.SetMovement(movement)
	}
}

func MoveDown(ecs *ECS) Command {
	return func(entityId data.GenerationalIndex) {
		state, err := ecs.GetState(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("MoveDown state error")
			return
		}
		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("MoveDown movement error")
			return
		}

		state.State = Moving
		state.OrientationY = South
		movement.Acceleration.Y = 1

		ecs.SetState(state)
		ecs.SetMovement(movement)
	}
}

func MoveLeft(ecs *ECS) Command {
	return func(entityId data.GenerationalIndex) {
		state, err := ecs.GetState(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("MoveLeft state error")
			return
		}
		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("MoveLeft movement error")
			return
		}

		state.State = Moving
		state.OrientationX = West
		movement.Acceleration.X = -1

		ecs.SetState(state)
		ecs.SetMovement(movement)
	}
}

func MoveRight(ecs *ECS) Command {
	return func(entityId data.GenerationalIndex) {
		state, err := ecs.GetState(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("MoveRight state error")
			return
		}

		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("MoveRight movement error")
			return
		}

		state.State = Moving
		state.OrientationX = East
		movement.Acceleration.X = 1

		ecs.SetState(state)
		ecs.SetMovement(movement)
	}
}

func Dodge(ecs *ECS) Command {
	return func(entityId data.GenerationalIndex) {
		state, err := ecs.GetState(entityId)
		if err != nil {
			log.Debug().Err(err).Any("entityId", entityId).Msg("Dodge state error")
			return
		}

		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		state.State = Dodging
		ecs.SetState(state)
	}
}
