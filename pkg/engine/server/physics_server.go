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
	bounds      [4]attribute.Dimension
	spatialHash *data.SpatialHash[attribute.Position]
}

func NewPhysicsServer(x, y, size float64) *PhysicsServer {
	s := &PhysicsServer{
		x:           x,
		y:           y,
		size:        size,
		spatialHash: data.NewSpatialHash[attribute.Position](32, 32),
	}

	xSize := x * size
	ySize := y * size
	// North
	s.bounds[0] = attribute.NewDimension(
		data.Vector{X: xSize / 2, Y: -size / 2},
		data.Vector{X: xSize, Y: size},
	)
	// South
	s.bounds[1] = attribute.NewDimension(
		data.Vector{X: xSize / 2, Y: ySize + size/2},
		data.Vector{X: xSize, Y: size},
	)
	// East
	s.bounds[2] = attribute.NewDimension(
		data.Vector{X: xSize + size/2, Y: ySize / 2},
		data.Vector{X: size, Y: ySize},
	)
	// West
	s.bounds[3] = attribute.NewDimension(
		data.Vector{X: -size / 2, Y: ySize / 2},
		data.Vector{X: size, Y: ySize},
	)

	return s
}

func (s *PhysicsServer) Load(e *core.ECS) {
	s.spatialHash.Drop()
	positions := e.GetAllPosition()

	for _, p := range positions {
		s.spatialHash.Insert(p, data.Vector{X: p.X, Y: p.Y})
	}
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

func DeltaPosition(p attribute.Position, m attribute.Movement) data.Vector {
	return data.Vector{X: p.X, Y: p.Y}.Add(m.Velocity)
}

func DeltaBounds(p attribute.Position, m attribute.Movement, d attribute.Dimension) data.Rectangle {
	return data.Bounds(DeltaPosition(p, m), d.Size)
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
		deltaP := DeltaPosition(p, m)
		s.spatialHash.Update(p, data.Vector{X: p.X, Y: p.Y}, deltaP)
		p.X = deltaP.X
		p.Y = deltaP.Y
		e.SetPosition(p)
		e.SetMovement(m)
		return
	}

	collisions := s.Collisions(e, p, DeltaBounds(p, m, d))
	for _, oob := range s.bounds[:] {
		p, m = checkOOBCollision(p, m, d, oob)
	}

	if len(collisions) == 0 {
		deltaP := DeltaPosition(p, m)
		s.spatialHash.Update(p, data.Vector{X: p.X, Y: p.Y}, deltaP)
		p.X = deltaP.X
		p.Y = deltaP.Y
		e.SetPosition(p)
		e.SetMovement(m)
		return
	}

	for _, collision := range collisions {
		p, m = HandleCollision(e, p, m, d, c, collision)
	}

	deltaP := DeltaPosition(p, m)
	s.spatialHash.Update(p, data.Vector{X: p.X, Y: p.Y}, deltaP)
	p.X = deltaP.X
	p.Y = deltaP.Y
	e.SetPosition(p)
	e.SetMovement(m)
	return
}

func (s *PhysicsServer) Collisions(
	e *core.ECS,
	p attribute.Position,
	r data.Rectangle,
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

		if r.Intersects(_d.Bounds) {
			results = append(results, _d)
		}
	}

	return results
}

func checkOOBCollision(
	p attribute.Position,
	m attribute.Movement,
	d attribute.Dimension,
	_d attribute.Dimension,
) (attribute.Position, attribute.Movement) {
	if m.Acceleration.X == 0 && m.Acceleration.Y == 0 {
		return p, m
	}

	xCollision := _d.Bounds.Intersects(
		data.Bounds(
			data.Vector{X: DeltaPosition(p, m).X, Y: p.Y},
			d.Size,
		),
	)
	yCollision := _d.Bounds.Intersects(
		data.Bounds(
			data.Vector{X: p.X, Y: DeltaPosition(p, m).Y},
			d.Size,
		),
	)

	if xCollision && m.Acceleration.X > 0 {
		p.X = _d.Bounds.MinPoint.X - d.Size.X/2 - CollisionBuffer
		m.Velocity.X = 0
	} else if xCollision && m.Acceleration.X < 0 {
		p.X = _d.Bounds.MaxPoint.X + d.Size.X/2 + CollisionBuffer
		m.Velocity.X = 0
	}

	if yCollision && m.Acceleration.Y > 0 {
		p.Y = _d.Bounds.MinPoint.Y - d.Size.Y/2 - CollisionBuffer
		m.Velocity.Y = 0
	} else if yCollision && m.Acceleration.Y < 0 {
		p.Y = _d.Bounds.MaxPoint.Y + d.Size.Y/2 + CollisionBuffer
		m.Velocity.Y = 0
	}

	return p, m
}

func HandleCollision(
	e *core.ECS,
	p attribute.Position,
	m attribute.Movement,
	d attribute.Dimension,
	c attribute.Collider,
	_d attribute.Dimension,
) (attribute.Position, attribute.Movement) {

	_c, err := e.GetCollider(_d.EntityId)
	if err != nil {
		return p, m
	}

	xCollision := _d.Bounds.Intersects(
		data.Bounds(
			data.Vector{X: DeltaPosition(p, m).X, Y: p.Y},
			d.Size,
		),
	)
	yCollision := _d.Bounds.Intersects(
		data.Bounds(
			data.Vector{X: p.X, Y: DeltaPosition(p, m).Y},
			d.Size,
		),
	)

	switch _c.ColliderType {
	case attribute.Immovable:
		if xCollision && m.Acceleration.X > 0 {
			p.X = _d.Bounds.MinPoint.X - d.Size.X/2 - CollisionBuffer
			m.Velocity.X = 0
		} else if xCollision && m.Acceleration.X < 0 {
			p.X = _d.Bounds.MaxPoint.X + d.Size.X/2 + CollisionBuffer
			m.Velocity.X = 0
		}

		if yCollision && m.Acceleration.Y > 0 {
			p.Y = _d.Bounds.MinPoint.Y - d.Size.Y/2 - CollisionBuffer
			m.Velocity.Y = 0
		} else if yCollision && m.Acceleration.Y < 0 {
			p.Y = _d.Bounds.MaxPoint.Y + d.Size.Y/2 + CollisionBuffer
			m.Velocity.Y = 0
		}
	case attribute.Impeding:
		if xCollision {
			m.Velocity = m.Velocity.ScaleX(1 - _c.ImpedingRate)
		}
		if yCollision {
			m.Velocity = m.Velocity.ScaleY(1 - _c.ImpedingRate)
		}
	case attribute.Moveable:
	}

	return p, m
}
