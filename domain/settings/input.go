package settings

import "errors"

type Input string

const (
	Accept        Input = "accept"
	Cancel        Input = "cancel"
	MoveUp        Input = "move_up"
	MoveDown      Input = "move_down"
	MoveLeft      Input = "move_left"
	MoveRight     Input = "move_right"
	Attack        Input = "attack"
	AttackUp      Input = "attack_up"
	AttackDown    Input = "attack_down"
	AttackLeft    Input = "attack_left"
	AttackRight   Input = "attack_right"
	Dodge         Input = "dodge"
	Interact      Input = "interact"
	UseItem       Input = "use_item"
	UseConsumable Input = "use_consumable"
	Map           Input = "map"
	Menu          Input = "menu"
	Err           Input = ""
)

var (
	ErrUnboundKey = errors.New("key is unbound")
  ErrNilKeyboard = errors.New("keyboard is nil")
)

type InputSettings struct {
	Keyboard map[Input]string
}

func (i *InputSettings) Validate() error {
	if i.Keyboard == nil {
    return ErrNilKeyboard
	}

	return nil
}

func AllInputs() []Input {
	return []Input{
		Accept,
		Cancel,
		MoveUp,
		MoveDown,
		MoveLeft,
		MoveRight,
		Attack,
		AttackUp,
		AttackDown,
		AttackLeft,
		AttackRight,
		Dodge,
		Interact,
		UseItem,
		UseConsumable,
		Map,
		Menu,
	}
}

func DefaultInputSettings() InputSettings {
	return InputSettings{
		Keyboard: DefaultKeyboard(),
	}
}

func DefaultKeyboard() map[Input]string {
	return map[Input]string{
		Accept:        "Enter",
		Cancel:        "Backspace",
		MoveUp:        "W",
		MoveDown:      "S",
		MoveLeft:      "A",
		MoveRight:     "D",
		Attack:        "MouseButtonLeft",
		AttackUp:      "ArrowUp",
		AttackDown:    "ArrowDown",
		AttackLeft:    "ArrowLeft",
		AttackRight:   "ArrowRight",
		Dodge:         "Space",
		Interact:      "E",
		UseItem:       "MouseButtonRight",
		UseConsumable: "Q",
		Map:           "M",
		Menu:          "KeyEscape",
	}
}

type InputService interface {
  CurrentInputs() []Input
  IsAny() bool
  IsPressed(i Input) bool
  IsJustPressed(i Input) bool
  GetCursor() (x, y int)
}
