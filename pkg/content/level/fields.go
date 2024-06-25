package level

import (
	"math/rand"

	"github.com/rs/zerolog/log"

	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/content/enemy"
	"github.com/maladroitthief/entree/pkg/content/environment"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/engine/level"
)

type fieldBlocks struct {
	world *content.World
}

func FieldBlockFactory(world *content.World) level.BlockFactory {
	if world == nil {
		log.Fatal().Msg("FieldBlockFactory World is nil")
	}

	bf := &fieldBlocks{
		world: world,
	}

	return bf
}

func (bf *fieldBlocks) AddPlayer(entity core.Entity, x, y float64) {
	// TODO: Handle this error
	position, _ := bf.world.ECS.GetPosition(entity)
	position.X = x
	position.Y = y

	bf.world.ECS.SetPosition(position)
}

func (bf *fieldBlocks) AddSolidBlock(x, y float64) {
	environment.Grass(bf.world, x, y)
}

func (bf *fieldBlocks) AddSolid(x, y float64) {
	entity := environment.Wall(bf.world, x, y)
	position, _ := bf.world.ECS.GetPosition(entity)
	position.X = x
	position.Y = y
	bf.world.ECS.SetPosition(position)
}

func (bf *fieldBlocks) AddSolid50(x, y float64) {
	roll := rand.Intn(100)
	if roll < 50 {
		entity := environment.Wall(bf.world, x, y)
		position, _ := bf.world.ECS.GetPosition(entity)
		position.X = x
		position.Y = y
		bf.world.ECS.SetPosition(position)
	}
}

func (bf *fieldBlocks) AddObstacle(x, y float64) {
	roll := rand.Intn(100)

	if roll < 10 {
		environment.Weeds(bf.world, x, y)
	} else {
		environment.Grass(bf.world, x, y)
	}
}

func (bf *fieldBlocks) AddEnemy(x, y float64) {
	enemy.NewOnyawn(bf.world, x, y)

	// 	position, _ := bf.world.ECS.GetPosition(entity)
	// 	position.X = x
	// 	position.Y = y
	// bf.world.ECS.SetPosition(position)
}
