package server

import (
	"math"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

const (
	CollisionBuffer = 0.1
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

func (s *PhysicsServer) Update(e *core.ECS) {
	s.spatialHash.Drop()
	physics := e.GetAllPhysics()

	for _, p := range physics {
		s.spatialHash.Insert(p, p.Bounds)
	}

	for _, p := range physics {
		p = UpdateVelocity(p)
		p = s.UpdatePosition(e, p)
		e.SetPhysics(p)
	}
}

func (s *PhysicsServer) Collisions(
	p attribute.Physics,
	r data.Rectangle,
) []attribute.Physics {

	results := []attribute.Physics{}
	candidates := s.spatialHash.SearchNeighbors(p.Position.X, p.Position.Y)
	for _, candidate := range candidates {
		if p == candidate {
			continue
		}

		if r.Intersects(candidate.Bounds) {
			results = append(results, candidate)
		}
	}

	return results
}

func UpdateVelocity(p attribute.Physics) attribute.Physics {
	p.Velocity = p.Velocity.ScaleXY(p.DeltaPosition.X, p.DeltaPosition.Y)
	if math.Signbit(p.DeltaPosition.X) != math.Signbit(p.Velocity.X) {
		p.Velocity.X = 0
	}

	if math.Signbit(p.DeltaPosition.Y) != math.Signbit(p.Velocity.Y) {
		p.Velocity.Y = 0
	}

	p.Velocity = p.Velocity.Add(p.DeltaPosition.Scale(p.Acceleration * p.Mass))
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

func (s *PhysicsServer) UpdatePosition(
	e *core.ECS,
	p attribute.Physics,
) attribute.Physics {

	nextPosition := p.Position.Add(p.Velocity)
	nextBounds := data.Bounds(nextPosition, p.Size)

	collisions := s.Collisions(p, nextBounds)

	for _, oob := range s.bounds[:] {
		nextPosition = data.Vector{
			X: UpdatePositionX(p, oob, nextPosition.X),
			Y: UpdatePositionY(p, oob, nextPosition.Y),
		}
	}

	if len(collisions) == 0 {
		p.Position = nextPosition

		return p
	}

	for _, ce := range collisions {
		nextPosition = data.Vector{
			X: UpdatePositionX(p, ce.Bounds, nextPosition.X),
			Y: UpdatePositionY(p, ce.Bounds, nextPosition.Y),
		}
	}

	p.Position = nextPosition

	return p
}

func UpdatePositionX(
	p attribute.Physics,
	r data.Rectangle,
	nextPositionX float64,
) float64 {
	if p.DeltaPosition.X == 0 {
		return nextPositionX
	}

	newBounds := data.Bounds(
		data.Vector{X: nextPositionX, Y: p.Position.Y},
		p.Size,
	)

	if !r.Intersects(newBounds) {
		return nextPositionX
	}

	if p.DeltaPosition.X > 0 {
		return r.MinPoint.X - p.Size.X/2 - CollisionBuffer
	}

	return r.MaxPoint.X + p.Size.X/2 + CollisionBuffer
}

func UpdatePositionY(
	p attribute.Physics,
	r data.Rectangle,
	nextPositionY float64,
) float64 {
	if p.DeltaPosition.Y == 0 {
		return nextPositionY
	}

	newBounds := data.Bounds(
		data.Vector{X: p.Position.X, Y: nextPositionY},
		p.Size,
	)

	if !r.Intersects(newBounds) {
		return nextPositionY
	}

	if p.DeltaPosition.Y > 0 {
		return r.MinPoint.Y - p.Size.Y/2 - CollisionBuffer
	}

	return r.MaxPoint.Y + p.Size.Y/2 + CollisionBuffer
}
