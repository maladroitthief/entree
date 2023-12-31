package level

import (
	"math/rand"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/content/enemy"
	"github.com/maladroitthief/entree/pkg/content/environment"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

const (
	Player     = '@'
	EmptySpace = '0'
	Solid      = '1'
	Solid50    = '2'
	Obstacle   = '5'
	Enemy      = 'e'
)

type BlockFactory interface {
	AddPlayer(e *core.ECS, p core.Entity, x, y float64)
	AddSolid(e *core.ECS, x, y float64)
	AddSolid50(e *core.ECS, x, y float64)
	AddObstacle(e *core.ECS, x, y float64)
	AddEnemy(e *core.ECS, x, y float64)
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
	dimension, _ := e.GetDimension(p.Id)
	position.X = x
	position.Y = y
	dimension.Polygon = dimension.Polygon.SetPosition(data.Vector{X: x, Y: y})

	e.SetPosition(position)
	e.SetDimension(dimension)
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

func (bf *blockFactory) AddEnemy(e *core.ECS, x, y float64) {
	enemy.NewOnyawn(e, x, y)
}
