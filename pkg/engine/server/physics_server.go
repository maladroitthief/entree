package server

import (
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/lattice"
	"github.com/maladroitthief/mosaic"
)

const (
	CollisionBuffer = 0.5
)

type (
	PhysicsServer struct {
		x         float64
		y         float64
		size      float64
		grid      *lattice.SpatialGrid[core.Entity]
		gameSpeed float64
	}
	physicsAttributes struct {
		entity    core.Entity
		position  core.Position
		movement  core.Movement
		dimension core.Dimension
	}
)

func NewPhysicsServer(world *content.World, x, y, size float64) *PhysicsServer {
	s := &PhysicsServer{
		x:         x,
		y:         y,
		size:      size,
		grid:      world.Grid,
		gameSpeed: 0.05,
	}

	return s
}

func (s *PhysicsServer) Update(ecs *core.ECS) {
	s.Load(ecs)
	movements := ecs.GetAllMovements()

	for _, m := range movements {
		m = s.UpdateMovement(m)
		s.UpdatePosition(ecs, m)
	}
}

func (s *PhysicsServer) Load(ecs *core.ECS) {
	s.grid.Drop()
	entities := ecs.GetAllEntities()
	for _, entity := range entities {
		collider, err := ecs.GetCollider(entity)
		if err != nil {
			continue
		}

		dimension, err := ecs.GetDimension(entity)
		if err != nil {
			continue
		}
		s.grid.Insert(entity, dimension.Bounds(), 1/collider.ImpedingRate)
	}
}

func (s *PhysicsServer) speedMaskVector(x, y float64) mosaic.Vector {
	speedMask := 1.0 - (1.0 - s.gameSpeed)
	return mosaic.NewVector(x, y).Scale(speedMask)
}

func (s *PhysicsServer) DeltaPosition(position core.Position, deltaVector mosaic.Vector) mosaic.Vector {
	return s.DeltaPositionXY(position, deltaVector.X, deltaVector.Y)
}

func (s *PhysicsServer) DeltaPositionXY(position core.Position, x, y float64) mosaic.Vector {
	delta := s.speedMaskVector(x, y)
	deltaPosition := mosaic.NewVector(position.X, position.Y).Add(delta)

	return deltaPosition
}

func (s *PhysicsServer) DeltaBounds(dimension core.Dimension, vector mosaic.Vector) mosaic.Polygon {
	return s.DeltaBoundsXY(dimension, vector.X, vector.Y)
}

func (s *PhysicsServer) DeltaBoundsXY(dimension core.Dimension, x, y float64) mosaic.Polygon {
	delta := s.speedMaskVector(x, y)
	deltaBounds := dimension.Polygon.Add(delta)

	return deltaBounds
}

func (s *PhysicsServer) updateAttributes(ecs *core.ECS, attr physicsAttributes) {
	deltaPosition := s.DeltaPosition(attr.position, attr.movement.Velocity)
	oldBounds := attr.dimension.Polygon.Bounds

	attr.position.X = deltaPosition.X
	attr.position.Y = deltaPosition.Y
	attr.dimension.Polygon = attr.dimension.Polygon.SetPosition(deltaPosition)

	collider, err := ecs.GetCollider(attr.entity)
	if err != nil {
		s.grid.Update(attr.entity, oldBounds, attr.dimension.Bounds(), 1.0)
	} else {
		s.grid.Update(attr.entity, oldBounds, attr.dimension.Bounds(), 1/collider.ImpedingRate)
	}

	if attr.movement.Acceleration.X == 0 && attr.movement.Acceleration.Y != 0 {
		state, err := ecs.GetState(attr.entity)
		if err == nil {
			state.OrientationX = core.Neutral
			ecs.SetState(state)
		}
	}
	attr.movement.Acceleration.X = 0
	attr.movement.Acceleration.Y = 0

	ecs.SetPosition(attr.position)
	ecs.SetMovement(attr.movement)
	ecs.SetDimension(attr.dimension)
}

func (s *PhysicsServer) UpdateMovement(movement core.Movement) core.Movement {
	force := movement.Acceleration.Scale(movement.Mass)
	switch f := mosaic.NewVector(force.X, 0); {
	case f.X < 0:
		if movement.Velocity.X > 0 {
			movement.Velocity.X = -1
		}
		movement.Velocity = movement.Velocity.Add(f)
		if movement.Velocity.X <= -movement.MaxVelocity {
			movement.Velocity.X = -movement.MaxVelocity
		}
	case f.X > 0:
		if movement.Velocity.X < 0 {
			movement.Velocity.X = 1
		}
		movement.Velocity = movement.Velocity.Add(f)
		if movement.Velocity.X >= movement.MaxVelocity {
			movement.Velocity.X = movement.MaxVelocity
		}
	case f.X == 0:
		movement.Velocity.X = 0
	}

	switch f := mosaic.NewVector(0, force.Y); {
	case f.Y < 0:
		if movement.Velocity.Y > 0 {
			movement.Velocity.Y = -1
		}
		movement.Velocity = movement.Velocity.Add(f)
		if movement.Velocity.Y <= -movement.MaxVelocity {
			movement.Velocity.Y = -movement.MaxVelocity
		}
	case f.Y > 0:
		if movement.Velocity.Y < 0 {
			movement.Velocity.Y = 1
		}
		movement.Velocity = movement.Velocity.Add(f)
		if movement.Velocity.Y >= movement.MaxVelocity {
			movement.Velocity.Y = movement.MaxVelocity
		}
	case f.Y == 0:
		movement.Velocity.Y = 0
	}

	magnitude := movement.Velocity.Magnitude()
	if magnitude > movement.MaxVelocity {
		movement.Velocity = movement.Velocity.Scale(movement.MaxVelocity / magnitude)
	}

	return movement
}

