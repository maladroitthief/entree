package canvas_test

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
)

type EntityMock struct {
	x            float64
	y            float64
	z            float64
	sizeX        float64
	sizeY        float64
	offsetX      float64
	offsetY      float64
	position     physics.Vector
	size         physics.Vector
	offset       physics.Vector
	bounds       physics.Rectangle
	scale        float64
	sheet        string
	sprite       string
	state        string
	stateCounter int
	orientationX canvas.OrientationX
	orientationY canvas.OrientationY
	components   []canvas.Component
	input        canvas.InputComponent
	physics      canvas.PhysicsComponent
	graphics     canvas.GraphicsComponent
}

func (e *EntityMock) Update(c *canvas.Canvas) {
	e.SetBounds()
	e.InputComponent().Update(e)
	e.PhysicsComponent().Update(e, c)
	e.GraphicsComponent().Update(e)
}

func (e *EntityMock) Send(msg, val string) {
	for _, c := range e.Components() {
		c.Receive(e, msg, val)
	}
}

func (e *EntityMock) SetPosition(x, y float64) {
	e.SetX(x)
	e.SetY(y)
}

func (e *EntityMock) X() float64 {
	return e.x * e.Scale()
}

func (e *EntityMock) SetX(x float64) {
	e.x = x / e.Scale()
}
func (e *EntityMock) Y() float64 {
	return e.y * e.Scale()
}

func (e *EntityMock) SetY(y float64) {
	e.y = y / e.Scale()
}

func (e *EntityMock) Z() float64 {
	return e.z
}

func (e *EntityMock) SetZ(z float64) {
	e.z = z
}

func (e *EntityMock) Position() (x, y float64) {
	return e.X(), e.Y()
}

func (e *EntityMock) Size() (x, y float64) {
	scale := e.Scale()

	return e.sizeX * scale, e.sizeY * scale
}

func (e *EntityMock) SetSize(x, y float64) {
	e.sizeX = x / e.Scale()
	e.sizeY = y / e.Scale()
}

func (e *EntityMock) Offset() (x, y float64) {
	scale := e.Scale()

	return e.offsetX * scale, e.offsetY * scale
}

func (e *EntityMock) SetOffset(x, y float64) {
	e.offsetX = x / e.Scale()
	e.offsetY = y / e.Scale()
}

func (e *EntityMock) Bounds() physics.Rectangle {
	if e.bounds == (physics.Rectangle{}) {
		e.SetBounds()
	}

	return e.bounds
}

func (e *EntityMock) SetBounds() {
	sizeX, sizeY := e.Size()
	e.bounds = physics.Bounds(
		physics.Vector{X: e.X(), Y: e.Y()},
		physics.Vector{X: sizeX, Y: sizeY},
	)
}

func (e *EntityMock) Scale() float64 {
	if e.scale <= 0 {
		e.scale = 1
	}

	return e.scale
}

func (e *EntityMock) SetScale(f float64) {
	e.scale = f
}

func (e *EntityMock) Sheet() string {
	return e.sheet
}

func (e *EntityMock) SetSheet(s string) {
	e.sheet = s
}

func (e *EntityMock) Sprite() string {
	return e.sprite
}

func (e *EntityMock) SetSprite(s string) {
	e.sprite = s
}

func (e *EntityMock) State() string {
	return e.state
}

func (e *EntityMock) SetState(s string) {
	e.state = s
}

func (e *EntityMock) StateCounter() int {
	return e.stateCounter
}

func (e *EntityMock) SetStateCounter(i int) {
	e.stateCounter = i
}

func (e *EntityMock) IncrementStateCounter() {
	e.stateCounter++
}

func (e *EntityMock) OrientationX() canvas.OrientationX {
	return e.orientationX
}

func (e *EntityMock) SetOrientationX(o canvas.OrientationX) {
	e.orientationX = o
}

func (e *EntityMock) OrientationY() canvas.OrientationY {
	return e.orientationY
}

func (e *EntityMock) SetOrientationY(o canvas.OrientationY) {
	e.orientationY = o
}

func (e *EntityMock) Components() []canvas.Component {
	return e.components
}

func (e *EntityMock) InputComponent() canvas.InputComponent {
	return e.input
}

func (e *EntityMock) SetInputComponent(c canvas.InputComponent) {
	e.input = c
}

func (e *EntityMock) PhysicsComponent() canvas.PhysicsComponent {
	return e.physics
}

func (e *EntityMock) SetPhysicsComponent(c canvas.PhysicsComponent) {
	e.physics = c
}

func (e *EntityMock) GraphicsComponent() canvas.GraphicsComponent {
	return e.graphics
}

func (e *EntityMock) SetGraphicsComponent(c canvas.GraphicsComponent) {
	e.graphics = c
}

type InputMock struct{}

func (m *InputMock) Update(canvas.Entity)                  {}
func (m *InputMock) Receive(canvas.Entity, string, string) {}

type PhysicsMock struct{}

func (m *PhysicsMock) Update(canvas.Entity, *canvas.Canvas)  {}
func (m *PhysicsMock) Receive(canvas.Entity, string, string) {}

type GraphicsMock struct{}

func (m *GraphicsMock) Update(canvas.Entity)                  {}
func (m *GraphicsMock) Receive(canvas.Entity, string, string) {}
