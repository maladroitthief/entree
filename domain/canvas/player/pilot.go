package player

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/physics"
)

func NewPilot(
	inputComponent canvas.InputComponent,
) canvas.Entity {
	physicsComponent := NewPlayerPhysicsComponent()
	graphicsComponent := NewPlayerGraphicsComponent(
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

	return &playerEntity{
		size:         physics.Vector{X: 12, Y: 18},
		position:     physics.Vector{X: 50, Y: 50},
		offset:       physics.Vector{X: 0, Y: -6},
		scale:        canvas.DefaultScale,
		sheet:        "pilot",
		sprite:       "idle_front_1",
		state:        "idle",
		stateCounter: 0,
		orientationX: canvas.Neutral,
		orientationY: canvas.South,
		components: []canvas.Component{
			inputComponent,
			physicsComponent,
			graphicsComponent,
		},
		input:    inputComponent,
		physics:  physicsComponent,
		graphics: graphicsComponent,
	}
}
