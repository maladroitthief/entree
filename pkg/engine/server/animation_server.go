package server

import (
	"github.com/maladroitthief/entree/pkg/engine/attribute"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

type AnimationServer struct {
}

func NewAnimationServer() *AnimationServer {
	s := &AnimationServer{}

	return s
}

func (s *AnimationServer) Update(e *core.ECS) {
	animations := e.GetAllAnimations()

	for _, a := range animations {
		if a.Static {
			continue
		}

		state, err := e.GetState(a.EntityId)
		if err != nil {
			continue
		}

		spriteName := state.State
		if state.OrientationY == attribute.South {
			spriteName = spriteName + "_front"
		} else {
			spriteName = spriteName + "_back"
		}

		if state.OrientationX != attribute.Neutral {
			spriteName = spriteName + "_side"
		}

		sprites, ok := a.Sprites[spriteName]
		if !ok {
			continue
		}

		spritesCount := len(a.Sprites[spriteName])
		speed := float64(a.Counter) / (a.Speed / float64(spritesCount))
		a.Variant = int(speed) % spritesCount
		a.Sprite = sprites[a.Variant]
		if int(speed) >= spritesCount {
			a.Counter = 0
		} else {
			a.Counter++
		}
		e.SetAnimation(a)
	}
}
