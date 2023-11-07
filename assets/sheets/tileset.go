package sheets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/maladroitthief/entree/pkg/content"
)

var (
	//go:embed tileset.png
	tilesSheet []byte
)

func TilesSheet() (*content.SpriteSheet, error) {
	image, _, err := image.Decode(bytes.NewBuffer(tilesSheet))
	if err != nil {
		return nil, err
	}

	ss, err := content.NewSpriteSheet(
		"tiles",
		image,
		32,
		32,
		0,
		32,
	)
	if err != nil {
		return nil, err
	}

	sprites := []content.Sprite{}

	sprites = append(sprites, Sprite("rock_1", 1, 1))
	sprites = append(sprites, Sprite("grass_1", 1, 2))
	sprites = append(sprites, Sprite("grass_2", 1, 3))
	sprites = append(sprites, Sprite("grass_3", 1, 4))
	sprites = append(sprites, Sprite("grass_4", 1, 5))

	for _, s := range sprites {
		err = ss.AddSprite(s)
		if err != nil {
			return nil, err
		}
	}

	return ss, nil
}
