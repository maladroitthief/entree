package canvas

type OrientationX int
type OrientationY int

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
	DeltaX            float64
	DeltaY            float64
	Acceleration      float64
	Deceleration      float64
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
	speed := float64(e.StateCounter) / (e.SpriteSpeed / float64(e.SpriteMaxVariants))
	e.SpriteVariant = int(speed)%e.SpriteMaxVariants + 1
}

func (e *Entity) Reset() {
	e.State = "idle"
	if e.DeltaX == 0 && e.DeltaY != 0 {
		e.OrientationX = Neutral
	}
	e.DeltaX = 0
	e.DeltaY = 0
}
