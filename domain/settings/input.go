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
)

type InputSettings struct {
	Keyboard map[Input]string
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
		Accept:        "KeyEnter",
		Cancel:        "KeyBackspace",
		MoveUp:        "KeyW",
		MoveDown:      "KeyS",
		MoveLeft:      "KeyA",
		MoveRight:     "KeyD",
		Attack:        "MouseButtonLeft",
		AttackUp:      "KeyArrowUp",
		AttackDown:    "KeyArrowDown",
		AttackLeft:    "KeyArrowLeft",
		AttackRight:   "KeyArrowRight",
		Dodge:         "KeySpace",
		Interact:      "KeyE",
		UseItem:       "MouseButtonRight",
		UseConsumable: "KeyQ",
		Map:           "KeyM",
		Menu:          "KeyEscape",
	}
}
