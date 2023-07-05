package game_test

import (
	"image"

	"github.com/maladroitthief/entree/domain/sprite"
)

type graphicsService struct {
}

func (svc *graphicsService) LoadSpriteSheet(ss sprite.SpriteSheet) {}

func (svc *graphicsService) GetSprite(sheetName, spriteName string) (image.Rectangle, error) {
	return image.Rectangle{}, nil
}

func (svc *graphicsService) GetSpriteSheet(sheetName string) (sprite.SpriteSheet, error) {
	return &spriteSheet{}, nil
}
