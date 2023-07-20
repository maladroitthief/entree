package physics

type Rectangle struct {
	MinPoint Vector
	MaxPoint Vector
	Vertices [4]Vector
}

func Bounds(position, size Vector) Rectangle {
	return NewRectangle(
		position.X-size.X/2,
		position.Y-size.Y/2,
		position.X+size.X/2,
		position.Y+size.Y/2,
	)
}

// NewRectangle creates a rectangle that has a counter-clockwise rotation of vectors
func NewRectangle(x1, y1, x2, y2 float64) Rectangle {
	minPoint := Vector{x1, y1}
	maxPoint := Vector{x2, y2}

	if x1 > x2 {
		minPoint.X = x2
		maxPoint.X = x1
	}

	if y1 > y2 {
		minPoint.Y = y2
		maxPoint.Y = y1
	}

	return Rectangle{
		MinPoint: minPoint,
		MaxPoint: maxPoint,
		Vertices: [4]Vector{
			{X: minPoint.X, Y: maxPoint.Y},
			{X: minPoint.X, Y: minPoint.Y},
			{X: maxPoint.X, Y: minPoint.Y},
			{X: maxPoint.X, Y: maxPoint.Y},
		},
	}
}

func (r Rectangle) Width() float64 {
	return r.MaxPoint.X - r.MinPoint.X
}

func (r Rectangle) Height() float64 {
	return r.MaxPoint.Y - r.MinPoint.Y
}

func (r Rectangle) Contains(x, y float64) bool {
	if x < r.MinPoint.X || x > r.MaxPoint.X {
		return false
	}

	if y < r.MinPoint.Y || y > r.MaxPoint.Y {
		return false
	}

	return true
}

func (r Rectangle) Intersects(s Rectangle) bool {
	d1x := s.MinPoint.X - r.MaxPoint.X
	d1y := s.MinPoint.Y - r.MaxPoint.Y
	d2x := r.MinPoint.X - s.MaxPoint.X
	d2y := r.MinPoint.Y - s.MaxPoint.Y

	if d1x > 0.0 || d1y > 0.0 {
		return false
	}

	if d2x > 0.0 || d2y > 0.0 {
		return false
	}

	return true
}

func (r Rectangle) Center() Vector {
	return Vector{
		X: r.MinPoint.X + (r.Width() / 2),
		Y: r.MinPoint.Y + (r.Height() / 2),
	}
}

func (r Rectangle) Vertex1() Vector {
	return Vector{
		X: r.MinPoint.X,
		Y: r.MaxPoint.Y,
	}
}

func (r Rectangle) Vertex2() Vector {
	return r.MaxPoint
}

func (r Rectangle) Vertex3() Vector {
	return Vector{
		X: r.MaxPoint.X,
		Y: r.MinPoint.Y,
	}
}

func (r Rectangle) Vertex4() Vector {
	return r.MinPoint
}

func (r Rectangle) Scale(c float64) Rectangle {
	scaledWidth := r.Width() * c
	scaledHeight := r.Height() * c

	return NewRectangle(
		r.MinPoint.X-(scaledWidth/2),
		r.MinPoint.Y-(scaledHeight/2),
		r.MaxPoint.X+(scaledWidth/2),
		r.MaxPoint.Y+(scaledHeight/2),
	)
}
