package level

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/canvas/environment"
	"github.com/maladroitthief/entree/domain/physics"
)

const (
	BlockSize  = 16
	Player     = 'P'
	EmptySpace = '0'
	Solid      = '1'
	Solid50    = '2'
)

type BlockFactory interface {
	AddWall(c *canvas.Canvas, x, y float64)
	AddPlayer(c *canvas.Canvas, p canvas.Entity, x, y float64)
}

type blockFactory struct {
}

func NewBlockFactory() BlockFactory {
	bf := &blockFactory{}

	return bf
}

func (bf *blockFactory) AddWall(c *canvas.Canvas, x, y float64) {
	c.AddEntity(environment.Wall(x, y))
}

func (bf *blockFactory) AddPlayer(c *canvas.Canvas, p canvas.Entity, x, y float64) {
	p.SetPosition(physics.Vector{X: x, Y: y})
	c.AddEntity(p)
}
