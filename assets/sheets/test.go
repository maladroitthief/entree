package sheets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/maladroitthief/entree/pkg/content"
)

var (
	//go:embed test.png
	testSheet []byte
)

func TestSheet() (*content.SpriteSheet, error) {
	image, _, err := image.Decode(bytes.NewBuffer(testSheet))
	if err != nil {
		return nil, err
	}

	ss, err := content.NewSpriteSheet(
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

	sprites := []content.Sprite{}
	sprites = append(sprites, Sprite("blank", 1, 1))
	sprites = append(sprites, Sprite("small_gravel", 1, 2))
	sprites = append(sprites, Sprite("big_gravel", 1, 3))
	sprites = append(sprites, Sprite("small_stones", 1, 4))
	sprites = append(sprites, Sprite("big_stones", 1, 5))
	sprites = append(sprites, Sprite("grass", 1, 6))
	sprites = append(sprites, Sprite("flowers", 1, 7))
	sprites = append(sprites, Sprite("tall_grass", 1, 8))
	sprites = append(sprites, Sprite("weeds", 3, 1))

	sprites = append(sprites, Sprite("wall_group_north_west", 17, 1))
	sprites = append(sprites, Sprite("wall_group_north", 17, 2))
	sprites = append(sprites, Sprite("wall_group_north_east", 17, 3))
	sprites = append(sprites, Sprite("wall_group_west", 18, 1))
	sprites = append(sprites, Sprite("wall_group", 18, 2))
	sprites = append(sprites, Sprite("wall_group_east", 18, 3))
	sprites = append(sprites, Sprite("wall_group_south_west", 19, 1))
	sprites = append(sprites, Sprite("wall_group_south", 19, 2))
	sprites = append(sprites, Sprite("wall_group_south_east", 19, 3))
	sprites = append(sprites, Sprite("wall_group_south_east", 19, 3))
	sprites = append(sprites, Sprite("wall_north_west", 17, 4))
	sprites = append(sprites, Sprite("wall_north_east", 17, 5))
	sprites = append(sprites, Sprite("wall_south_west", 18, 4))
	sprites = append(sprites, Sprite("wall_south_east", 18, 5))
	sprites = append(sprites, Sprite("wall", 19, 4))
	sprites = append(sprites, Sprite("wall_tall", 19, 5))

	for _, s := range sprites {
		err = ss.AddSprite(s)
		if err != nil {
			return nil, err
		}
	}

	return ss, nil
}
