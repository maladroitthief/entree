package attribute

import "github.com/maladroitthief/entree/common/data"

type OrientationX int
type OrientationY int

const (
	Neutral OrientationX = iota
	West
	East
	South OrientationY = iota
	North
)

type State struct {
  Id data.GenerationalIndex
	EntityId data.GenerationalIndex

	State        string
	Counter      int
	OrientationX OrientationX
	OrientationY OrientationY
}

func NewState() State {
	return State{
		State:        "idle",
		Counter:      0,
		OrientationX: Neutral,
		OrientationY: South,
	}
}
