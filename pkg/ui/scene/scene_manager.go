package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/ui/input"
	"github.com/maladroitthief/entree/pkg/ui/window"
)

const (
	transitionMaxCount = 20
)

type SceneManager struct {
	input input.InputHandler
	wm    *window.WindowManager

	current         Scene
	next            Scene
	transitionCount int

	debug bool
}

func NewSceneManager() *SceneManager {
	sm := &SceneManager{}

	return sm
}

func (s *SceneManager) Update() error {
	if s.transitionCount <= 0 {
		return s.current.Update(
			&GameState{
				SceneManager: s,
				Input:        s.input,
			},
		)
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return nil
	}

	s.current = s.next
	s.next = nil

	return nil
}

func (s *SceneManager) Draw(r *ebiten.Image) {
	transitionFrom := ebiten.NewImage(s.wm.GetWidth(), s.wm.GetHeight())
	transitionTo := ebiten.NewImage(s.wm.GetWidth(), s.wm.GetHeight())

	if s.transitionCount == 0 {
		s.current.Draw(r)
		return
	}

	transitionFrom.Clear()
	s.current.Draw(transitionFrom)

	transitionTo.Clear()
	s.next.Draw(transitionTo)

	r.DrawImage(transitionFrom, nil)

	alpha := 1 - float64(s.transitionCount)/float64(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	r.DrawImage(transitionTo, op)
}

func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}
}

func (s *SceneManager) SetInput(i input.InputHandler) {
	s.input = i
}

func (s *SceneManager) SetWindowManager(w *window.WindowManager) {
	s.wm = w
}

func (s *SceneManager) SetDebug(b bool) {
	s.debug = b
}

func (s *SceneManager) GetDebug() bool {
	return s.debug
}
