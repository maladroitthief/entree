package server

import (
	"math"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

const (
	CollisionBuffer = 0.5
)

type PhysicsServer struct {
	x           float64
	y           float64
	size        float64
	bounds      [4]data.Rectangle
	spatialHash *data.SpatialHash[attribute.Physics]
}

func NewPhysicsServer(x, y, size float64) *PhysicsServer {
	s := &PhysicsServer{
		x:           x,
		y:           y,
		size:        size,
		spatialHash: data.NewSpatialHash[attribute.Physics](32, 32),
	}

	xSize := x * size
	ySize := y * size
	// North
	s.bounds[0] = data.NewRectangle(-size, 0, xSize+size, -size)
	// South
	s.bounds[1] = data.NewRectangle(-size, ySize, xSize+size, ySize+size)
	// East
	s.bounds[2] = data.NewRectangle(xSize, ySize+size, xSize+size, -size)
	// West
	s.bounds[3] = data.NewRectangle(-size, ySize+size, 0, -size)

	return s
}

func (s *PhysicsServer) Load(e *core.ECS) {
	s.spatialHash.Drop()
	physics := e.GetAllPhysics()

	for _, p := range physics {
		s.spatialHash.Insert(p, p.Bounds)
	}
}

func (s *PhysicsServer) Update(e *core.ECS) {
	s.Load(e)
	physics := e.GetAllPhysics()

	for _, p := range physics {
		p = UpdateVelocity(p)
		p = s.UpdatePosition(e, p)
		e.SetPhysics(p)
	}
}

func UpdateVelocity(p attribute.Physics) attribute.Physics {
	p.Velocity = p.Velocity.ScaleXY(p.Acceleration.X, p.Acceleration.Y)
	if math.Signbit(p.Acceleration.X) != math.Signbit(p.Velocity.X) {
		p.Velocity.X = 0
	}

	if math.Signbit(p.Acceleration.Y) != math.Signbit(p.Velocity.Y) {
		p.Velocity.Y = 0
	}

	p.Velocity = p.Velocity.Add(p.Acceleration.Scale(p.Mass))
	direction := data.Vector{X: 1, Y: 1}

	if p.Velocity.X < 0 {
		direction.X = -1
	}

	if p.Velocity.Y < 0 {
		direction.Y = -1
	}

	if math.Abs(p.Velocity.X) > p.MaxVelocity {
		p.Velocity.X = p.MaxVelocity
	}

	if math.Abs(p.Velocity.Y) > p.MaxVelocity {
		p.Velocity.Y = p.MaxVelocity
	}

	p.Velocity = p.Velocity.ScaleXY(direction.X, direction.Y)

	m := p.Velocity.Magnitude()
	if m > p.MaxVelocity {
		p.Velocity = p.Velocity.Scale(p.MaxVelocity / m)
	}

	return p
}

func DeltaPosition(p attribute.Physics) data.Vector {
	return p.Position.Add(p.Velocity)
}

func DeltaBounds(p attribute.Physics) data.Rectangle {
	return data.Bounds(DeltaPosition(p), p.Size)
}

func (s *PhysicsServer) UpdatePosition(
	e *core.ECS,
	p attribute.Physics,
) attribute.Physics {

	collisions := s.Collisions(p, DeltaBounds(p))
	for _, oob := range s.bounds[:] {
		p = CheckOOBCollision(p, oob)
	}

	if len(collisions) == 0 {
		p.Position = DeltaPosition(p)
		s.spatialHash.Update(p, p.Bounds, DeltaBounds(p))

		return p
	}

	for _, ce := range collisions {
		p = HandleCollision(p, ce)
	}

	p.Position = DeltaPosition(p)
	s.spatialHash.Update(p, p.Bounds, DeltaBounds(p))

	return p
}

func (s *PhysicsServer) Collisions(
	p attribute.Physics,
	r data.Rectangle,
) []attribute.Physics {

	results := []attribute.Physics{}
	candidates := s.spatialHash.SearchNeighbors(p.Position.X, p.Position.Y)
	for i := 0; i < len(candidates); i++ {
		if p == candidates[i] {
			continue
		}

		if r.Intersects(candidates[i].Bounds) {
			results = append(results, candidates[i])
		}
	}

	return results
}

func CheckOOBCollision(
	p attribute.Physics,
	r data.Rectangle,
) attribute.Physics {
	if p.Acceleration.X == 0 && p.Acceleration.Y == 0 {
		return p
	}

	deltaPosition := DeltaPosition(p)
	xBounds := data.Bounds(
		data.Vector{X: deltaPosition.X, Y: p.Position.Y},
		p.Size,
	)
	yBounds := data.Bounds(
		data.Vector{X: p.Position.X, Y: deltaPosition.Y},
		p.Size,
	)
	xCollision := r.Intersects(xBounds)
	yCollision := r.Intersects(yBounds)

	if xCollision && p.Acceleration.X > 0 {
		p.Position.X = r.MinPoint.X - p.Size.X/2 - CollisionBuffer
		p.Velocity.X = 0
	} else if xCollision && p.Acceleration.X < 0 {
		p.Position.X = r.MaxPoint.X + p.Size.X/2 + CollisionBuffer
		p.Velocity.X = 0
	}

	if yCollision && p.Acceleration.Y > 0 {
		p.Position.Y = r.MinPoint.Y - p.Size.Y/2 - CollisionBuffer
		p.Velocity.Y = 0
	} else if yCollision && p.Acceleration.Y < 0 {
		p.Position.Y = r.MaxPoint.Y + p.Size.Y/2 + CollisionBuffer
		p.Velocity.Y = 0
	}

	return p
}

func HandleCollision(
	p attribute.Physics,
	ce attribute.Physics,
) attribute.Physics {

	xCollision := ce.Bounds.Intersects(
		data.Bounds(
			data.Vector{X: DeltaPosition(p).X, Y: p.Position.Y},
			p.Size,
		),
	)
	yCollision := ce.Bounds.Intersects(
		data.Bounds(
			data.Vector{X: p.Position.X, Y: DeltaPosition(p).Y},
			p.Size,
		),
	)

	switch ce.CollisionType {
	case attribute.Immovable:
		if xCollision && p.Acceleration.X > 0 {
			p.Position.X = ce.Bounds.MinPoint.X - p.Size.X/2 - CollisionBuffer
			p.Velocity.X = 0
		} else if xCollision && p.Acceleration.X < 0 {
			p.Position.X = ce.Bounds.MaxPoint.X + p.Size.X/2 + CollisionBuffer
			p.Velocity.X = 0
		}

		if yCollision && p.Acceleration.Y > 0 {
			p.Position.Y = ce.Bounds.MinPoint.Y - p.Size.Y/2 - CollisionBuffer
			p.Velocity.Y = 0
		} else if yCollision && p.Acceleration.Y < 0 {
			p.Position.Y = ce.Bounds.MaxPoint.Y + p.Size.Y/2 + CollisionBuffer
			p.Velocity.Y = 0
		}
	case attribute.Impeding:
		if xCollision {
			p.Velocity = p.Velocity.ScaleX(1 - ce.ImpedingRate)
		}
		if yCollision {
			p.Velocity = p.Velocity.ScaleY(1 - ce.ImpedingRate)
		}
	case attribute.Moveable:
	}

	return p
}
