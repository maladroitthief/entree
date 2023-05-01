package player

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics/collision"
)

func NewPilot(
	input canvas.InputComponent,
) *canvas.Entity {
	physics := NewPlayerPhysicsComponent()
	graphics := NewPlayerGraphicsComponent(
		map[string]int{
			"idle_front":      6,
			"idle_front_side": 4,
			"idle_back":       6,
			"idle_back_side":  4,
			"move_front":      6,
			"move_front_side": 6,
			"move_back":       6,
			"move_back_side":  6,
		},
	)

	return &canvas.Entity{
		Size:         collision.Vector{X: 32, Y: 32},
		Position:     collision.Vector{X: 100, Y: 100},
		Sheet:        "pilot",
		Sprite:       "idle_front_1",
		State:        "idle",
		StateCounter: 0,
		OrientationX: canvas.Neutral,
		OrientationY: canvas.South,
		Components: []canvas.Component{
			input,
			physics,
			graphics,
		},
		Input:    input,
		Physics:  physics,
		Graphics: graphics,
	}
}
