package canvas

import "github.com/maladroitthief/entree/domain/sprite"

type Entity struct {
	Width  int
	Height int
	X      int
	Y      int
	State  string
}

func (e *Entity) Draw() *sprite.Sprite {
	return nil
}
