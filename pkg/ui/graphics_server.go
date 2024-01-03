package ui

import (
	"errors"
	"image"

	"github.com/maladroitthief/entree/pkg/content"
)

var (
	ErrSheetNotFound = errors.New("sprite sheet not loaded")
)

type GraphicsServer struct {
	sheets map[string]*content.SpriteSheet
}

func NewGraphicsServer() (*GraphicsServer, error) {
	return &GraphicsServer{
		sheets: make(map[string]*content.SpriteSheet),
	}, nil
}

func (svc *GraphicsServer) LoadSpriteSheet(ss *content.SpriteSheet) {
	svc.sheets[ss.Name()] = ss
}

func (svc *GraphicsServer) Sprite(
	sheetName,
	spriteName string,
) (image.Rectangle, error) {
	ss, ok := svc.sheets[sheetName]
	if !ok {
		return image.Rectangle{}, ErrSheetNotFound
	}

	return ss.SpriteRectangle(spriteName)
}

func (svc *GraphicsServer) SpriteSheet(
	sheetName string,
) (*content.SpriteSheet, error) {
	ss, ok := svc.sheets[sheetName]
	if !ok {
		return ss, ErrSheetNotFound
	}

	return ss, nil
}
