package attribute

import "github.com/maladroitthief/entree/common/data"

type Behavior int

const (
	None Behavior = iota
	Input
)

type AI struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Behavior Behavior
}

func NewAI(b Behavior) AI {
	return AI{
		Behavior: b,
	}
}
