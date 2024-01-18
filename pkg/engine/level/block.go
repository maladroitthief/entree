package level

import (
	"math/rand"

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
	AddPlayer(p core.Entity, x, y float64)
	AddSolid(x, y float64)
	AddSolid50(x, y float64)
	AddObstacle(x, y float64)
	AddEnemy(x, y float64)
}

type blockFactory struct {
	ecs *core.ECS
}

func NewBlockFactory(ecs *core.ECS) BlockFactory {
	bf := &blockFactory{
		ecs: ecs,
	}

	return bf
}

func (bf *blockFactory) AddPlayer(p core.Entity, x, y float64) {
	// TODO: Handle this error
	position, _ := bf.ecs.GetPosition(p.Id)
	position.X = x
	position.Y = y

	bf.ecs.SetPosition(position)
}

func (bf *blockFactory) AddSolid(x, y float64) {
	environment.Wall(bf.ecs, x, y)
}

func (bf *blockFactory) AddSolid50(x, y float64) {
	roll := rand.Intn(100)
	if roll < 50 {
		environment.Wall(bf.ecs, x, y)
	}
}

func (bf *blockFactory) AddObstacle(x, y float64) {
	roll := rand.Intn(100)

	if roll < 10 {
		environment.Weeds(bf.ecs, x, y)
	} else {
		environment.Grass(bf.ecs, x, y)
	}
}

func (bf *blockFactory) AddEnemy(x, y float64) {
	enemy.NewOnyawn(bf.ecs, x, y)
}
