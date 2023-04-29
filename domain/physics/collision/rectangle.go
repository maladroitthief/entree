package collision

type Rectangle struct {
	MinPoint Vector
	MaxPoint Vector
	Vertices [4]Vector
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
	if x <= r.MinPoint.X || x >= r.MaxPoint.X {
		return false
	}

	if y <= r.MinPoint.Y || y >= r.MaxPoint.Y {
		return false
	}

	return true
}

func (r Rectangle) Intersects(s Rectangle) bool {
	// check if X positions are out of bounds
	if r.MinPoint.X > s.MaxPoint.X || r.MaxPoint.X < s.MinPoint.X {
		return false
	}

	// check if Y positions are out of bounds
	if r.MinPoint.Y > s.MaxPoint.Y || r.MaxPoint.Y < s.MinPoint.Y {
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
