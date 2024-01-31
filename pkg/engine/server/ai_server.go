package server

import (
	bt "github.com/maladroitthief/entree/common/data/behavior_tree"
	"github.com/maladroitthief/entree/pkg/content"
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

func BFS(world *content.World, entity, target core.Entity) {

}
