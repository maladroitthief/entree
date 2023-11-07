package ui

import "github.com/maladroitthief/entree/common/data"

type Camera struct {
	ViewPort data.Vector
	Zoom     float64
	X        float64
	Y        float64
}

func NewCamera(x, y float64, viewPort data.Vector) *Camera {
	return &Camera{
		X:        x,
		Y:        y,
		ViewPort: viewPort,
		Zoom:     5,
	}
}

func (c *Camera) ViewPortCenter() data.Vector {
	return data.Vector{
		X: c.ViewPort.X / 2,
		Y: c.ViewPort.Y / 2,
	}
}
