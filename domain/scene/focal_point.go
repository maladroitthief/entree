package scene

type FocalPoint interface {
	Position() (x, y float64)
}

type focalPoint struct {
	x float64
	y float64
}

func (fp *focalPoint) Position() (x, y float64) {
	return fp.x, fp.y
}
