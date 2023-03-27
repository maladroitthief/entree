package application

import (
	"errors"
	"image"

	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/domain/sprite"
)

var (
	ErrSheetNotFound = errors.New("sprite sheet not in service")
)

type GraphicsService struct {
	log    logs.Logger
	sheets map[string]sprite.SpriteSheet
}

func NewGraphicsService(
	logger logs.Logger,
) *GraphicsService {
	if logger == nil {
		panic("nil graphics logger")
	}

	return &GraphicsService{
		log:    logger,
		sheets: make(map[string]sprite.SpriteSheet),
	}
}

func (svc *GraphicsService) LoadSpriteSheet(ss sprite.SpriteSheet) {
	svc.sheets[ss.GetName()] = ss
}

func (svc *GraphicsService) GetSprite(
	sheetName,
	spriteName string,
) (image.Rectangle, error) {
	ss, ok := svc.sheets[sheetName]
	if !ok {
		return image.Rectangle{}, ErrSheetNotFound
	}

	// TODO: potentially cache the loaded rectangles
	return ss.SpriteRectangle(spriteName)
}

func (svc *GraphicsService) GetSpriteSheet(
	sheetName string,
) (sprite.SpriteSheet, error) {
	ss, ok := svc.sheets[sheetName]
	if !ok {
		return ss, ErrSheetNotFound
	}

	return ss, nil
}
