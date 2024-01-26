package level

import (
	"math/rand"

	"github.com/maladroitthief/entree/pkg/content"
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
	world *content.World
}

func NewBlockFactory(world *content.World) BlockFactory {
	bf := &blockFactory{
		world: world,
	}

	return bf
}

func (bf *blockFactory) AddPlayer(p core.Entity, x, y float64) {
	// TODO: Handle this error
	position, _ := bf.world.ECS.GetPosition(p.Id)
	position.X = x
	position.Y = y

	bf.world.ECS.SetPosition(position)
}

func (bf *blockFactory) AddSolid(x, y float64) {
	environment.Wall(bf.world.ECS, x, y)
}

func (bf *blockFactory) AddSolid50(x, y float64) {
	roll := rand.Intn(100)
	if roll < 50 {
		environment.Wall(bf.world.ECS, x, y)
	}
}

func (bf *blockFactory) AddObstacle(x, y float64) {
	roll := rand.Intn(100)

	if roll < 10 {
		environment.Weeds(bf.world.ECS, x, y)
	} else {
		environment.Grass(bf.world.ECS, x, y)
	}
}

func (bf *blockFactory) AddEnemy(x, y float64) {
	e := enemy.NewOnyawn(bf.world)

	position, _ := bf.world.ECS.GetPosition(e.Id)
	position.X = x
	position.Y = y

	bf.world.ECS.SetPosition(position)
}
