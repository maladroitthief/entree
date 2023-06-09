package scene

import (
	"github.com/maladroitthief/entree/domain/physics"
)

type Camera interface {
	Update()
	ViewPort() physics.Vector
	SetViewPort(physics.Vector)
	Position() physics.Vector
	SetFocalPoint(FocalPoint)
	Zoom() float64
	SetZoom(float64)
	ViewportCenter() physics.Vector
}

type camera struct {
	viewPort   physics.Vector
	zoom       float64
	focalPoint FocalPoint
}

func NewCamera(fp FocalPoint, viewPort physics.Vector) Camera {
	return &camera{
		focalPoint: fp,
		viewPort:   viewPort,
		zoom:       1,
	}
}

func (c *camera) Update() {

}

func (c *camera) ViewPort() physics.Vector {
	return c.viewPort
}

func (c *camera) SetViewPort(vp physics.Vector) {
	c.viewPort = vp
}

func (c *camera) Position() physics.Vector {
	return c.focalPoint.Position()
}

func (c *camera) SetFocalPoint(fp FocalPoint) {
	c.focalPoint = fp
}

func (c *camera) Zoom() float64 {
	return c.zoom
}

func (c *camera) SetZoom(z float64) {
	c.zoom = z
}

func (c *camera) ViewportCenter() physics.Vector {
	return physics.Vector{
		X: c.viewPort.X / 2,
		Y: c.viewPort.Y / 2,
	}
}
