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
	spatialHash *data.SpatialHash[attribute.Position]
	oob         [4]attribute.Dimension
}

func NewPhysicsServer(x, y, size float64) *PhysicsServer {
	s := &PhysicsServer{
		x:           x,
		y:           y,
		size:        size,
		spatialHash: data.NewSpatialHash[attribute.Position](32, 32),
		oob:         OOB(x, y, size),
	}

	return s
}

func (s *PhysicsServer) Load(e *core.ECS) {
	s.spatialHash.Drop()
	positions := e.GetAllPosition()

	for _, p := range positions {
		s.spatialHash.Insert(p, data.Vector{X: p.X, Y: p.Y})
	}
}

func DeltaPosition(p attribute.Position, v data.Vector) data.Vector {
	return data.Vector{X: p.X, Y: p.Y}.Add(v)
}

func DeltaPositionXY(p attribute.Position, x, y float64) data.Vector {
	return data.Vector{X: p.X, Y: p.Y}.Add(data.Vector{X: x, Y: y})
}

func DeltaBounds(d attribute.Dimension, v data.Vector) data.Polygon {
	return d.Bounds.Add(v)
}

func DeltaBoundsXY(d attribute.Dimension, x, y float64) data.Polygon {
	return d.Bounds.Add(data.Vector{X: x, Y: y})
}

func (s *PhysicsServer) Update(e *core.ECS) {
	s.Load(e)
	movements := e.GetAllMovement()

	for _, m := range movements {
		m = UpdateMovement(m)
		s.UpdatePosition(e, m)
	}
}

func UpdateMovement(m attribute.Movement) attribute.Movement {
	m.Velocity = m.Velocity.ScaleXY(m.Acceleration.X, m.Acceleration.Y)
	if math.Signbit(m.Acceleration.X) != math.Signbit(m.Velocity.X) {
		m.Velocity.X = 0
	}

	if math.Signbit(m.Acceleration.Y) != math.Signbit(m.Velocity.Y) {
		m.Velocity.Y = 0
	}

	m.Velocity = m.Velocity.Add(m.Acceleration.Scale(m.Mass))
	direction := data.Vector{X: 1, Y: 1}

	if m.Velocity.X < 0 {
		direction.X = -1
	}

	if m.Velocity.Y < 0 {
		direction.Y = -1
	}

	if math.Abs(m.Velocity.X) > m.MaxVelocity {
		m.Velocity.X = m.MaxVelocity
	}

	if math.Abs(m.Velocity.Y) > m.MaxVelocity {
		m.Velocity.Y = m.MaxVelocity
	}

	m.Velocity = m.Velocity.ScaleXY(direction.X, direction.Y)

	magnitude := m.Velocity.Magnitude()
	if magnitude > m.MaxVelocity {
		m.Velocity = m.Velocity.Scale(m.MaxVelocity / magnitude)
	}

	return m
}

func (s *PhysicsServer) UpdatePosition(
	e *core.ECS,
	m attribute.Movement,
) {

	p, err := e.GetPosition(m.EntityId)
	if err != nil {
		return
	}

	d, err := e.GetDimension(m.EntityId)
	if err != nil {
		return
	}

	c, err := e.GetCollider(m.EntityId)
	if err != nil {
		s.updateAttributes(e, p, m, d)
		return
	}

	p, m, d = s.HandleOOB(e, p, m, d)

	collisions := s.Collisions(e, p, m, d)
	if len(collisions) == 0 {
		s.updateAttributes(e, p, m, d)
		return
	}

	for _, collision := range collisions {
		p, m, d = HandleCollision(e, p, m, d, c, collision)
	}

	s.updateAttributes(e, p, m, d)
	return
}

func (s *PhysicsServer) updateAttributes(
	e *core.ECS,
	p attribute.Position,
	m attribute.Movement,
	d attribute.Dimension,
) {
	deltaPosition := DeltaPosition(p, m.Velocity)
	s.spatialHash.Update(p, data.Vector{X: p.X, Y: p.Y}, deltaPosition)
	p.X = deltaPosition.X
	p.Y = deltaPosition.Y
	d.Bounds = d.Bounds.SetPosition(deltaPosition)

	e.SetPosition(p)
	e.SetMovement(m)
	e.SetDimension(d)
}

