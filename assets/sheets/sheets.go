package sheets

import (
	"fmt"

	"github.com/maladroitthief/entree/domain/sprite"
)

func SpriteArray(name string, row, columnStart, columnEnd int) []sprite.Sprite {
	if columnStart > columnEnd {
		return []sprite.Sprite{
			{
				Name:   name,
				Row:    row,
				Column: columnStart,
			},
		}
	}

	spriteCount := columnEnd - columnStart + 1
	results := make([]sprite.Sprite, spriteCount)

	for i := 0; i < spriteCount; i++ {
		results[i] = sprite.Sprite{
			Name:   fmt.Sprintf("%s_%d", name, columnStart+i),
			Row:    row,
			Column: columnStart + i,
		}
	}

	return results
}
