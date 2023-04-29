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

func (v Vector) LeftNormal() Vector {
	return Vector{
		X: -v.Y,
		Y: v.X,
	}
}

func (v Vector) RightNormal() Vector {
	return Vector{
		X: v.Y,
		Y: -v.X,
	}
}

func (v Vector) Add(w Vector) Vector {
	return Vector{
		X: v.X + w.X,
		Y: v.Y + w.Y,
	}
}

func (v Vector) Subtract(w Vector) Vector {
	return Vector{
		X: v.X - w.X,
		Y: v.Y - w.Y,
	}
}

func (v Vector) Scale(c float64) Vector {
	return Vector{
		X: v.X * c,
		Y: v.Y * c,
	}
}

func (v Vector) ScaleX(c float64) Vector {
	return v.ScaleXY(c, 1)
}

func (v Vector) ScaleY(c float64) Vector {
	return v.ScaleXY(1, c)
}

func (v Vector) ScaleXY(cx, cy float64) Vector {
	return Vector{
		X: v.X * cx,
		Y: v.Y * cy,
	}
}

func (v Vector) Normalize() Vector {
	c := v.Magnitude()
	if c == 0 {
		c = 1
	}

	return v.Scale(1 / c)
}

func (v Vector) Projection(w Vector) Vector {
	return w.Scale(v.DotProduct(w) / w.DotProduct(w))
}