func (s *PhysicsServer) Collisions(
	e *core.ECS,
	p attribute.Position,
	m attribute.Movement,
	d attribute.Dimension,
) []attribute.Dimension {

	results := []attribute.Dimension{}
	candidates := s.spatialHash.SearchNeighbors(p.X, p.Y)
	for i := 0; i < len(candidates); i++ {
		if p.EntityId == candidates[i].EntityId {
			continue
		}

		_d, err := e.GetDimension(candidates[i].EntityId)
		if err != nil {
			continue
		}

		_, intersects := DeltaBounds(d, m.Velocity).Intersects(_d.Bounds)
		if intersects {
			results = append(results, _d)
		}
	}

	return results
}

func HandleCollision(
	e *core.ECS,
	p attribute.Position,
	m attribute.Movement,
	d attribute.Dimension,
	c attribute.Collider,
	_d attribute.Dimension,
) (attribute.Position, attribute.Movement, attribute.Dimension) {

	_c, err := e.GetCollider(_d.EntityId)
	if err != nil {
		return p, m, d
	}

	switch _c.ColliderType {
	case attribute.Immovable:
		xMTV, xCollision := _d.Bounds.Intersects(DeltaBoundsXY(d, m.Velocity.X, 0))
		if xCollision && m.Acceleration.X != 0 {
			translation := DeltaPositionXY(p, m.Velocity.X, 0).Add(xMTV)
			p.X = translation.X
			m.Velocity.X = 0
			d.Bounds = d.Bounds.SetPosition(data.Vector{X: p.X, Y: p.Y})
		}

		yMTV, yCollision := _d.Bounds.Intersects(DeltaBoundsXY(d, 0, m.Velocity.Y))
		if yCollision && m.Acceleration.Y != 0 {
			translation := DeltaPositionXY(p, 0, m.Velocity.Y).Add(yMTV)
			p.Y = translation.Y
			m.Velocity.Y = 0
			d.Bounds = d.Bounds.SetPosition(data.Vector{X: p.X, Y: p.Y})
		}
	case attribute.Impeding:
		m.Velocity = m.Velocity.Scale(1 - _c.ImpedingRate)
	case attribute.Moveable:
	}

	return p, m, d
}

func (s *PhysicsServer) HandleOOB(
	e *core.ECS,
	p attribute.Position,
	m attribute.Movement,
	d attribute.Dimension,
) (attribute.Position, attribute.Movement, attribute.Dimension) {

	for _, oob := range s.oob {
		xMTV, xCollision := oob.Bounds.Intersects(DeltaBoundsXY(d, m.Velocity.X, 0))
		if xCollision && m.Acceleration.X != 0 {
			translation := DeltaPositionXY(p, m.Velocity.X, 0).Add(xMTV)
			p.X = translation.X
			m.Velocity.X = 0
			d.Bounds = d.Bounds.SetPosition(data.Vector{X: p.X, Y: p.Y})
		}

		yMTV, yCollision := oob.Bounds.Intersects(DeltaBoundsXY(d, 0, m.Velocity.Y))
		if yCollision && m.Acceleration.Y != 0 {
			translation := DeltaPositionXY(p, 0, m.Velocity.Y).Add(yMTV)
			p.Y = translation.Y
			m.Velocity.Y = 0
			d.Bounds = d.Bounds.SetPosition(data.Vector{X: p.X, Y: p.Y})
		}
	}

	return p, m, d
}

func OOB(x, y, size float64) [4]attribute.Dimension {
	entities := [4]core.Entity{}

	xSize := x * size
	ySize := y * size
	positions := [4]data.Vector{
		{X: xSize / 2, Y: -size / 2},
		{X: xSize / 2, Y: ySize + size/2},
		{X: xSize + size/2, Y: ySize / 2},
		{X: -size / 2, Y: ySize / 2},
	}
	sizes := [4]data.Vector{
		{X: xSize, Y: size},
		{X: xSize, Y: size},
		{X: size, Y: ySize},
		{X: size, Y: ySize},
	}

	dimensions := [4]attribute.Dimension{}
	for i := 0; i < len(entities); i++ {
		dimensions[i] = attribute.NewDimension(positions[i], sizes[i])
	}

	return dimensions
}
