package content

import (
	"errors"
	"image"
)

var (
	ErrNameBlank            = errors.New("name cannot be blank")
	ErrRowSize              = errors.New("rows must be at least 1")
	ErrColumnSize           = errors.New("columns must be at least 1")
	ErrSpriteSize           = errors.New("sprite size must be at least 1")
	ErrSpriteRowOOB         = errors.New("sprite row is out of sheet bounds")
	ErrSpriteColumnOOB      = errors.New("sprite column is out of sheet bounds")
	ErrSpriteOffsetNegative = errors.New("sprite offset is negative")
	ErrSpriteNotFound       = errors.New("sprite does not exist in this sheet")
)

type Sprite struct {
	Name   string
	Row    int
	Column int
}

type SpriteSheet struct {
	name       string
	image      image.Image
	rows       int
	columns    int
	offset     int
	spriteSize int
	sprites    map[string]Sprite
}

func NewSpriteSheet(
	name string,
	image image.Image,
	rows int,
	columns int,
	offset int,
	size int,
) (*SpriteSheet, error) {
	if name == "" {
		return nil, ErrNameBlank
	}

	if rows < 1 {
		return nil, ErrRowSize
	}

	if columns < 1 {
		return nil, ErrColumnSize
	}

	if size < 1 {
		return nil, ErrSpriteSize
	}

	if offset < 0 {
		return nil, ErrSpriteOffsetNegative
	}

	return &SpriteSheet{
		name:       name,
		image:      image,
		rows:       rows,
		columns:    columns,
		offset:     offset,
		spriteSize: size,
		sprites:    make(map[string]Sprite),
	}, nil
}

func (ss *SpriteSheet) AddSprite(s Sprite) error {
	if s.Row > ss.rows || s.Row <= 0 {
		return ErrSpriteRowOOB
	}

	if s.Column > ss.columns || s.Column <= 0 {
		return ErrSpriteColumnOOB
	}

	ss.sprites[s.Name] = s

	return nil
}

func (ss *SpriteSheet) Name() string {
	return ss.name
}

func (ss *SpriteSheet) Image() image.Image {
	return ss.image
}

func (ss *SpriteSheet) SpriteRectangle(name string) (image.Rectangle, error) {
	s, ok := ss.sprites[name]
	if !ok {
		return image.Rectangle{}, ErrSpriteNotFound
	}

	startingX := (s.Column-1)*ss.spriteSize + (s.Column-1)*ss.offset
	startingY := (s.Row-1)*ss.spriteSize + (s.Row-1)*ss.offset

	point_1 := image.Pt(startingX, startingY)
	point_2 := image.Pt(point_1.X+ss.spriteSize, point_1.Y+ss.spriteSize)

	return image.Rectangle{Min: point_1, Max: point_2}, nil
}
