package environment

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics/collision"
)

type environmentEntity struct {
	position     collision.Vector
	size         collision.Vector
	bounds       collision.Rectangle
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

func (e *environmentEntity) Update(c *canvas.Canvas) {
	e.SetBounds()
	e.InputComponent().Update(e)
	e.PhysicsComponent().Update(e, c)
	e.GraphicsComponent().Update(e)
}

func (e *environmentEntity) Send(msg, val string) {
	for _, c := range e.Components() {
		c.Receive(e, msg, val)
	}
}

func (e *environmentEntity) Position() collision.Vector {
	return e.position.Scale(e.Scale())
}

func (e *environmentEntity) SetPosition(v collision.Vector) {
	e.position = v.Scale(1 / e.Scale())
}

func (e *environmentEntity) Size() collision.Vector {
	return e.size.Scale(e.Scale())
}

func (e *environmentEntity) SetSize(v collision.Vector) {
	e.size = v.Scale(1 / e.Scale())
}

func (e *environmentEntity) Bounds() collision.Rectangle {
	return e.bounds
}

func (e *environmentEntity) SetBounds() {
	e.bounds = collision.NewRectangle(
		e.Position().X,
		e.Position().Y,
		e.Position().X+e.Size().X,
		e.Position().Y+e.Size().Y,
	)
}

func (e *environmentEntity) Scale() float64 {
	if e.scale <= 0 {
		return 1
	}

	return e.scale
}

func (e *environmentEntity) SetScale(f float64) {
	e.scale = f
}

func (e *environmentEntity) Sheet() string {
	return e.sheet
}

func (e *environmentEntity) SetSheet(s string) {
	e.sheet = s
}

func (e *environmentEntity) Sprite() string {
	return e.sprite
}

func (e *environmentEntity) SetSprite(s string) {
	e.sprite = s
}

func (e *environmentEntity) State() string {
	return e.state
}

func (e *environmentEntity) SetState(s string) {
	e.state = s
}

func (e *environmentEntity) StateCounter() int {
	return e.stateCounter
}

func (e *environmentEntity) SetStateCounter(i int) {
	e.stateCounter = i
}

func (e *environmentEntity) IncrementStateCounter() {
	e.stateCounter++
}

func (e *environmentEntity) OrientationX() canvas.OrientationX {
	return e.orientationX
}

func (e *environmentEntity) SetOrientationX(o canvas.OrientationX) {
	e.orientationX = o
}

func (e *environmentEntity) OrientationY() canvas.OrientationY {
	return e.orientationY
}

func (e *environmentEntity) SetOrientationY(o canvas.OrientationY) {
	e.orientationY = o
}

func (e *environmentEntity) Components() []canvas.Component {
	return e.components
}

func (e *environmentEntity) SetComponents(c []canvas.Component) {
	e.components = c
}

func (e *environmentEntity) InputComponent() canvas.InputComponent {
	return e.input
}

func (e *environmentEntity) SetInputComponent(c canvas.InputComponent) {
	e.input = c
}

func (e *environmentEntity) PhysicsComponent() canvas.PhysicsComponent {
	return e.physics
}

func (e *environmentEntity) SetPhysicsComponent(c canvas.PhysicsComponent) {
	e.physics = c
}

func (e *environmentEntity) GraphicsComponent() canvas.GraphicsComponent {
	return e.graphics
}

func (e *environmentEntity) SetGraphicsComponent(c canvas.GraphicsComponent) {
	e.graphics = c
}
