package action

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
