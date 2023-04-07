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
  e.DeltaReset()

	for _, input := range i.inputSvc.CurrentInputs() {
		switch input {
		case settings.MoveUp:
      e.DeltaY--
		case settings.MoveDown:
      e.DeltaY++
		case settings.MoveRight:
      e.DeltaX++
		case settings.MoveLeft:
      e.DeltaX--
		}
	}
}

