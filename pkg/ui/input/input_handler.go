package input

import "github.com/maladroitthief/entree/pkg/engine/action"

type InputHandler interface {
	Update()
	Load()
	Reset()

	BindKey(action.Action)
	BindMouseButton(action.Action)
	BindGamepadButton(action.Action)
}
