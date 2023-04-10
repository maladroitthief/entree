package sheets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/maladroitthief/entree/domain/sprite"
)

var (
	//go:embed pilot.png
	pilotSheet []byte
)

func PilotSheet() (sprite.SpriteSheet, error) {
	image, _, err := image.Decode(bytes.NewBuffer(pilotSheet))
	if err != nil {
		return nil, err
	}

	ss, err := sprite.NewSpriteSheet(
		"pilot",
		image,
		8,
		8,
		0,
		32,
	)
	if err != nil {
		return nil, err
	}

	sprites := []sprite.Sprite{}
	sprites = append(sprites, SpriteArray("idle_front", 1, 1, 6)...)
	sprites = append(sprites, SpriteArray("idle_front_side", 2, 1, 4)...)
	sprites = append(sprites, SpriteArray("idle_back", 3, 1, 6)...)
	sprites = append(sprites, SpriteArray("idle_back_side", 4, 1, 4)...)
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
