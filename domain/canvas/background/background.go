package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics/collision"
)

type backgroundEntity struct {
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

func (e *backgroundEntity) Update(c *canvas.Canvas) {
	e.SetBounds()
	e.InputComponent().Update(e)
	e.PhysicsComponent().Update(e, c)
	e.GraphicsComponent().Update(e)
}

func (e *backgroundEntity) Send(msg, val string) {
	for _, c := range e.Components() {
		c.Receive(e, msg, val)
	}
}

func (e *backgroundEntity) Position() collision.Vector {
	return e.position.Scale(e.Scale())
}

func (e *backgroundEntity) SetPosition(v collision.Vector) {
	e.position = v.Scale(1 / e.Scale())
}

func (e *backgroundEntity) Size() collision.Vector {
	return e.size.Scale(e.Scale())
}

func (e *backgroundEntity) SetSize(v collision.Vector) {
	e.size = v.Scale(1 / e.Scale())
}

func (e *backgroundEntity) Bounds() collision.Rectangle {
	return e.bounds
}

func (e *backgroundEntity) SetBounds() {
	e.bounds = collision.NewRectangle(
		e.Position().X,
		e.Position().Y,
		e.Position().X+e.Size().X,
		e.Position().Y+e.Size().Y,
	)
}

func (e *backgroundEntity) Scale() float64 {
	if e.scale <= 0 {
		return 1
	}

	return e.scale
}

func (e *backgroundEntity) SetScale(f float64) {
	e.scale = f
}

func (e *backgroundEntity) Sheet() string {
	return e.sheet
}

func (e *backgroundEntity) SetSheet(s string) {
	e.sheet = s
}

func (e *backgroundEntity) Sprite() string {
	return e.sprite
}

func (e *backgroundEntity) SetSprite(s string) {
	e.sprite = s
}

func (e *backgroundEntity) State() string {
	return e.state
}

func (e *backgroundEntity) SetState(s string) {
	e.state = s
}

func (e *backgroundEntity) StateCounter() int {
	return e.stateCounter
}

func (e *backgroundEntity) SetStateCounter(i int) {
	e.stateCounter = i
}

func (e *backgroundEntity) IncrementStateCounter() {
	e.stateCounter++
}

func (e *backgroundEntity) OrientationX() canvas.OrientationX {
	return e.orientationX
}

func (e *backgroundEntity) SetOrientationX(o canvas.OrientationX) {
	e.orientationX = o
}

func (e *backgroundEntity) OrientationY() canvas.OrientationY {
	return e.orientationY
}

func (e *backgroundEntity) SetOrientationY(o canvas.OrientationY) {
	e.orientationY = o
}

func (e *backgroundEntity) Components() []canvas.Component {
	return e.components
}

func (e *backgroundEntity) SetComponents(c []canvas.Component) {
	e.components = c
}

func (e *backgroundEntity) InputComponent() canvas.InputComponent {
	return e.input
}

func (e *backgroundEntity) SetInputComponent(c canvas.InputComponent) {
	e.input = c
}

func (e *backgroundEntity) PhysicsComponent() canvas.PhysicsComponent {
	return e.physics
}

func (e *backgroundEntity) SetPhysicsComponent(c canvas.PhysicsComponent) {
	e.physics = c
}

func (e *backgroundEntity) GraphicsComponent() canvas.GraphicsComponent {
	return e.graphics
}

func (e *backgroundEntity) SetGraphicsComponent(c canvas.GraphicsComponent) {
	e.graphics = c
}
