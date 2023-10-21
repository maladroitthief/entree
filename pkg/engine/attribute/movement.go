package attribute

import "github.com/maladroitthief/entree/common/data"

const (
	BaseMaxVelocity  = 3
	BaseMass         = 10
	BaseImpedingRate = 0.35
	MaxImpedingRate  = 1.0
)

type Movement struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Velocity      data.Vector
	MaxVelocity   float64
	Mass          float64
	Acceleration  data.Vector
}

func NewMovement() Movement {
	return Movement{
		Velocity:      data.Vector{X: 0, Y: 0},
		MaxVelocity:   BaseMaxVelocity,
		Mass:          BaseMass,
	}
}
