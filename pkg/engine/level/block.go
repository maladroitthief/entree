package level

import (
	"github.com/maladroitthief/entree/pkg/content/environment"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

const (
	BlockSize  = 16
	Player     = 'P'
	EmptySpace = '0'
	Solid      = '1'
	Solid50    = '2'
)

type BlockFactory interface {
	AddWall(e *core.ECS, x, y float64)
	AddPlayer(e *core.ECS, p core.Entity, x, y float64)
}

type blockFactory struct {
}

func NewBlockFactory() BlockFactory {
	bf := &blockFactory{}

	return bf
}

func (bf *blockFactory) AddWall(e *core.ECS, x, y float64) {
	environment.Wall(e, x, y)
}

func (bf *blockFactory) AddPlayer(e *core.ECS, p core.Entity, x, y float64) {
	// TODO: Handle this error
	physics, _ := e.GetPhysics(p.Id)
	physics.Position.X = x
	physics.Position.Y = y

	e.SetPhysics(physics)
}
