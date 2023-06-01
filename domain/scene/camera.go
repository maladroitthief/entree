package scene

import "github.com/maladroitthief/entree/domain/physics"

type Camera struct {
	ViewPort physics.Vector
	Position physics.Vector
	Zoom     float64
}
