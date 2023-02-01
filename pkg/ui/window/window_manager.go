package window

import "github.com/hajimehoshi/ebiten/v2"

type WindowManager struct {
	width  int
	height int
	title  string
}

func NewWindowManager(width int, height int, title string) *WindowManager {
	wm := &WindowManager{
		width:  width,
		height: height,
		title:  title,
	}
  
  wm.Update()
	return wm
}

func (w *WindowManager) Update() {
	ebiten.SetWindowSize(w.width, w.height)
	ebiten.SetWindowTitle(w.title)
}

func (w *WindowManager) GetWidth() int {
	return w.width
}

func (w *WindowManager) GetHeight() int {
	return w.height
}

func (w *WindowManager) GetTitle() string {
	return w.title
}

