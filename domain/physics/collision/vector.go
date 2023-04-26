package collision

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v Vector) DotProduct(w Vector) float64 {
	return (v.X * w.X) + (v.Y * w.Y)
}

func (v Vector) Scale(c float64) Vector {
	return Vector{
		X: v.X * c,
		Y: v.Y * c,
	}
}

func (v Vector) Projection(w Vector) Vector {
	return w.Scale(v.DotProduct(w) / w.DotProduct(w))
}
