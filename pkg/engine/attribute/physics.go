package attribute

import "github.com/maladroitthief/entree/common/data"

const (
	BaseMaxVelocity  = 3
	BaseMass         = 10
	BaseAcceleration = 1
)

type Physics struct {
  Id data.GenerationalIndex
	EntityId data.GenerationalIndex

	Position      data.Vector
	ZLevel        int
	Velocity      data.Vector
	MaxVelocity   float64
	Mass          float64
	Acceleration  float64
	DeltaPosition data.Vector
	Size          data.Vector
	Scale         float64
	Offset        data.Vector
	Bounds        data.Rectangle
}

func NewPhysics(position data.Vector, z int, size data.Vector) Physics {
	return Physics{
		Position:     position,
		ZLevel:       z,
		Velocity:     data.Vector{X: 0, Y: 0},
		MaxVelocity:  BaseMaxVelocity,
		Acceleration: BaseAcceleration,
		Mass:         BaseMass,
		Size:         size,
		Scale:        1,
		Offset:       data.Vector{X: 0, Y: 0},
		Bounds:       data.Bounds(position, size),
	}
}
