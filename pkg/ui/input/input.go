package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	gamepadIDs                 []ebiten.GamepadID
	virtualGamepadButtonStates map[virtualGamepadButton]int
	gamepadConfig              gamepadConfig
}

func (i *Input) Update() {

}
