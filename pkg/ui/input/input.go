package input

type Input int

const (
	Accept Input = iota
	Cancel
	MoveUp
	MoveDown
	MoveLeft
	MoveRight
	AttackUp
	AttackDown
	AttackLeft
	AttackRight
	Dodge
	Interact
	UseItem
	UseConsumable
	Map
	Menu
)

func Inputs() []Input {
	return []Input{
		Accept,
		Cancel,
		MoveUp,
		MoveDown,
		MoveLeft,
		MoveRight,
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
