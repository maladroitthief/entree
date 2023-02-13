package character

type Eugene struct {
	x            float64
	y            float64
	mass         int
	acceleration float64
	maxVelocity  float64
}

func (e *Eugene) Update() {

}

func(e *Eugene) GetX() float64{
  return e.x
}

func(e *Eugene) SetX(x float64){
  e.x = x
}

func(e *Eugene) GetY() float64{
  return e.y
}

func(e *Eugene) SetY(y float64){
  e.y = y
}

func(e *Eugene) GetMass() int{
  return e.mass
}

func(e *Eugene) SetMass(m int){
  e.mass = m
}

func(e *Eugene) GetAcceleration() float64{
  return e.acceleration
}

func(e *Eugene) SetAcceleration(a float64){
  e.acceleration = a
}

func(e *Eugene) GetMaxVelocity() float64{
  return e.maxVelocity
}

func(e *Eugene) SetMaxVelocity(v float64){
  e.maxVelocity = v
}

