package background

import (
	"github.com/maladroitthief/entree/domain/canvas"
)

func NewGrass(x, y int) *canvas.Entity {
	return &canvas.Entity{
		Width:             16,
		Height:            16,
		X:                 x,
		Y:                 y,
		Sheet:             "test",
		Sprite:            "grass",
		SpriteSpeed:       40,
		SpriteVariant:     1,
		SpriteMaxVariants: 1,
		OrientationX:      canvas.Neutral,
		OrientationY:      canvas.South,
		Input:             &backgroundInput{},
		Physics:           &backgroundPhysics{},
		Graphics:          &backgroundGraphics{},
	}
}
