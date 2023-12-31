package sheets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/maladroitthief/entree/pkg/content"
)

var (
	//go:embed federico.png
	federicoSheet []byte
)

func FedericoSheet() (*content.SpriteSheet, error) {
	image, _, err := image.Decode(bytes.NewBuffer(federicoSheet))
	if err != nil {
		return nil, err
	}

	ss, err := content.NewSpriteSheet(
		"federico",
		image,
		8,
		6,
		0,
		32,
	)
	if err != nil {
		return nil, err
	}

	sprites := []content.Sprite{}
	sprites = append(sprites, SpriteArray("idle_front", 1, 1, 6)...)
	sprites = append(sprites, SpriteArray("idle_front_side", 2, 1, 6)...)
	sprites = append(sprites, SpriteArray("idle_back", 3, 1, 6)...)
	sprites = append(sprites, SpriteArray("idle_back_side", 4, 1, 6)...)
	sprites = append(sprites, SpriteArray("move_front", 5, 1, 6)...)
	sprites = append(sprites, SpriteArray("move_front_side", 6, 1, 6)...)
	sprites = append(sprites, SpriteArray("move_back", 7, 1, 6)...)
	sprites = append(sprites, SpriteArray("move_back_side", 8, 1, 6)...)

	for _, s := range sprites {
		err = ss.AddSprite(s)
		if err != nil {
			return nil, err
		}
	}

	return ss, nil
}
