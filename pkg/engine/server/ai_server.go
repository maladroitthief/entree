package server

import (
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

type AIServer struct {
	btManager bt.Manager
}

func NewAIServer() *AIServer {
	s := &AIServer{}

	return s
}

func (s *AIServer) Update(e *core.ECS) {
	// ais := e.GetAllAI()

	// for _, ai := range ais {
	// 	switch ai.BehaviorType {
	// 	case core.Computer:
	// 		ProcessInput(e, ai, inputs)
	// 	}
	// }
}
