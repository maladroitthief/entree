package attribute

import "github.com/maladroitthief/entree/common/data"

type CollisionType int

const (
	BaseMaxVelocity  = 3
	BaseMass         = 10
	BaseImpedingRate = 0.35
	MaxImpedingRate  = 1.0

	Immovable CollisionType = iota
	Moveable
	Impeding
)

type Physics struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Position      data.Vector
	ZLevel        int
	Velocity      data.Vector
	MaxVelocity   float64
	Mass          float64
	Acceleration  data.Vector
	Size          data.Vector
	Scale         float64
	Offset        data.Vector
	Bounds        data.Rectangle
	CollisionType CollisionType
	ImpedingRate  float64
}

func NewPhysics(position data.Vector, z int, size data.Vector) Physics {
	return Physics{
		Position:      position,
		ZLevel:        z,
		Velocity:      data.Vector{X: 0, Y: 0},
		MaxVelocity:   BaseMaxVelocity,
		Mass:          BaseMass,
		Size:          size,
		Scale:         1,
		Offset:        data.Vector{X: 0, Y: 0},
		Bounds:        data.Bounds(position, size),
		CollisionType: Moveable,
		ImpedingRate:  BaseImpedingRate,
	}
}
