package player

import (
	"github.com/maladroitthief/entree/domain/canvas"
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
  
  e := canvas.NewEntity()
  e.SetSize(12, 18)
  e.SetOffset(0, -6)
  e.SetSheet("pilot")
  e.SetSprite("idle_front_1")
  e.SetState("idle")
  e.SetInputComponent(inputComponent)
  e.SetPhysicsComponent(physicsComponent)
  e.SetGraphicsComponent(graphicsComponent)

  return e
}
