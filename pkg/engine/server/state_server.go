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
		if state.State == core.Idling {
			state.Counter = 0
		} else {
			state.Counter++
		}
		e.SetState(state)
	}
}
