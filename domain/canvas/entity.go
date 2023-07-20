package canvas

import (
	"github.com/maladroitthief/entree/domain/physics"
)

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
	DefaultMaxVelocity  = 3
	DefaultMass         = 10
	DefaultSpriteSpeed  = 100
	DefaultSize         = 16
	CollisionBuffer     = 0.1
)

type Entity interface {
	Update(*Canvas)
	Send(msg, val string)
	X() float64
	SetX(float64)
	Y() float64
	SetY(float64)
	Z() float64
	SetZ(float64)
	Position() (x, y float64)
	Size() (x, y float64)
	SetSize(x, y float64)
	Offset() (x, y float64)
	SetOffset(x, y float64)
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
	if deltaX == 0 {
		return positionX
	}

	sizeX, sizeY := e.Size()
	newBounds := physics.Bounds(
		physics.Vector{X: positionX, Y: e.Y()},
		physics.Vector{X: sizeX, Y: sizeY},
	)

	if !r.Intersects(newBounds) {
		return positionX
	}

	if deltaX > 0 {
		return r.MinPoint.X - sizeX/2 - CollisionBuffer
	}

	return r.MaxPoint.X + sizeX/2 + CollisionBuffer
}

func collisionVectorY(e Entity, r physics.Rectangle, deltaY, positionY float64) float64 {
	if deltaY == 0 {
		return positionY
	}

	sizeX, sizeY := e.Size()
	newBounds := physics.Bounds(
		physics.Vector{X: e.X(), Y: positionY},
		physics.Vector{X: sizeX, Y: sizeY},
	)

	if !r.Intersects(newBounds) {
		return positionY
	}

	if deltaY > 0 {
		return r.MinPoint.Y - sizeY/2 - CollisionBuffer
	}

	return r.MaxPoint.Y + sizeY/2 + CollisionBuffer
}

type entity struct {
	x            float64
	y            float64
	z            float64
	sizeX        float64
	sizeY        float64
	offsetX      float64
	offsetY      float64
	bounds       physics.Rectangle
	scale        float64
	sheet        string
	sprite       string
	state        string
	stateCounter int
	orientationX OrientationX
	orientationY OrientationY
	components   []Component
	input        InputComponent
	physics      PhysicsComponent
	graphics     GraphicsComponent
}

func NewEntity() Entity {
	e := &entity{
		sizeX:        DefaultSize,
		sizeY:        DefaultSize,
		scale:        DefaultScale,
		stateCounter: 0,
		orientationX: Neutral,
		orientationY: South,
	}

	return e
}

func (e *entity) Update(c *Canvas) {
	e.SetBounds()
	e.InputComponent().Update(e)
	e.PhysicsComponent().Update(e, c)
	e.GraphicsComponent().Update(e)
}

func (e *entity) Send(msg, val string) {
	for _, c := range e.Components() {
		c.Receive(e, msg, val)
	}
}

func (e *entity) SetPosition(x, y float64) {
	e.SetX(x)
	e.SetY(y)
}

func (e *entity) X() float64 {
	return e.x * e.Scale()
}

func (e *entity) SetX(x float64) {
	e.x = x / e.Scale()
}
func (e *entity) Y() float64 {
	return e.y * e.Scale()
}

func (e *entity) SetY(y float64) {
	e.y = y / e.Scale()
}

func (e *entity) Z() float64 {
	return e.z
}

func (e *entity) SetZ(z float64) {
	e.z = z
}

func (e *entity) Position() (x, y float64) {
	return e.X(), e.Y()
}

func (e *entity) Size() (x, y float64) {
	scale := e.Scale()

	return e.sizeX * scale, e.sizeY * scale
}

func (e *entity) SetSize(x, y float64) {
	e.sizeX = x / e.Scale()
	e.sizeY = y / e.Scale()
}

func (e *entity) Offset() (x, y float64) {
	scale := e.Scale()

	return e.offsetX * scale, e.offsetY * scale
}

func (e *entity) SetOffset(x, y float64) {
	e.offsetX = x / e.Scale()
	e.offsetY = y / e.Scale()
}

func (e *entity) Bounds() physics.Rectangle {
	if e.bounds == (physics.Rectangle{}) {
		e.SetBounds()
	}

	return e.bounds
}

func (e *entity) SetBounds() {
	sizeX, sizeY := e.Size()
	e.bounds = physics.Bounds(
		physics.Vector{X: e.X(), Y: e.Y()},
		physics.Vector{X: sizeX, Y: sizeY},
	)
}

func (e *entity) Scale() float64 {
	if e.scale <= 0 {
		e.scale = 1
	}

	return e.scale
}

func (e *entity) SetScale(f float64) {
	e.scale = f
}

func (e *entity) Sheet() string {
	return e.sheet
}

func (e *entity) SetSheet(s string) {
	e.sheet = s
}

func (e *entity) Sprite() string {
	return e.sprite
}

func (e *entity) SetSprite(s string) {
	e.sprite = s
}

func (e *entity) State() string {
	return e.state
}

func (e *entity) SetState(s string) {
	e.state = s
}

func (e *entity) StateCounter() int {
	return e.stateCounter
}

func (e *entity) SetStateCounter(i int) {
	e.stateCounter = i
}

func (e *entity) IncrementStateCounter() {
	e.stateCounter++
}

func (e *entity) OrientationX() OrientationX {
	return e.orientationX
}

func (e *entity) SetOrientationX(o OrientationX) {
	e.orientationX = o
}

func (e *entity) OrientationY() OrientationY {
	return e.orientationY
}

func (e *entity) SetOrientationY(o OrientationY) {
	e.orientationY = o
}

func (e *entity) Components() []Component {
	return []Component{
		e.physics,
		e.graphics,
		e.input,
	}
}

func (e *entity) InputComponent() InputComponent {
	return e.input
}

func (e *entity) SetInputComponent(c InputComponent) {
	e.input = c
}

func (e *entity) PhysicsComponent() PhysicsComponent {
	return e.physics
}

func (e *entity) SetPhysicsComponent(c PhysicsComponent) {
	e.physics = c
}

func (e *entity) GraphicsComponent() GraphicsComponent {
	return e.graphics
}

func (e *entity) SetGraphicsComponent(c GraphicsComponent) {
	e.graphics = c
}
