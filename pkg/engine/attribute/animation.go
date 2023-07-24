package attribute

import "github.com/maladroitthief/entree/common/data"

const (
	DefaultSize  = 16
	DefaultSpeed = 60
)

type Animation struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Static      bool
	Speed       float64
	Variant     int
	VariantMax  int
	SpriteSheet string
	Sprite      string
}

func NewAnimation(sheet, sprite string) Animation {
	return Animation{
		Speed:       DefaultSpeed,
		Variant:     1,
		VariantMax:  1,
		SpriteSheet: sheet,
		Sprite:      sprite,
	}
}
