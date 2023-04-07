package input

import (
	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/settings"
)

type PlayerInputComponent struct {
	inputSvc settings.InputService
}

func NewPlayerInputComponent(inputSvc settings.InputService) *PlayerInputComponent {
	pic := &PlayerInputComponent{
    inputSvc: inputSvc,
  }

	return pic
}

func (i *PlayerInputComponent) Update(e *canvas.Entity) {
  e.Reset()

	for _, input := range i.inputSvc.CurrentInputs() {
		switch input {
		case settings.MoveUp:
      e.State = "move"
      e.OrientationY = canvas.North
      e.DeltaY--
		case settings.MoveDown:
      e.State = "move"
      e.OrientationY = canvas.South
      e.DeltaY++
		case settings.MoveRight:
      e.State = "move"
      e.OrientationX = canvas.East
      e.DeltaX++
		case settings.MoveLeft:
      e.State = "move"
      e.OrientationX = canvas.West
      e.DeltaX--
		}
	}
}

