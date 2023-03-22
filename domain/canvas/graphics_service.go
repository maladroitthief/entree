package canvas

import "image"

type GraphicsService interface {
	GetSprite(sheetName, spriteName string) (image.Rectangle, error)
}
