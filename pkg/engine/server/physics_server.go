package server

import (
	"math"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/rs/zerolog/log"
)

const (
	CollisionBuffer = 0.5
)

type PhysicsServer struct {
	x    float64
	y    float64
	size float64
	grid *data.SpatialGrid[core.Entity]
}

func NewPhysicsServer(world *content.World, x, y, size float64) *PhysicsServer {
	s := &PhysicsServer{
		x:    x,
		y:    y,
		size: size,
		grid: world.Grid,
	}

	return s
}

func (s *PhysicsServer) Load(e *core.ECS) {
	s.grid.Drop()
	entities := e.GetAllEntities()
	for _, entity := range entities {
		dimension, err := e.GetDimension(entity.DimensionId)
		if err != nil {
			continue
		}
		s.grid.Insert(entity, dimension.Bounds())
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
	log.Debug().Msg("PhysicsServer.Update()")
	s.Load(e)
	movements := e.GetAllMovement()

	for _, m := range movements {
		m = s.UpdateMovement(m)
		s.UpdatePosition(e, m)
	}
}

func (s *PhysicsServer) UpdateMovement(m core.Movement) core.Movement {
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

	p, err := e.GetPosition(m.EntityId)
	if err != nil {
		return
	}

	d, err := e.GetDimension(m.EntityId)
	if err != nil {
		return
	}
	d.Polygon = d.Polygon.SetPosition(data.Vector{X: p.X, Y: p.Y})

	c, err := e.GetCollider(m.EntityId)
	if err != nil {
		s.updateAttributes(e, p, m, d)
		return
	}

	p, m, d = s.HandleOutOfBounds(e, p, m, d)

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

	s.grid.Update(entity, oldBounds, d.Bounds())

	if m.Acceleration.X == 0 && m.Acceleration.Y != 0 {
		state, err := e.GetState(p.EntityId)
		if err == nil {
			state.OrientationX = core.Neutral
			e.SetState(state)
		}
	}
	m.Acceleration.X = 0
	m.Acceleration.Y = 0

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
	results := []core.Dimension{}
	entities := s.grid.FindNear(d.Bounds())
	for i := 0; i < len(entities); i++ {
		_d, err := e.GetDimension(entities[i].Id)
		if err != nil {
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

func (s *PhysicsServer) HandleOutOfBounds(
	e *core.ECS,
	p core.Position,
	m core.Movement,
	d core.Dimension,
) (core.Position, core.Movement, core.Dimension) {
	sizeX := s.x * s.size
	sizeY := s.y * s.size
	center := data.Vector{X: sizeX / 2, Y: sizeY / 2}
	oob := data.NewRectangle(center, sizeX, sizeY).ToPolygon()

	xMTV, xContained := oob.ContainsPolygon(DeltaBoundsXY(d, m.Velocity.X, 0))
	if !xContained && m.Acceleration.X != 0 {
		translation := DeltaPositionXY(p, m.Velocity.X, 0).Add(xMTV)
		p.X = translation.X
		m.Velocity.X = 0
		d.Polygon = d.Polygon.SetPosition(data.Vector{X: p.X, Y: p.Y})
	}

	yMTV, yContained := oob.ContainsPolygon(DeltaBoundsXY(d, 0, m.Velocity.Y))
	if !yContained && m.Acceleration.Y != 0 {
		translation := DeltaPositionXY(p, 0, m.Velocity.Y).Add(yMTV)
		p.Y = translation.Y
		m.Velocity.Y = 0
		d.Polygon = d.Polygon.SetPosition(data.Vector{X: p.X, Y: p.Y})
	}

	return p, m, d
}
