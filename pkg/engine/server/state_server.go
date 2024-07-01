package server

import (
	"github.com/maladroitthief/entree/pkg/engine/core"
)

type StateServer struct {
}

func NewStateServer() *StateServer {
	s := &StateServer{}

	return s
}

func (s *StateServer) Update(e *core.ECS) {
	states := e.GetAllStates()

	for _, state := range states {
		if state.Idling() {
			state.DodgeCounter = 0
		} else {
			state.DodgeCounter++
		}
		e.SetState(state)
	}
}
