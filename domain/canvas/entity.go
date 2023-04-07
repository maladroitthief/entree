package canvas

type OrientationX int
type OrientationY int

const (
	Neutral OrientationX = iota
	West
	East
	South OrientationY = iota
	North
	DefaultAcceleration = 1
	DefaultDeceleration = 1
	DefaultMaxVelocity  = 10
	DefaultMass         = 5
)

type InputComponent interface {
	Update(*Entity)
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
	Acceleration      int
	Deceleration      int
	MaxVelocity       int
	Mass              int
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

func (e *Entity) Update(c *Canvas) {
	e.Input.Update(e)
	e.Physics.Update(e, c)
	e.Graphics.Update(e)
}

func (e *Entity) VariantUpdate() {
	speed := float32(e.StateCounter) / e.SpriteSpeed
	e.SpriteVariant = int(speed)%e.SpriteMaxVariants + 1
}

func (e *Entity) Reset() {
  e.State = "idle"
  e.OrientationX = Neutral
  e.OrientationY = South
  e.DeltaX = 0
  e.DeltaY = 0
}
