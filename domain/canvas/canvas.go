package canvas

type Canvas struct {
  entities []*Entity
}

func NewCanvas() *Canvas {
  return &Canvas{}
}

func (c *Canvas) Entities() []*Entity {
  return c.entities
}
