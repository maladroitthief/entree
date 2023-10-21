package attribute

import "github.com/maladroitthief/entree/common/data"

type ColliderType int

const (
	Immovable ColliderType = iota
	Moveable
	Impeding
)

type Collider struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	ColliderType ColliderType
	ImpedingRate  float64
}

func NewCollider() Collider {
	return Collider{
		ColliderType: Moveable,
		ImpedingRate:  BaseImpedingRate,
	}
}
