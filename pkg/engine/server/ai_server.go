package server

import (
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
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
