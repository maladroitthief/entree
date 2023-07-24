package sheets

import (
	"fmt"

	"github.com/maladroitthief/entree/pkg/content"
)

func SpriteArray(name string, row, columnStart, columnEnd int) []content.Sprite {
	if columnStart > columnEnd {
		return []content.Sprite{
			{
				Name:   name,
				Row:    row,
				Column: columnStart,
			},
		}
	}

	spriteCount := columnEnd - columnStart + 1
	results := make([]content.Sprite, spriteCount)

	for i := 0; i < spriteCount; i++ {
		results[i] = content.Sprite{
			Name:   fmt.Sprintf("%s_%d", name, columnStart+i),
			Row:    row,
			Column: columnStart + i,
		}
	}

	return results
}

func Sprite(name string, row, col int) content.Sprite {
  return content.Sprite{
      Name:   name,
      Row:    row,
      Column: col,
  }
}
