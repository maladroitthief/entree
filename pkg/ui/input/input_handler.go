package input

import "github.com/maladroitthief/entree/pkg/engine/core"

type InputHandler interface {
}

type inputHandler struct {
	keyboardConfig *keyboardConfig
	actionStates   map[Input]int
	actionBindings map[Input]core.Action
}

func NewInputHandler() InputHandler {
	ih := &inputHandler{}
	ih.keyboardConfig = NewKeyboardConfig()

	return ih
}

func (ih *inputHandler) Update() []core.Action {
	if ih.actionStates == nil {
		ih.actionStates = map[Input]int{}
	}

	for _, a := range Inputs() {
		if ih.keyboardConfig.IsPressed(a) {
			ih.actionStates[a]++
			continue
		}

		ih.actionStates[a] = 0
	}

	return nil
}


func (ih *inputHandler) BindAction(i Input, a core.Action) {
	if ih.actionBindings == nil {
		ih.actionBindings = map[Input]core.Action{}
	}

	ih.actionBindings[i] = a
}

func (ih *inputHandler) IsAny() bool {
	return ih.keyboardConfig.IsAnyKey()
}

func (ih *inputHandler) actionState(a Input) int {
	if ih.actionStates == nil {
		return 0
	}

	return ih.actionStates[a]
}
