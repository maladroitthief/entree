package canvas

type Canvas struct {
  entities []*Entity
}

func NewCanvas() *Canvas {
  return &Canvas{}
}

func (c *Canvas) Draw() []*Entity {
  return c.entities
}
