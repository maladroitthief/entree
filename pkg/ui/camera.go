package ui

import (
	"github.com/maladroitthief/mosaic"
)

type Camera struct {
	ViewPort mosaic.Vector
	Zoom     float64
	X        float64
	Y        float64
}

func NewCamera(x, y float64, viewPort mosaic.Vector) *Camera {
	return &Camera{
		X:        x,
		Y:        y,
		ViewPort: viewPort,
		Zoom:     5,
	}
}

func (c *Camera) ViewPortCenter() mosaic.Vector {
	return mosaic.Vector{
		X: c.ViewPort.X / 2,
		Y: c.ViewPort.Y / 2,
	}
}
