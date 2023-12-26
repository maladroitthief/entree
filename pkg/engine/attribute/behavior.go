package attribute

import (
	"github.com/maladroitthief/entree/common/data"
)

type Status int
type Archetype int

const (
	RUNNING Status = iota
	SUCCESS
	FAILURE

	COMPOSITE_NODE Archetype = iota
	DECORATOR_NODE
	LEAF_NODE
)



type Behavior struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Status    Status
	Archetype Archetype
	_Parent   *Behavior
	Parent    data.GenerationalIndex
	Children  []data.GenerationalIndex

	Tick func() (Status, error)
}

func Inverter() Behavior {
	b := Behavior{
		Status:    SUCCESS,
		Archetype: DECORATOR_NODE,
		Children:  make(),
	}

	inverter := func() (Status, error) {

	}

	return b
}
