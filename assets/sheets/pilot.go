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
		32,
	)
	if err != nil {
		return nil, err
	}

	sprites := []sprite.Sprite{
		{
			Name:   "idle_front_1",
			Row:    1,
			Column: 1,
		},
		{
			Name:   "idle_front_2",
			Row:    1,
			Column: 2,
		},
		{
			Name:   "idle_front_3",
			Row:    1,
			Column: 3,
		},
		{
			Name:   "idle_front_4",
			Row:    1,
			Column: 4,
		},
		{
			Name:   "idle_front_5",
			Row:    1,
			Column: 5,
		},
		{
			Name:   "idle_front_6",
			Row:    1,
			Column: 6,
		},
		{
			Name:   "idle_front_side_1",
			Row:    2,
			Column: 1,
		},
		{
			Name:   "idle_front_side_2",
			Row:    2,
			Column: 2,
		},
		{
			Name:   "idle_front_side_3",
			Row:    2,
			Column: 3,
		},
		{
			Name:   "idle_front_side_4",
			Row:    2,
			Column: 4,
		},
		{
			Name:   "idle_back_1",
			Row:    3,
			Column: 1,
		},
		{
			Name:   "idle_back_2",
			Row:    3,
			Column: 2,
		},
		{
			Name:   "idle_back_3",
			Row:    3,
			Column: 3,
		},
		{
			Name:   "idle_back_4",
			Row:    3,
			Column: 4,
		},
		{
			Name:   "idle_back_5",
			Row:    3,
			Column: 5,
		},
		{
			Name:   "idle_back_6",
			Row:    3,
			Column: 6,
		},
		{
			Name:   "idle_back_side_1",
			Row:    4,
			Column: 1,
		},
		{
			Name:   "idle_back_side_2",
			Row:    4,
			Column: 2,
		},
		{
			Name:   "idle_back_side_3",
			Row:    4,
			Column: 3,
		},
		{
			Name:   "idle_back_side_4",
			Row:    4,
			Column: 4,
		},
		{
			Name:   "move_front_1",
			Row:    5,
			Column: 1,
		},
		{
			Name:   "move_front_2",
			Row:    5,
			Column: 2,
		},
		{
			Name:   "move_front_3",
			Row:    5,
			Column: 3,
		},
		{
			Name:   "move_front_4",
			Row:    5,
			Column: 4,
		},
		{
			Name:   "move_front_5",
			Row:    5,
			Column: 5,
		},
		{
			Name:   "move_front_6",
			Row:    5,
			Column: 6,
		},
		{
			Name:   "move_front_side_1",
			Row:    6,
			Column: 1,
		},
		{
			Name:   "move_front_side_2",
			Row:    6,
			Column: 2,
		},
		{
			Name:   "move_front_side_3",
			Row:    6,
			Column: 3,
		},
		{
			Name:   "move_front_side_4",
			Row:    6,
			Column: 4,
		},
		{
			Name:   "move_front_side_5",
			Row:    6,
			Column: 5,
		},
		{
			Name:   "move_front_side_6",
			Row:    6,
			Column: 6,
		},
		{
			Name:   "move_back_1",
			Row:    7,
			Column: 1,
		},
		{
			Name:   "move_back_2",
			Row:    7,
			Column: 2,
		},
		{
			Name:   "move_back_3",
			Row:    7,
			Column: 3,
		},
		{
			Name:   "move_back_4",
			Row:    7,
			Column: 4,
		},
		{
			Name:   "move_back_5",
			Row:    7,
			Column: 5,
		},
		{
			Name:   "move_back_6",
			Row:    7,
			Column: 6,
		},
		{
			Name:   "move_back_side_1",
			Row:    8,
			Column: 1,
		},
		{
			Name:   "move_back_side_2",
			Row:    8,
			Column: 2,
		},
		{
			Name:   "move_back_side_3",
			Row:    8,
			Column: 3,
		},
		{
			Name:   "move_back_side_4",
			Row:    8,
			Column: 4,
		},
		{
			Name:   "move_back_side_5",
			Row:    8,
			Column: 5,
		},
		{
			Name:   "move_back_side_6",
			Row:    8,
			Column: 6,
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
