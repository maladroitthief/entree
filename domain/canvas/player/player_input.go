package player

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

func (i *PlayerInputComponent) Update(e canvas.Entity) {
	canvas.ResetEntity(e)

	for _, input := range i.inputSvc.CurrentInputs() {
		switch input {
		case settings.MoveUp:
			e.SetState("move")
			e.SetOrientationY(canvas.North)
			e.Send("DeltaPositionY", "-1")
		case settings.MoveDown:
			e.SetState("move")
			e.SetOrientationY(canvas.South)
			e.Send("DeltaPositionY", "1")
		case settings.MoveRight:
			e.SetState("move")
			e.SetOrientationX(canvas.East)
			e.Send("DeltaPositionX", "1")
		case settings.MoveLeft:
			e.SetState("move")
			e.SetOrientationX(canvas.West)
			e.Send("DeltaPositionX", "-1")
		}
	}
}

func (i *PlayerInputComponent) Receive(e canvas.Entity, msg, val string) {
}
