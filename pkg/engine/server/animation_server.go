package server

import (
	"fmt"

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

		speed := float64(a.Counter) / (a.Speed / float64(a.VariantMax))
		a.Variant = int(speed)%a.VariantMax + 1
		a.Sprite = fmt.Sprintf("%s_%d", spriteName, a.Variant)
		if int(speed) >= a.VariantMax {
			a.Counter = 0
		} else {
			a.Counter++
		}
		e.SetAnimation(a)
	}
}
