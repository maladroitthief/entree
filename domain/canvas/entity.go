package canvas

import "image"

type Entity struct {
	Width  int
	Height int
	X      int
	Y      int
	Sheet  string
	State  string
}

func (e *Entity) CurrentSprite(gSvc GraphicsService) (image.Rectangle, error) {
	return gSvc.GetSprite(e.Sheet, e.State)
}
