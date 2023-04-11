package sheets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/maladroitthief/entree/domain/sprite"
)

var (
	//go:embed test.png
	testSheet []byte
)

func TestSheet() (sprite.SpriteSheet, error) {
	image, _, err := image.Decode(bytes.NewBuffer(testSheet))
	if err != nil {
		return nil, err
	}

	ss, err := sprite.NewSpriteSheet(
		"test",
		image,
		22,
		49,
		1,
		16,
	)
	if err != nil {
		return nil, err
	}

	sprites := []sprite.Sprite{}
	sprites = append(sprites, SpriteArray("blank", 1, 1, 0)...)
	sprites = append(sprites, SpriteArray("small_gravel", 1, 2, 0)...)
	sprites = append(sprites, SpriteArray("big_gravel", 1, 3, 0)...)
	sprites = append(sprites, SpriteArray("small_stones", 1, 4, 0)...)
	sprites = append(sprites, SpriteArray("big_stones", 1, 5, 0)...)
	sprites = append(sprites, SpriteArray("grass", 1, 6, 0)...)
	sprites = append(sprites, SpriteArray("flowers", 1, 7, 0)...)
	sprites = append(sprites, SpriteArray("tall_grass", 1, 8, 0)...)

	for _, s := range sprites {
		err = ss.AddSprite(s)
		if err != nil {
			return nil, err
		}
	}

	return ss, nil
}
