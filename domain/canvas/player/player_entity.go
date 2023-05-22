package player

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics/collision"
)

type playerEntity struct {
	position     collision.Vector
	size         collision.Vector
	offset       collision.Vector
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

func (e *playerEntity) Update(c *canvas.Canvas) {
	e.SetBounds()
	e.InputComponent().Update(e)
	e.PhysicsComponent().Update(e, c)
	e.GraphicsComponent().Update(e)
}

func (e *playerEntity) Send(msg, val string) {
	for _, c := range e.Components() {
		c.Receive(e, msg, val)
	}
}

func (e *playerEntity) Position() collision.Vector {
	return e.position.Scale(e.Scale())
}

func (e *playerEntity) SetPosition(v collision.Vector) {
	e.position = v.Scale(1 / e.Scale())
}

func (e *playerEntity) Size() collision.Vector {
	return e.size.Scale(e.Scale())
}

func (e *playerEntity) SetSize(v collision.Vector) {
	e.size = v.Scale(1 / e.Scale())
}

func (e *playerEntity) Offset() collision.Vector {
	return e.offset
}

func (e *playerEntity) Bounds() collision.Rectangle {
	return e.bounds
}

func (e *playerEntity) SetBounds() {
	e.bounds = collision.Bounds(e.Position(), e.Size())
}

func (e *playerEntity) Scale() float64 {
	if e.scale <= 0 {
		return 1
	}

	return e.scale
}

func (e *playerEntity) SetScale(f float64) {
	e.scale = f
}

func (e *playerEntity) Sheet() string {
	return e.sheet
}

func (e *playerEntity) SetSheet(s string) {
	e.sheet = s
}

func (e *playerEntity) Sprite() string {
	return e.sprite
}

func (e *playerEntity) SetSprite(s string) {
	e.sprite = s
}

func (e *playerEntity) State() string {
	return e.state
}

func (e *playerEntity) SetState(s string) {
	e.state = s
}

func (e *playerEntity) StateCounter() int {
	return e.stateCounter
}

func (e *playerEntity) SetStateCounter(i int) {
	e.stateCounter = i
}

func (e *playerEntity) IncrementStateCounter() {
	e.stateCounter++
}

func (e *playerEntity) OrientationX() canvas.OrientationX {
	return e.orientationX
}

func (e *playerEntity) SetOrientationX(o canvas.OrientationX) {
	e.orientationX = o
}

func (e *playerEntity) OrientationY() canvas.OrientationY {
	return e.orientationY
}

func (e *playerEntity) SetOrientationY(o canvas.OrientationY) {
	e.orientationY = o
}

func (e *playerEntity) Components() []canvas.Component {
	return e.components
}

func (e *playerEntity) SetComponents(c []canvas.Component) {
	e.components = c
}

func (e *playerEntity) InputComponent() canvas.InputComponent {
	return e.input
}

func (e *playerEntity) SetInputComponent(c canvas.InputComponent) {
	e.input = c
}

func (e *playerEntity) PhysicsComponent() canvas.PhysicsComponent {
	return e.physics
}

func (e *playerEntity) SetPhysicsComponent(c canvas.PhysicsComponent) {
	e.physics = c
}

func (e *playerEntity) GraphicsComponent() canvas.GraphicsComponent {
	return e.graphics
}

func (e *playerEntity) SetGraphicsComponent(c canvas.GraphicsComponent) {
	e.graphics = c
}
