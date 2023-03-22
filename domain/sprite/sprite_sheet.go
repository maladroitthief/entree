package sprite

import (
	"errors"
	"image"
	"os"
)

var (
	ErrNameBlank       = errors.New("name cannot be blank")
	ErrRowSize         = errors.New("rows must be at least 1")
	ErrColumnSize      = errors.New("columns must be at least 1")
	ErrSpriteSize      = errors.New("sprite size must be at least 1")
	ErrSpriteRowOOB    = errors.New("sprite row is out of sheet bounds")
	ErrSpriteColumnOOB = errors.New("sprite column is out of sheet bounds")
	ErrSpriteNotFound  = errors.New("sprite does not exist in this sheet")
)

type spriteSheet struct {
	Name       string
	Path       string
	Rows       int
	Columns    int
	SpriteSize int
	Sprites    map[string]Sprite
}

type SpriteSheet interface {
	AddSprite(s Sprite) error
	DrawRectangle(name string) (image.Rectangle, error)
  GetName() string
}

func NewSpriteSheet(
	name string,
	path string,
	rows int,
	columns int,
	size int,
) (SpriteSheet, error) {
	// Ensure the file opens
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

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

	return &spriteSheet{
		Name:       name,
		Path:       path,
		Rows:       rows,
		Columns:    columns,
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

func (ss *spriteSheet) DrawRectangle(name string) (image.Rectangle, error) {
	s, ok := ss.Sprites[name]
	if ok != true {
		return image.Rectangle{}, ErrSpriteNotFound
	}

	point_1 := image.Pt((s.Column-1)*ss.SpriteSize, (s.Row-1)*ss.SpriteSize)
	point_2 := image.Pt(point_1.X+ss.SpriteSize, point_1.Y+ss.SpriteSize)

	return image.Rectangle{Min: point_1, Max: point_2}, nil
}
