package collision

type Rectangle struct {
	minPoint Point
	maxPoint Point
}

func NewRectangle(x1, y1, x2, y2 float64) Rectangle {
	minPoint := Point{x1, y1}
	maxPoint := Point{x2, y2}

	if x1 > x2 {
		minPoint.X = x2
		maxPoint.X = x1
	}

	if y1 > y2 {
		minPoint.Y = y2
		maxPoint.Y = y1
	}
	return Rectangle{
		minPoint: minPoint,
		maxPoint: maxPoint,
	}
}

func (r Rectangle) Width() float64 {
	return r.maxPoint.X - r.minPoint.X
}

func (r Rectangle) Height() float64 {
	return r.maxPoint.Y - r.minPoint.Y
}

func (r Rectangle) Contains(x, y float64) bool {
	if x <= r.minPoint.X || x >= r.maxPoint.X {
		return false
	}

	if y <= r.minPoint.Y || y >= r.maxPoint.Y {
		return false
	}

	return true
}

func (r Rectangle) Intersects(s Rectangle) bool {
	// check if X positions are out of bounds
	if r.minPoint.X > s.maxPoint.X || r.maxPoint.X < s.minPoint.X {
		return false
	}

	// check if Y positions are out of bounds
	if r.minPoint.Y > s.maxPoint.Y || r.maxPoint.Y < s.minPoint.Y {
		return false
	}

	return true
}
