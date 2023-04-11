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
	sprites = append(sprites, Sprite("blank", 1, 1))
	sprites = append(sprites, Sprite("small_gravel", 1, 2))
	sprites = append(sprites, Sprite("big_gravel", 1, 3))
	sprites = append(sprites, Sprite("small_stones", 1, 4))
	sprites = append(sprites, Sprite("big_stones", 1, 5))
	sprites = append(sprites, Sprite("grass", 1, 6))
	sprites = append(sprites, Sprite("flowers", 1, 7))
	sprites = append(sprites, Sprite("tall_grass", 1, 8))
	sprites = append(sprites, Sprite("tall_grass", 1, 8))

	for _, s := range sprites {
		err = ss.AddSprite(s)
		if err != nil {
			return nil, err
		}
	}

	return ss, nil
}
