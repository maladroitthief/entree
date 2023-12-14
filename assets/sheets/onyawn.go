package sheets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/maladroitthief/entree/pkg/content"
)

var (
	//go:embed onyawn.png
	onyawnSheet []byte
)

func OnyawnSheet() (*content.SpriteSheet, error) {
	image, _, err := image.Decode(bytes.NewBuffer(onyawnSheet))
	if err != nil {
		return nil, err
	}

	ss, err := content.NewSpriteSheet(
		"onyawn",
		image,
		6,
		6,
		0,
		32,
	)
	if err != nil {
		return nil, err
	}

	sprites := []content.Sprite{}
	sprites = append(sprites, SpriteArray("idle_front", 1, 1, 2)...)
	sprites = append(sprites, SpriteArray("idle_front_side", 2, 1, 2)...)
	sprites = append(sprites, SpriteArray("idle_back", 3, 1, 2)...)
	sprites = append(sprites, SpriteArray("move_front", 4, 1, 6)...)
	sprites = append(sprites, SpriteArray("move_front_side", 5, 1, 6)...)
	sprites = append(sprites, SpriteArray("move_back", 6, 1, 6)...)

	for _, s := range sprites {
		err = ss.AddSprite(s)
		if err != nil {
			return nil, err
		}
	}

	return ss, nil
}
