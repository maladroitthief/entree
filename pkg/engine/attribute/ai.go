package attribute

import "github.com/maladroitthief/entree/common/data"

type BehaviorType int

const (
	None BehaviorType = iota
	Input
)

type AI struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	BehaviorType BehaviorType

	RootBehavior   data.GenerationalIndex
	ActiveBehavior data.GenerationalIndex
	ActiveSequence bool
}

func NewAI(b BehaviorType) AI {
	return AI{
		BehaviorType: b,
	}
}
