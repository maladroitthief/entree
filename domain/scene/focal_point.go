package scene

import "github.com/maladroitthief/entree/domain/physics"

type FocalPoint interface {
	Position() physics.Vector
}

type focalPoint struct {
	position physics.Vector
}

func (fp *focalPoint) Position() physics.Vector {
	return fp.position
}
