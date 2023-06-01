package canvas

import "github.com/maladroitthief/entree/domain/physics"

type Canvas struct {
	entities []Entity
	x        int
	y        int
	size     int
	quadTree *physics.QuadTree[Entity]
}

func NewCanvas(x, y, size int) *Canvas {
	c := &Canvas{
		x:    x,
		y:    y,
		size: size,
		quadTree: physics.NewQuadTree[Entity](
			0, physics.NewRectangle(0, 0, float64(x*size), float64(y*size)),
		),
	}

	return c
}

func (c *Canvas) AddEntity(e Entity) {
	c.entities = append(c.entities, e)
}

func (c *Canvas) Update() {
	// dump the quadtree and rebuild it
	c.quadTree.Clear()
	for _, entity := range c.entities {
		c.quadTree.Insert(
			physics.NewQuadTreeItem(
				entity,
				physics.Bounds(entity.Position(), entity.Size()),
			),
		)
	}
}

func (c *Canvas) Entities() []Entity {
	return c.entities
}

func (c *Canvas) Collisions(e Entity, r physics.Rectangle) []Entity {
	results := []Entity{}

	// Broad phase
	candidates := c.quadTree.Get(r)

	// Narrow phase
	for _, candidate := range candidates {
		if e == candidate {
			continue
		}

		if r.Intersects(candidate.Bounds()) {
			results = append(results, candidate)
		}
	}

	return results
}
