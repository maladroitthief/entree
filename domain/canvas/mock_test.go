package canvas_test

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
)

type EntityMock struct {
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

func (e *EntityMock) Position() physics.Vector {
	return e.position
}

func (e *EntityMock) SetPosition(v physics.Vector) {
	e.position = v
}

func (e *EntityMock) Size() physics.Vector {
	return e.size
}

func (e *EntityMock) SetSize(v physics.Vector) {
	e.size = v
}

func (e *EntityMock) Offset() physics.Vector {
	return e.offset
}

func (e *EntityMock) Bounds() physics.Rectangle {
	return e.bounds
}

func (e *EntityMock) SetBounds() {
	e.bounds = physics.NewRectangle(
		e.Position().X,
		e.Position().Y,
		e.Position().X+e.Size().X,
		e.Position().Y+e.Size().Y,
	)
}

func (e *EntityMock) Scale() float64 {
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

func (e *EntityMock) SetComponents(c []canvas.Component) {
	e.components = c
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
