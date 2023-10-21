package level

import (
	"math/rand"

	"github.com/maladroitthief/entree/pkg/content/environment"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

const (
	BlockSize  = 16
	Player     = '@'
	EmptySpace = '0'
	Solid      = '1'
	Solid50    = '2'
	Obstacle   = '5'
)

type BlockFactory interface {
	AddPlayer(e *core.ECS, p core.Entity, x, y float64)
	AddSolid(e *core.ECS, x, y float64)
	AddSolid50(e *core.ECS, x, y float64)
	AddObstacle(e *core.ECS, x, y float64)
}

type blockFactory struct {
}

func NewBlockFactory() BlockFactory {
	bf := &blockFactory{}

	return bf
}

func (bf *blockFactory) AddPlayer(e *core.ECS, p core.Entity, x, y float64) {
	// TODO: Handle this error
	position, _ := e.GetPosition(p.Id)
	position.Position.X = x
	position.Position.Y = y

	e.SetPosition(position)
}

func (bf *blockFactory) AddSolid(e *core.ECS, x, y float64) {
	environment.Wall(e, x, y)
}

func (bf *blockFactory) AddSolid50(e *core.ECS, x, y float64) {
	roll := rand.Intn(100)
	if roll < 50 {
		environment.Wall(e, x, y)
	}
}

func (bf *blockFactory) AddObstacle(e *core.ECS, x, y float64) {
	roll := rand.Intn(100)

	if roll < 10 {
		environment.Weeds(e, x, y)
	} else {
		environment.Grass(e, x, y)
	}
}
