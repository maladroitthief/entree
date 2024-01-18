package level

import (
	"math/rand"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/content/enemy"
	"github.com/maladroitthief/entree/pkg/content/environment"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/engine/level"
)

type fieldBlocks struct {
	ecs *core.ECS
}

func FieldBlockFactory(ecs *core.ECS) level.BlockFactory {
	bf := &fieldBlocks{
		ecs: ecs,
	}

	return bf
}

func (bf *fieldBlocks) AddPlayer(p core.Entity, x, y float64) {
	// TODO: Handle this error
	position, _ := bf.ecs.GetPosition(p.Id)
	dimension, _ := bf.ecs.GetDimension(p.Id)
	position.X = x
	position.Y = y
	dimension.Polygon = dimension.Polygon.SetPosition(data.Vector{X: x, Y: y})

	bf.ecs.SetPosition(position)
	bf.ecs.SetDimension(dimension)
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
	enemy.NewOnyawn(bf.ecs, x, y)
}
