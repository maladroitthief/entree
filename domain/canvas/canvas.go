package canvas

import "github.com/maladroitthief/entree/domain/physics"

type Canvas struct {
	entities []Entity
	x        int
	y        int
	size     int
	bounds   [4]physics.Rectangle
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
	c.createBounds()

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

func (c *Canvas) Bounds() []physics.Rectangle {
	return c.bounds[:]
}

func (c *Canvas) createBounds() {
	xSize := float64(c.x * c.size)
	ySize := float64(c.y * c.size)
	size := float64(c.size)

	// North
	c.bounds[0] = physics.NewRectangle(-size, 0, xSize+size, -size)
	// South
	c.bounds[1] = physics.NewRectangle(-size, ySize, xSize+size, ySize+size)
	// East
	c.bounds[2] = physics.NewRectangle(xSize, ySize+size, xSize+size, -size)
	// West
	c.bounds[3] = physics.NewRectangle(-size, ySize+size, 0, -size)
}

func (c *Canvas) OutOfBounds(e Entity) bool {
	for _, bounds := range c.bounds {
		if bounds.Intersects(e.Bounds()) {
			return true
		}
	}

	return false
}
