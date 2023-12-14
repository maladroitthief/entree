package attribute

import (
	"fmt"

	"github.com/maladroitthief/entree/common/data"
)

const (
	DefaultSize  = 32
	DefaultSpeed = 60
)

type Animation struct {
	Id       data.GenerationalIndex
	EntityId data.GenerationalIndex

	Counter     int
	Static      bool
	Speed       float64
	Variant     int
	VariantMax  int
	SpriteSheet string
	Sprite      string
	Sprites     map[string][]string
}

func NewAnimation(sheet, defaultSprite string) Animation {
	return Animation{
		Speed:       DefaultSpeed,
		Counter:     0,
		Variant:     1,
		VariantMax:  1,
		SpriteSheet: sheet,
		Sprite:      defaultSprite,
		Sprites:     map[string][]string{},
	}
}

func SpriteArray(spriteName string, count int) []string {
	results := make([]string, count)

	for i := 0; i < count; i++ {
		results[i] = fmt.Sprintf("%s_%d", spriteName, i+1)
	}

	return results
}
