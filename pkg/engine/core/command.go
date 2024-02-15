package core

import (
	"github.com/rs/zerolog/log"
)

type (
	Command func(entity Entity)
)

func Idle(ecs *ECS) Command {
	return func(entity Entity) {
		entity, err := ecs.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("Idle entity error")
			return
		}

		state, err := ecs.GetState(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("Idle state error")
			return
		}

		state.State = Idling
		ecs.SetState(state)
	}
}

func MoveUp(ecs *ECS) Command {
	return func(entity Entity) {
		entity, err := ecs.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveUp entity error")
			return
		}

		state, err := ecs.GetState(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveUp state error")
			return
		}

		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveUp movement error")
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
	return func(entity Entity) {
		entity, err := ecs.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveDown entity error")
			return
		}

		state, err := ecs.GetState(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveDown state error")
			return
		}
		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveDown movement error")
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
	return func(entity Entity) {
		entity, err := ecs.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveLeft entity error")
			return
		}

		state, err := ecs.GetState(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveLeft state error")
			return
		}
		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveLeft movement error")
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
	return func(entity Entity) {
		entity, err := ecs.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveRight entity error")
			return
		}

		state, err := ecs.GetState(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveRight state error")
			return
		}

		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveRight movement error")
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
	return func(entity Entity) {
		entity, err := ecs.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("Dodge entity error")
			return
		}

		state, err := ecs.GetState(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("Dodge state error")
			return
		}

		if state.State == Dodging && state.Counter <= DodgeDuration {
			return
		}

		state.State = Dodging
		ecs.SetState(state)
	}
}
