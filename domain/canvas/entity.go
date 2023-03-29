package canvas

import "github.com/maladroitthief/entree/domain/action"

type OrientationX int
type OrientationY int

const (
	Neutral OrientationX = iota
	West
	East
	South OrientationY = iota
	North
)

type InputComponent interface {
	Update(*Entity, []action.Input)
}

type PhysicsComponent interface {
	Update(*Entity, *Canvas)
}

type GraphicsComponent interface {
	Update(*Entity)
}
type Entity struct {
	Width             int
	Height            int
	X                 int
	Y                 int
	DeltaX            int
	DeltaY            int
	Sheet             string
	Sprite            string
	SpriteSpeed       float32
	SpriteVariant     int
	SpriteMaxVariants int
	State             string
	StateCounter      int
	OrientationX      OrientationX
	OrientationY      OrientationY
	Input             InputComponent
	Physics           PhysicsComponent
	Graphics          GraphicsComponent
}

func (e *Entity) Update(actions []action.Input, c *Canvas) {
	e.Input.Update(e, actions)
	e.Physics.Update(e, c)
	e.Graphics.Update(e)
}

func (e *Entity) VariantUpdate() {
	speed := float32(e.StateCounter) / e.SpriteSpeed
	e.SpriteVariant = int(speed)%e.SpriteMaxVariants + 1
}
