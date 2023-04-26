package canvas

import "github.com/maladroitthief/entree/domain/physics/collision"

type Canvas struct {
	entities []*Entity
	x        int
	y        int
	size     int
	hashMap  collision.Hashmap[*Entity]
}

func NewCanvas(x, y, size int) *Canvas {
	c := &Canvas{
		x:    x,
		y:    y,
		size: size,
	}

	return c
}

func (c *Canvas) AddEntity(e *Entity) {
	c.entities = append(c.entities, e)
}

func (c *Canvas) Entities() []*Entity {
	return c.entities
}
