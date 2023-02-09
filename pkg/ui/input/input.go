package input

type virtualAction int

const (
	Accept virtualAction = iota
  Cancel

	Up
	Down
	Left
	Right

  Attack
  Dodge
  Interact
  UseItem
  UseConsumable

  Map
  Menu
)

var virtualActions = []virtualAction{
	Accept,
  Cancel,
	Up,
	Down,
	Left,
	Right,
  Attack,
  Dodge,
  Interact,
  UseItem,
  UseConsumable,
  Map,
  Menu,
}

type Input struct {
	keyboardConfig *keyboardConfig
	actionStates   map[virtualAction]int
}

func NewInput() *Input {
	i := &Input{}
	i.keyboardConfig = NewKeyboardConfig()

	return i
}

func (i *Input) Update() {
	if i.actionStates == nil {
		i.actionStates = map[virtualAction]int{}
	}

	for _, a := range virtualActions {
		if i.keyboardConfig.IsPressed(a) {
			i.actionStates[a]++
			continue
		}

		i.actionStates[a] = 0
	}
}

func (i *Input) actionState(a virtualAction) int {
	if i.actionStates == nil {
		return 0
	}

	return i.actionStates[a]
}

func (i *Input) IsAnyAction() bool {
	return i.keyboardConfig.IsAnyKey()
}
