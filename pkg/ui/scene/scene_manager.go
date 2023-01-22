package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/ui/input"
)

const (
	transitionMaxCount = 20
	ScreenWidth        = 256
	ScreenHeight       = 240
)

var (
	transitionFrom = ebiten.NewImage(ScreenWidth, ScreenHeight)
	transitionTo   = ebiten.NewImage(ScreenWidth, ScreenHeight)
)

type SceneManager struct {
	current         Scene
	next            Scene
	transitionCount int
}

func NewSceneManager() *SceneManager {
	sm := &SceneManager{}

	return sm
}

func (s *SceneManager) Update(input *input.Input) error {
	if s.transitionCount <= 0 {
    return nil
		//return s.current.Update(
		//	&GameState{
		//		SceneManager: s,
		//		Input:        input,
		//	},
		//)
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
