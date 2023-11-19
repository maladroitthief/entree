package data

import (
	"fmt"
	"math"
)

type Polygon struct {
	Position    Vector
	Vectors     []Vector
	CalcVectors []Vector
	Planes      []Plane
}

// NewPolygon accepts an array of vectors in CCW rotation
func NewPolygon(position Vector, vectors []Vector) Polygon {
	p := Polygon{
		Position: position,
		Vectors:  vectors,
	}

	return p.Update()
}

func (p Polygon) Update() Polygon {
	p.CalcVectors = p.calcVectors()
	p.Planes = p.calcPlanes()
	return p
}

func (p Polygon) Copy(q Polygon) Polygon {
	position := q.Position.Clone()
	vectors := make([]Vector, len(q.Vectors))
	for i := 0; i < len(vectors); i++ {
		vectors[i] = q.Vectors[i].Clone()
	}

	return NewPolygon(position, vectors)
}

func (p Polygon) Clone() Polygon {
	return p.Copy(p)
}

func (p Polygon) CheckPosition(position Vector) Polygon {
	return p.Clone().SetPosition(position)
}

func (p Polygon) SetPosition(position Vector) Polygon {
	if p.Position == position {
		return p
	}

	p.Position = position

	return p.Update()
}

func (p Polygon) Info() string {
	return fmt.Sprintf("%+v", p.Vectors)
}

func (p Polygon) Add(v Vector) Polygon {
	q := p.Clone()
	q.Position = q.Position.Add(v)
	return q.Update()
}

func (p Polygon) Contains(v Vector) bool {
	for _, plane := range p.Planes {
		if plane.DistanceTo(v) > 0 {
			return false
		}
	}

	return true
}

func (p Polygon) Intersects(q Polygon) (Vector, bool) {
	normal := Vector{}
	distance := math.MaxFloat64

	for _, plane := range p.Planes {
		minP, maxP := projectVectors(plane.Normal, p.CalcVectors)
		minQ, maxQ := projectVectors(plane.Normal, q.CalcVectors)

		if minP >= maxQ || minQ >= maxP {
			return Vector{}, false
		}

		planeDistance := math.Min(maxQ-minP, maxP-minQ)
		if planeDistance < distance {
			distance = planeDistance
			normal = plane.Normal
		}
	}

	for _, plane := range q.Planes {
		minP, maxP := projectVectors(plane.Normal, p.CalcVectors)
		minQ, maxQ := projectVectors(plane.Normal, q.CalcVectors)

		if minP >= maxQ || minQ >= maxP {
			return Vector{}, false
		}

		planeDistance := math.Min(maxQ-minP, maxP-minQ)
		if planeDistance < distance {
			distance = planeDistance
			normal = plane.Normal
		}
	}

	if normal.DotProduct(q.Position.Subtract(p.Position)) < 0 {
		normal = normal.Invert()
	}

	result := normal.Scale(distance)
	return result, true
}

func projectVectors(axis Vector, vectors []Vector) (min, max float64) {
	min = math.MaxFloat64
	max = -math.MaxFloat64

	for _, v := range vectors {
		projection := v.DotProduct(axis)

		if projection < min {
			min = projection
		}
		if projection > max {
			max = projection
		}
	}

	return min, max
}

func (p Polygon) calcVectors() []Vector {
	vectors := make([]Vector, len(p.Vectors))
	for i := 0; i < len(vectors); i++ {
		vectors[i] = p.Position.Add(p.Vectors[i])
	}

	return vectors
}

func (p Polygon) calcPlanes() []Plane {
	planes := make([]Plane, len(p.CalcVectors))
	for i := 0; i < len(planes)-1; i++ {
		planes[i] = NewPlane(p.CalcVectors[i], p.CalcVectors[i+1])
	}
	planes[len(planes)-1] = NewPlane(p.CalcVectors[len(planes)-1], p.CalcVectors[0])

	return planes
}
