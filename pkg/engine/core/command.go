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

		state = state.SetIdle()
		ecs.SetState(state)
	}
}

func MoveY(ecs *ECS, value float64) Command {
	return func(entity Entity) {
		entity, err := ecs.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveY entity error")
			return
		}

		state, err := ecs.GetState(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveY state error")
			return
		}

		if state.Check(Dodging) && state.DodgeCounter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveY movement error")
			return
		}

		state = state.Set(Moving)
		if value < 0 {
			state.OrientationY = North
		} else {
			state.OrientationY = South
		}
		movement.Force.Y = value

		ecs.SetState(state)
		ecs.SetMovement(movement)
	}
}

func MoveX(ecs *ECS, value float64) Command {
	return func(entity Entity) {
		entity, err := ecs.GetEntity(entity.Id)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveX entity error")
			return
		}

		state, err := ecs.GetState(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveX state error")
			return
		}
		if state.Check(Dodging) && state.DodgeCounter <= DodgeDuration {
			return
		}

		movement, err := ecs.GetMovement(entity)
		if err != nil {
			log.Debug().Err(err).Any("entity", entity).Msg("MoveX movement error")
			return
		}

		state = state.Set(Moving)
		if value < 0 {
			state.OrientationX = West
		} else {
			state.OrientationX = East
		}
		movement.Force.X = value

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

		if state.Check(Dodging) && state.DodgeCounter <= DodgeDuration {
			return
		}

		state = state.Set(Dodging)
		ecs.SetState(state)
	}
}
