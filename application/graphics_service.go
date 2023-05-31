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

type GraphicsService interface {
	LoadSpriteSheet(ss sprite.SpriteSheet)
	GetSprite(sheetName, spriteName string) (image.Rectangle, error)
	GetSpriteSheet(sheetName string) (sprite.SpriteSheet, error)
}

type graphicsService struct {
	log    logs.Logger
	sheets map[string]sprite.SpriteSheet
}

func NewGraphicsService(logger logs.Logger) GraphicsService {
	if logger == nil {
		panic("nil graphics logger")
	}

	return &graphicsService{
		log:    logger,
		sheets: make(map[string]sprite.SpriteSheet),
	}
}

func (svc *graphicsService) LoadSpriteSheet(ss sprite.SpriteSheet) {
	svc.sheets[ss.GetName()] = ss
}

func (svc *graphicsService) GetSprite(
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

func (svc *graphicsService) GetSpriteSheet(
	sheetName string,
) (sprite.SpriteSheet, error) {
	ss, ok := svc.sheets[sheetName]
	if !ok {
		return ss, ErrSheetNotFound
	}

	return ss, nil
}
