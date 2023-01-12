package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/ui/input"
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

func (s *SceneManager) Update(input *input.Input) error {
	if s.transitionCount <= 0 {
		return s.current.Update(
			&GameState{
				SceneManager: s,
				Input:        input,
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
	if s.transitionCount == 0 {
    s.current.Draw(r)
    return
	}

  transitionFrom.
}
