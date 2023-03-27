package assets

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
		1,
		3,
		32,
	)
	if err != nil {
		return nil, err
	}

	sprites := []sprite.Sprite{
		{
			Name:   "idle_down",
			Row:    1,
			Column: 1,
		},
		{
			Name:   "idle_left",
			Row:    1,
			Column: 2,
		},
		{
			Name:   "idle_up",
			Row:    1,
			Column: 3,
		},
	}

	for _, s := range sprites {
		err = ss.AddSprite(s)
		if err != nil {
			return nil, err
		}
	}

	return ss, nil
}
