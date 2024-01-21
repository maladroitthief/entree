package level

import (
	"math/rand"

	"github.com/maladroitthief/entree/pkg/content/enemy"
	"github.com/maladroitthief/entree/pkg/content/environment"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/engine/level"
	"github.com/maladroitthief/entree/pkg/engine/server"
)

type fieldBlocks struct {
	ecs *core.ECS
	ai  *server.AIServer
}

func FieldBlockFactory(ecs *core.ECS, ai *server.AIServer) level.BlockFactory {
	bf := &fieldBlocks{
		ecs: ecs,
		ai:  ai,
	}

	return bf
}

func (bf *fieldBlocks) AddPlayer(p core.Entity, x, y float64) {
	// TODO: Handle this error
	position, _ := bf.ecs.GetPosition(p.Id)
	position.X = x
	position.Y = y

	bf.ecs.SetPosition(position)
}

func (bf *fieldBlocks) AddSolidBlock(x, y float64) {
	environment.Grass(bf.ecs, x, y)
}

func (bf *fieldBlocks) AddSolid(x, y float64) {
	environment.Wall(bf.ecs, x, y)
}

func (bf *fieldBlocks) AddSolid50(x, y float64) {
	roll := rand.Intn(100)
	if roll < 50 {
		environment.Wall(bf.ecs, x, y)
	}
}

func (bf *fieldBlocks) AddObstacle(x, y float64) {
	roll := rand.Intn(100)

	if roll < 10 {
		environment.Weeds(bf.ecs, x, y)
	} else {
		environment.Grass(bf.ecs, x, y)
	}
}

func (bf *fieldBlocks) AddEnemy(x, y float64) {
	e := enemy.NewOnyawn(bf.ecs, bf.ai)

	position, _ := bf.ecs.GetPosition(e.Id)
	position.X = x
	position.Y = y

	bf.ecs.SetPosition(position)
}
