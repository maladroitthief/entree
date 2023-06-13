package canvas

import "github.com/maladroitthief/entree/domain/physics"

type OrientationX int
type OrientationY int

const (
	Neutral OrientationX = iota
	West
	East
	South OrientationY = iota
	North
	DefaultScale        = 1
	DefaultAcceleration = 1.5
	DefaultMaxVelocity  = 5
	DefaultMass         = 10
	DefaultSpriteSpeed  = 40
	CollisionBuffer     = 1
)

type Entity interface {
	Update(*Canvas)
	Send(msg, val string)
	Position() physics.Vector
	SetPosition(physics.Vector)
	Size() physics.Vector
	SetSize(physics.Vector)
	Offset() physics.Vector
	Bounds() physics.Rectangle
	SetBounds()
	Scale() float64
	SetScale(float64)
	Sheet() string
	SetSheet(string)
	Sprite() string
	SetSprite(string)
	State() string
	SetState(string)
	StateCounter() int
	SetStateCounter(int)
	IncrementStateCounter()
	OrientationX() OrientationX
	SetOrientationX(OrientationX)
	OrientationY() OrientationY
	SetOrientationY(OrientationY)
	Components() []Component
	SetComponents([]Component)
	InputComponent() InputComponent
	SetInputComponent(InputComponent)
	PhysicsComponent() PhysicsComponent
	SetPhysicsComponent(PhysicsComponent)
	GraphicsComponent() GraphicsComponent
	SetGraphicsComponent(GraphicsComponent)
}

func ResetEntity(e Entity) {
	e.SetState("idle")
	e.Send("Reset", "")
}

func CollisionVector(e Entity, r physics.Rectangle, deltaPosition, position physics.Vector) physics.Vector {
	return physics.Vector{
		X: collisionVectorX(e, r, deltaPosition.X, position.X),
		Y: collisionVectorY(e, r, deltaPosition.Y, position.Y),
	}
}

func collisionVectorX(e Entity, r physics.Rectangle, deltaX, positionX float64) float64 {
	// If we are not moving in X, do nothing
	if deltaX == 0 {
		return positionX
	}

	newBounds := physics.Bounds(
		physics.Vector{X: positionX, Y: e.Position().Y},
		e.Size(),
	)

	// if no physics occurs, allow the movement
	if !r.Intersects(newBounds) {
		return positionX
	}

	// Set the position to the bounds of the object we are colliding with. This step avoids issues where an entity would stop prematurely before contacting another entity
	if deltaX > 0 {
		return r.MinPoint.X - e.Size().X/2 - CollisionBuffer
	}

	return r.MaxPoint.X + e.Size().X/2 + CollisionBuffer
}

func collisionVectorY(e Entity, r physics.Rectangle, deltaY, positionY float64) float64 {
	// If we are not moving in Y, do nothing
	if deltaY == 0 {
		return positionY
	}

	newBounds := physics.Bounds(
		physics.Vector{X: e.Position().X, Y: positionY},
		e.Size(),
	)

	// if no physics occurs, allow the movement
	if !r.Intersects(newBounds) {
		return positionY
	}

	// Set the position to the bounds of the object we are colliding with. This step avoids issues where an entity would stop prematurely before contacting another entity
	if deltaY > 0 {
		return r.MinPoint.Y - e.Size().Y/2 - CollisionBuffer
	}

	return r.MaxPoint.Y + e.Size().Y/2 + CollisionBuffer
}
