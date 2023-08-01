package ui

import "github.com/maladroitthief/entree/common/data"

type Camera struct {
	ViewPort data.Vector
	Zoom     float64
	Position data.Vector
}

func NewCamera(position data.Vector, viewPort data.Vector) *Camera {
	return &Camera{
		Position: position,
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