func (s *PhysicsServer) UpdatePosition(
	ecs *core.ECS,
	movement core.Movement,
) {
	entity, err := ecs.GetEntity(movement.EntityId)
	if err != nil {
		return
	}

	position, err := ecs.GetPosition(entity)
	if err != nil {
		return
	}

	dimension, err := ecs.GetDimension(entity)
	if err != nil {
		return
	}
	dimension.Polygon = dimension.Polygon.SetPosition(mosaic.Vector{X: position.X, Y: position.Y})
	attr := physicsAttributes{entity, position, movement, dimension}
	collider, err := ecs.GetCollider(entity)
	if err != nil {
		s.updateAttributes(ecs, attr)
		return
	}

	attr = s.HandleOutOfBounds(ecs, attr)

	collisions := s.Collisions(ecs, attr)
	if len(collisions) == 0 {
		s.updateAttributes(ecs, attr)
		return
	}

	for _, collision := range collisions {
		attr = s.HandleCollision(ecs, attr, collider, collision)
	}

	s.updateAttributes(ecs, attr)
	return
}

func (s *PhysicsServer) Collisions(ecs *core.ECS, attr physicsAttributes) []core.Dimension {
	results := []core.Dimension{}
	entities := s.grid.FindNear(attr.dimension.Bounds())
	for i := 0; i < len(entities); i++ {
		_d, err := ecs.GetDimension(entities[i])
		if err != nil {
			continue
		}

		_, intersects := s.DeltaBounds(attr.dimension, attr.movement.Velocity).Intersects(_d.Polygon)
		if intersects {
			results = append(results, _d)
		}
	}

	return results
}

func (s *PhysicsServer) HandleCollision(ecs *core.ECS, attr physicsAttributes, collider core.Collider, collision core.Dimension) physicsAttributes {
	collisionEntity, err := ecs.GetEntity(collision.EntityId)
	if err != nil {
		return attr
	}

	collisionCollider, err := ecs.GetCollider(collisionEntity)
	if err != nil {
		return attr
	}

	switch collisionCollider.ColliderType {
	case core.Immovable:
		xMTV, xCollision := collision.Polygon.Intersects(s.DeltaBoundsXY(attr.dimension, attr.movement.Velocity.X, 0))
		if xCollision && attr.movement.Acceleration.X != 0 {
			translation := s.DeltaPositionXY(attr.position, attr.movement.Velocity.X, 0).Add(xMTV)
			attr.position.X = translation.X
			attr.movement.Velocity.X = 0
			attr.dimension.Polygon = attr.dimension.Polygon.SetPosition(mosaic.Vector{X: attr.position.X, Y: attr.position.Y})
		}

		yMTV, yCollision := collision.Polygon.Intersects(s.DeltaBoundsXY(attr.dimension, 0, attr.movement.Velocity.Y))
		if yCollision && attr.movement.Acceleration.Y != 0 {
			translation := s.DeltaPositionXY(attr.position, 0, attr.movement.Velocity.Y).Add(yMTV)
			attr.position.Y = translation.Y
			attr.movement.Velocity.Y = 0
			attr.dimension.Polygon = attr.dimension.Polygon.SetPosition(mosaic.Vector{X: attr.position.X, Y: attr.position.Y})
		}
	case core.Impeding:
		attr.movement.Velocity = attr.movement.Velocity.Scale(1 - collisionCollider.ImpedingRate)
	case core.Moveable:
	}

	return attr
}

func (s *PhysicsServer) HandleOutOfBounds(ecs *core.ECS, attr physicsAttributes) physicsAttributes {
	sizeX := s.x * s.size
	sizeY := s.y * s.size
	center := mosaic.Vector{X: sizeX / 2, Y: sizeY / 2}
	oob := mosaic.NewRectangle(center, sizeX, sizeY).ToPolygon()

	xMTV, xContained := oob.ContainsPolygon(s.DeltaBoundsXY(attr.dimension, attr.movement.Velocity.X, 0))
	if !xContained && attr.movement.Acceleration.X != 0 {
		translation := s.DeltaPositionXY(attr.position, attr.movement.Velocity.X, 0).Add(xMTV)
		attr.position.X = translation.X
		attr.movement.Velocity.X = 0
		attr.dimension.Polygon = attr.dimension.Polygon.SetPosition(mosaic.Vector{X: attr.position.X, Y: attr.position.Y})
	}

	yMTV, yContained := oob.ContainsPolygon(s.DeltaBoundsXY(attr.dimension, 0, attr.movement.Velocity.Y))
	if !yContained && attr.movement.Acceleration.Y != 0 {
		translation := s.DeltaPositionXY(attr.position, 0, attr.movement.Velocity.Y).Add(yMTV)
		attr.position.Y = translation.Y
		attr.movement.Velocity.Y = 0
		attr.dimension.Polygon = attr.dimension.Polygon.SetPosition(mosaic.Vector{X: attr.position.X, Y: attr.position.Y})
	}

	return attr
}
