package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type virtualGamepadButton int

const (
	Ok virtualGamepadButton = iota
	Quit

	// Directional inputs
	Up
	Down
	Left
	Right
)

type Input struct {
	gamepadIDs     []ebiten.GamepadID
	keyboardConfig keyboardConfig
}

func (i *Input) Update() {
  return
}
