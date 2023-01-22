package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type action int

const (
	Ok action = iota
	Quit

	// Directional inputs
	Up
	Down
	Left
	Right
)

var actions = []action{
  Ok,
  Quit,
  Up,
  Down,
  Left,
  Right,
}

type Input struct {
	gamepadIDs     []ebiten.GamepadID
	keyboardConfig *keyboardConfig
}

func NewInput() *Input {
  i := &Input{}
  i.keyboardConfig = NewKeyboardConfig()

  return i
}

func (i *Input) Update() {
  return
}
