package adapter_test

import (
	"image"

	"github.com/maladroitthief/entree/domain/sprite"
)

type spriteSheet struct {
}

func (ss *spriteSheet) AddSprite(s sprite.Sprite) error {
	return nil
}

func (ss *spriteSheet) SpriteRectangle(name string) (image.Rectangle, error) {
	return image.Rectangle{}, nil
}

func (ss *spriteSheet) GetName() string {
	return ""
}

func (ss *spriteSheet) GetImage() image.Image {
	return image.Rectangle{}
}
