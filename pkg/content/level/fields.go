package level

import (
	"math/rand"

	"github.com/maladroitthief/entree/pkg/content/environment"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/engine/level"
)
type fieldBlocks struct {
}

func FieldBlockFactory() level.BlockFactory {
	bf := &fieldBlocks{}

	return bf
}

func (bf *fieldBlocks) AddPlayer(e *core.ECS, p core.Entity, x, y float64) {
	// TODO: Handle this error
	position, _ := e.GetPosition(p.Id)
	position.X = x
	position.Y = y

	e.SetPosition(position)
}

func (bf *fieldBlocks) AddSolidBlock(e *core.ECS, x, y float64) {
	environment.Grass(e, x, y)
}

func (bf *fieldBlocks) AddSolid(e *core.ECS, x, y float64) {
	environment.Wall(e, x, y)
}

func (bf *fieldBlocks) AddSolid50(e *core.ECS, x, y float64) {
	roll := rand.Intn(100)
	if roll < 50 {
		environment.Wall(e, x, y)
	}
}

func (bf *fieldBlocks) AddObstacle(e *core.ECS, x, y float64) {
	roll := rand.Intn(100)

	if roll < 10 {
		environment.Weeds(e, x, y)
	} else {
		environment.Grass(e, x, y)
	}
}
