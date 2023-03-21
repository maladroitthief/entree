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
	repo   sprite.Repository
	sheets map[string]sprite.SpriteSheet
}

func NewGraphicsService(
	logger logs.Logger,
	repo sprite.Repository,
) *GraphicsService {
	if logger == nil {
		panic("nil graphics logger")
	}

	if repo == nil {
		panic("nil graphics repo")
	}

	return &GraphicsService{
		log:    logger,
		repo:   repo,
		sheets: make(map[string]sprite.SpriteSheet),
	}
}

func (svc *GraphicsService) LoadSpriteSheet(path string) error {
	ss, err := svc.repo.GetSpriteSheet(path)
	if err != nil {
		return err
	}

	svc.sheets[ss.GetName()] = ss

	return nil
}

func (svc *GraphicsService) GetSprite(
	sheetName,
	spriteName string,
) (image.Rectangle, error) {
	ss, ok := svc.sheets[sheetName]
	if ok != true {
		return image.Rectangle{}, ErrSheetNotFound
	}

	return ss.Draw(spriteName)
}
