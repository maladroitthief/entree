package input

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
	keyboardConfig *keyboardConfig
	actionStates   map[action]int
}

func NewInput() *Input {
	i := &Input{}
	i.keyboardConfig = NewKeyboardConfig()

	return i
}

func (i *Input) Update() {
	if i.actionStates == nil {
		i.actionStates = map[action]int{}
	}

	for _, a := range actions {
		if i.keyboardConfig.IsPressed(a) {
			i.actionStates[a]++
			continue
		}

		i.actionStates[a] = 0
	}
}

func (i *Input) actionState(a action) int {
	if i.actionStates == nil {
		return 0
	}

	return i.actionStates[a]
}

func (i *Input) IsAnyAction() bool {
	return i.keyboardConfig.IsAnyKey()
}
