package canvas

import (
	"github.com/maladroitthief/entree/domain/physics/collision"
)

type OrientationX int
type OrientationY int

const (
	Neutral OrientationX = iota
	West
	East
	South OrientationY = iota
	North
	DefaultScale        = 3
	DefaultAcceleration = 1.5
	DefaultMaxVelocity  = 5
	DefaultMass         = 10
	DefaultSpriteSpeed  = 40
)

type Entity interface {
	Update(*Canvas)
	Send(msg, val string)
	Position() collision.Vector
	SetPosition(collision.Vector)
	Size() collision.Vector
	SetSize(collision.Vector)
	Bounds() collision.Rectangle
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

func CollisionVector(e, ce Entity, deltaPosition, position collision.Vector) collision.Vector {
	newPosition := position

	// Set the X position
	if deltaPosition.X != 0 {
		newBounds := collision.NewRectangle(
			position.X,
			e.Bounds().MinPoint.Y,
			position.X+e.Size().X,
			e.Bounds().MaxPoint.Y,
		)
		if ce.Bounds().Intersects(newBounds) {
			if deltaPosition.X > 0 {
				newPosition.X = ce.Bounds().MinPoint.X - e.Size().X - 1
			} else {
				newPosition.X = ce.Bounds().MaxPoint.X + 1
			}
		}
	}

	// Set the Y position
	if deltaPosition.Y != 0 {
		newBounds := collision.NewRectangle(
			e.Bounds().MinPoint.X,
			position.Y,
			e.Bounds().MaxPoint.X,
			position.Y+e.Size().Y,
		)
		if ce.Bounds().Intersects(newBounds) {
			if deltaPosition.Y > 0 {
				newPosition.Y = ce.Bounds().MinPoint.Y - e.Size().Y - 1
			} else {
				newPosition.Y = ce.Bounds().MaxPoint.Y + 1
			}
		}
	}

	return newPosition
}
