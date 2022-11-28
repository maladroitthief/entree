package screen

import "github.com/hajimehoshi/ebiten/v2"

type screenManager struct {
	screenTitle  string
	screenHeight int
	screenWidth  int
}

type ScreenManager interface {
	Update()
}

func NewScreenManager(title string, width int, height int) ScreenManager {
	sm := &screenManager{
		screenTitle:  title,
		screenWidth:  width,
		screenHeight: height,
	}

	return sm
}

func (sm *screenManager) Update() {
	ebiten.SetWindowTitle(sm.screenTitle)
	ebiten.SetWindowSize(sm.screenWidth, sm.screenHeight)
}
