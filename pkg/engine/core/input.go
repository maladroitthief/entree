package core

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
	AllInputs = []Input{
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
)
