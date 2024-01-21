package server

import (
	"time"

	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

type AIServer struct {
	btManager bt.Manager
}

func NewAIServer() *AIServer {
	s := &AIServer{
		btManager: bt.NewManager(),
	}

	return s
}

func (s *AIServer) Add(ai core.AI) {
	duration := time.Millisecond * 100
	ticker := bt.NewTicker(ai.Context, duration, ai.Node)
	s.btManager.Add(ticker)
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
