package canvas

import "image"

type Entity struct {
	Width  int
	Height int
	X      int
	Y      int
	Sheet  string
	State  string
	Image  *image.Image
}
