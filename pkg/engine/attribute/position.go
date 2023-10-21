package attribute

import "github.com/maladroitthief/entree/common/data"

type Position struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Position data.Vector
}

func NewPosition(position data.Vector) Position {
	return Position{
		Position: position,
	}
}
