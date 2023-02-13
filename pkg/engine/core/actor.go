package core

type Actor interface {
	Update()

	New(InputHandler) Actor

	GetX() float64
	SetX(float64)

	GetY() float64
	SetY(float64)

	GetMass() int
	SetMass(int)

	GetAcceleration() float64
	SetAcceleration(float64)

	GetMaxVelocity() float64
	SetMaxVelocity(float64)
}

type InputHandler interface {
	Update(Actor)
}
