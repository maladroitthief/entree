package server

import (
	"math"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

const (
	CollisionBuffer = 0.5
)

type PhysicsServer struct {
	log         logs.Logger
	x           float64
	y           float64
	size        float64
	spatialHash *data.SpatialHash[core.Entity]
}

func NewPhysicsServer(e *core.ECS, log logs.Logger, x, y, size float64) *PhysicsServer {
	s := &PhysicsServer{
		log:         log,
		x:           x,
		y:           y,
		size:        size,
		spatialHash: data.NewSpatialHash[core.Entity](int(x), int(y), 33),
	}

	s.log.Debug("NewPhysicsServer()", nil)
	return s
}

func (s *PhysicsServer) Load(e *core.ECS) {
	s.spatialHash.Drop()
	entities := e.GetAllEntities()
	for _, entity := range entities {
		dimension, err := e.GetDimension(entity.DimensionId)
		if err != nil {
			continue
		}
		s.spatialHash.Insert(entity, dimension.Bounds())
	}
}

func DeltaPosition(p core.Position, v data.Vector) data.Vector {
	return data.Vector{X: p.X, Y: p.Y}.Add(v)
}

func DeltaPositionXY(p core.Position, x, y float64) data.Vector {
	return data.Vector{X: p.X, Y: p.Y}.Add(data.Vector{X: x, Y: y})
}

func DeltaBounds(d core.Dimension, v data.Vector) data.Polygon {
	return d.Polygon.Add(v)
}

func DeltaBoundsXY(d core.Dimension, x, y float64) data.Polygon {
	return d.Polygon.Add(data.Vector{X: x, Y: y})
}

func (s *PhysicsServer) Update(e *core.ECS) {
	s.log.Debug("PhysicsServer", "Update()")
	s.Load(e)
	movements := e.GetAllMovement()

	for _, m := range movements {
		m = s.UpdateMovement(m)
		s.UpdatePosition(e, m)
	}
}

func (s *PhysicsServer) UpdateMovement(m core.Movement) core.Movement {
	s.log.Debug("PhysicsServer", "UpdateMovement()")
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
	m core.Movement,
) {

	s.log.Debug("PhysicsServer", "UpdatePosition()")
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
	p core.Position,
	m core.Movement,
	d core.Dimension,
) {
	entity, err := e.GetEntity(p.EntityId)
	if err != nil {
		return
	}

	deltaPosition := DeltaPosition(p, m.Velocity)
	oldBounds := d.Polygon.Bounds

	p.X = deltaPosition.X
	p.Y = deltaPosition.Y
	d.Polygon = d.Polygon.SetPosition(deltaPosition)

	s.spatialHash.Update(entity, oldBounds, d.Bounds())

	e.SetPosition(p)
	e.SetMovement(m)
	e.SetDimension(d)
}

func (s *PhysicsServer) Collisions(
	e *core.ECS,
	p core.Position,
	m core.Movement,
	d core.Dimension,
) []core.Dimension {
	s.log.Debug("Start", "PhysicsServer.Collisions()")
	results := []core.Dimension{}
	entities := s.spatialHash.FindNear(d.Bounds())
	s.log.Debug("entities length", len(entities))
	for i := 0; i < len(entities); i++ {
		_d, err := e.GetDimension(entities[i].Id)
		if err != nil {
			s.log.Error("dimension not found", "PhysicsServer.Collisions()", err)
			continue
		}

		_, intersects := DeltaBounds(d, m.Velocity).Intersects(_d.Polygon)
		if intersects {
			results = append(results, _d)
		}
	}

	return results
}

func HandleCollision(
	e *core.ECS,
	p core.Position,
	m core.Movement,
	d core.Dimension,
	c core.Collider,
	_d core.Dimension,
) (core.Position, core.Movement, core.Dimension) {

	_c, err := e.GetCollider(_d.EntityId)
	if err != nil {
		return p, m, d
	}

	switch _c.ColliderType {
	case core.Immovable:
		xMTV, xCollision := _d.Polygon.Intersects(DeltaBoundsXY(d, m.Velocity.X, 0))
		if xCollision && m.Acceleration.X != 0 {
			translation := DeltaPositionXY(p, m.Velocity.X, 0).Add(xMTV)
			p.X = translation.X
			m.Velocity.X = 0
			d.Polygon = d.Polygon.SetPosition(data.Vector{X: p.X, Y: p.Y})
		}

		yMTV, yCollision := _d.Polygon.Intersects(DeltaBoundsXY(d, 0, m.Velocity.Y))
		if yCollision && m.Acceleration.Y != 0 {
			translation := DeltaPositionXY(p, 0, m.Velocity.Y).Add(yMTV)
			p.Y = translation.Y
			m.Velocity.Y = 0
			d.Polygon = d.Polygon.SetPosition(data.Vector{X: p.X, Y: p.Y})
		}
	case core.Impeding:
		m.Velocity = m.Velocity.Scale(1 - _c.ImpedingRate)
	case core.Moveable:
	}

	return p, m, d
}
