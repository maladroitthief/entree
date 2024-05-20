package server

import (
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/lattice"
	"github.com/maladroitthief/mosaic"
)

const (
	// CollisionBuffer is used to avoid sticking to the hidden edge of grouped polygons
	CollisionBuffer = 0.001
)

type (
	PhysicsServer struct {
		x         float64
		y         float64
		size      float64
		ecs       *core.ECS
		grid      *lattice.SpatialGrid[core.Entity]
		gameSpeed float64
	}
	body struct {
		entity         core.Entity
		position       core.Position
		movement       core.Movement
		dimension      core.Dimension
		collider       core.Collider
		startingBounds mosaic.Rectangle
	}
)

func NewPhysicsServer(world *content.World, x, y, size float64) *PhysicsServer {
	s := &PhysicsServer{
		x:         x,
		y:         y,
		size:      size,
		grid:      world.Grid,
		ecs:       world.ECS,
		gameSpeed: 0.05,
	}

	return s
}

func (s *PhysicsServer) Update(ecs *core.ECS) {
	s.ResetGrid()
	movements := ecs.GetAllMovements()

	for _, m := range movements {
		var err error
		body := body{movement: m}

		body.entity, err = ecs.GetEntity(m.EntityId)
		if err != nil {
			continue
		}

		body.position, err = ecs.GetPosition(body.entity)
		if err != nil {
			continue
		}

		body.dimension, err = ecs.GetDimension(body.entity)
		if err != nil {
			continue
		}
		body.dimension.Polygon = body.dimension.Polygon.SetPosition(
			mosaic.Vector{X: body.position.X, Y: body.position.Y},
		)
		body.startingBounds = body.dimension.Polygon.Bounds

		body = s.movementUpdate(body)
		s.collisionUpdate(body)
	}
}

func (s *PhysicsServer) ResetGrid() {
	s.grid.Drop()
	entities := s.ecs.GetAllEntities()
	for _, entity := range entities {
		collider, err := s.ecs.GetCollider(entity)
		if err != nil {
			continue
		}

		dimension, err := s.ecs.GetDimension(entity)
		if err != nil {
			continue
		}
		s.grid.Insert(entity, dimension.Bounds(), 1/collider.ImpedingRate)
	}
}

func (s *PhysicsServer) movementUpdate(body body) body {
	force := body.movement.Force.Scale(body.movement.Mass)
	delta := s.deltaPosition(body.position, force)
	body.position.X, body.position.Y = delta.X, delta.Y
	body.dimension.Polygon = body.dimension.Polygon.SetPosition(delta)

	return body
}

func (s *PhysicsServer) collisionUpdate(body body) {
	var err error
	body.collider, err = s.ecs.GetCollider(body.entity)
	if err != nil {
		s.setBody(body)
		return
	}

	body = s.resolveOutOfBounds(body)
	collisions := s.checkCollisions(body)
	if len(collisions) == 0 {
		s.setBody(body)
		return
	}

	for _, collision := range collisions {
		body = s.resolveCollision(body, collision)
	}
	s.setBody(body)

	return
}

func (s *PhysicsServer) speedMaskVector(x, y float64) mosaic.Vector {
	speedMask := 1.0 - (1.0 - s.gameSpeed)
	return mosaic.NewVector(x, y).Scale(speedMask)
}

func (s *PhysicsServer) deltaPosition(position core.Position, deltaVector mosaic.Vector) mosaic.Vector {
	delta := s.speedMaskVector(deltaVector.X, deltaVector.Y)
	deltaPosition := mosaic.NewVector(position.X, position.Y).Add(delta)

	return deltaPosition
}

func (s *PhysicsServer) setBody(body body) {
	multiplier := 1.0
	calculatedBounds := body.dimension.Bounds()

	collider, err := s.ecs.GetCollider(body.entity)
	if err == nil {
		multiplier = 1 / collider.ImpedingRate
	}
	s.grid.Update(body.entity, body.startingBounds, calculatedBounds, multiplier)

	if body.movement.Force.X == 0 && body.movement.Force.Y != 0 {
		state, err := s.ecs.GetState(body.entity)
		if err == nil {
			state.OrientationX = core.Neutral
			s.ecs.SetState(state)
		}
	}

	body.movement.Force.X, body.movement.Force.Y = 0, 0
	s.ecs.SetPosition(body.position)
	s.ecs.SetMovement(body.movement)
	s.ecs.SetDimension(body.dimension)
}

func (s *PhysicsServer) checkCollisions(body body) []core.Dimension {
	collisions := []core.Dimension{}
	entities := s.grid.FindNear(body.dimension.Bounds())

	for i := 0; i < len(entities); i++ {
		entityDimension, err := s.ecs.GetDimension(entities[i])
		if err != nil {
			continue
		}

		_, depth := body.dimension.Polygon.Intersects(entityDimension.Polygon)
		if depth != 0.0 {
			collisions = append(collisions, entityDimension)
		}
	}

	return collisions
}

func (s *PhysicsServer) resolveCollision(body body, objectDimension core.Dimension) body {
	object, err := s.ecs.GetEntity(objectDimension.EntityId)
	if err != nil || object.Id == body.entity.Id {
		return body
	}

	objectCollider, err := s.ecs.GetCollider(object)
	if err != nil {
		return body
	}

	switch objectCollider.ColliderType {
	case core.Immovable:
		normal, depth := objectDimension.Polygon.Intersects(body.dimension.Polygon)
		deltaP := mosaic.NewVector(body.position.X, body.position.Y).Add(
			normal.Scale(depth + CollisionBuffer),
		)
		body.position.X, body.position.Y = deltaP.X, deltaP.Y
		body.dimension.Polygon = body.dimension.Polygon.SetPosition(deltaP)

	case core.Impeding:
		body.movement.Velocity = body.movement.Velocity.Scale(1 - objectCollider.ImpedingRate)

	case core.Moveable:
		normal, depth := body.dimension.Polygon.Intersects(objectDimension.Polygon)
		objectPosition, err := s.ecs.GetPosition(object)
		if err != nil {
			return body
		}

		deltaP := mosaic.NewVector(body.position.X, body.position.Y).Add(
			normal.Scale(-(depth + CollisionBuffer)),
		)
		body.position.X, body.position.Y = deltaP.X, deltaP.Y
		body.dimension.Polygon = body.dimension.Polygon.SetPosition(deltaP)

		deltaP = mosaic.NewVector(objectPosition.X, objectPosition.Y).Add(
			normal.Scale(depth + CollisionBuffer),
		)
		objectPosition.X, objectPosition.Y = deltaP.X, deltaP.Y
		objectDimension.Polygon = objectDimension.Polygon.SetPosition(deltaP)

		s.ecs.SetPosition(objectPosition)
		s.ecs.SetDimension(objectDimension)
	}

	return body
}

func (s *PhysicsServer) resolveOutOfBounds(body body) body {
	sizeX, sizeY := s.x*s.size, s.y*s.size

	center := mosaic.Vector{X: sizeX / 2, Y: sizeY / 2}
	oob := mosaic.NewRectangle(center, sizeX, sizeY).ToPolygon()

	normal, depth := oob.ContainsPolygon(body.dimension.Polygon)
	newPosition := mosaic.NewVector(body.position.X, body.position.Y).Add(normal.Scale(depth))
	body.position.X, body.position.Y = newPosition.X, newPosition.Y
	body.dimension.Polygon = body.dimension.Polygon.SetPosition(newPosition)

	return body
}
