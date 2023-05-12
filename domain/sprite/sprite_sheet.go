package sprite

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

type spriteSheet struct {
	Name       string
	Image      image.Image
	Rows       int
	Columns    int
	Offset     int
	SpriteSize int
	Sprites    map[string]Sprite
}

type SpriteSheet interface {
	AddSprite(s Sprite) error
	SpriteRectangle(name string) (image.Rectangle, error)
	GetName() string
	GetImage() image.Image
}

func NewSpriteSheet(
	name string,
	image image.Image,
	rows int,
	columns int,
	offset int,
	size int,
) (SpriteSheet, error) {
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

	return &spriteSheet{
		Name:       name,
		Image:      image,
		Rows:       rows,
		Columns:    columns,
		Offset:     offset,
		SpriteSize: size,
		Sprites:    make(map[string]Sprite),
	}, nil
}

func (ss *spriteSheet) AddSprite(s Sprite) error {
	if s.Row > ss.Rows || s.Row <= 0 {
		return ErrSpriteRowOOB
	}

	if s.Column > ss.Columns || s.Column <= 0 {
		return ErrSpriteColumnOOB
	}

	ss.Sprites[s.Name] = s

	return nil
}

func (ss *spriteSheet) GetName() string {
	return ss.Name
}

func (ss *spriteSheet) GetImage() image.Image {
	return ss.Image
}

func (ss *spriteSheet) SpriteRectangle(name string) (image.Rectangle, error) {
	s, ok := ss.Sprites[name]
	if !ok {
		return image.Rectangle{}, ErrSpriteNotFound
	}

	startingX := (s.Column-1)*ss.SpriteSize + (s.Column-1)*ss.Offset
	startingY := (s.Row-1)*ss.SpriteSize + (s.Row-1)*ss.Offset

	point_1 := image.Pt(startingX, startingY)
	point_2 := image.Pt(point_1.X+ss.SpriteSize, point_1.Y+ss.SpriteSize)

	return image.Rectangle{Min: point_1, Max: point_2}, nil
}
