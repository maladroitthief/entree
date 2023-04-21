package canvas

import "github.com/maladroitthief/entree/domain/physics/collision"

type OrientationX int
type OrientationY int
type Size collision.Point
type Position collision.Point
type DeltaPosition collision.Point
type Velocity collision.Point

const (
	Neutral OrientationX = iota
	West
	East
	South OrientationY = iota
	North
	DefaultAcceleration = 3
	DefaultDeceleration = 15
	DefaultMaxVelocity  = 10
	DefaultMass         = 1
)

type Entity struct {
	Size              Size
	Position          Position
	DeltaPosition     DeltaPosition
	Velocity          Velocity
	Acceleration      float64
	VelocityX         float64
	VelocityY         float64
	MaxVelocity       float64
	Mass              float64
	Sheet             string
	Sprite            string
	SpriteSpeed       float64
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

func (e *Entity) Update(c *Canvas) {
	e.Input.Update(e)
	e.Physics.Update(e, c)
	e.Graphics.Update(e)
}

func (e *Entity) VariantUpdate() {
	speed := float64(e.StateCounter) /
		(e.SpriteSpeed / float64(e.SpriteMaxVariants))
	e.SpriteVariant = int(speed)%e.SpriteMaxVariants + 1
}

func (e *Entity) Reset() {
	e.State = "idle"
	dx, dy := e.DeltaPosition.X, e.DeltaPosition.Y
	if dx == 0 && dy != 0 {
		e.OrientationX = Neutral
	}
	e.DeltaPosition = DeltaPosition{0, 0}
}
