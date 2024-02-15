package core

type Input string

const (
	InputAccept        Input = "accept"
	InputCancel        Input = "cancel"
	InputMoveUp        Input = "move_up"
	InputMoveDown      Input = "move_down"
	InputMoveLeft      Input = "move_left"
	InputMoveRight     Input = "move_right"
	InputAttack        Input = "attack"
	InputAttackUp      Input = "attack_up"
	InputAttackDown    Input = "attack_down"
	InputAttackLeft    Input = "attack_left"
	InputAttackRight   Input = "attack_right"
	InputDodge         Input = "dodge"
	InputInteract      Input = "interact"
	InputUseItem       Input = "use_item"
	InputUseConsumable Input = "use_consumable"
	InputMap           Input = "map"
	InputMenu          Input = "menu"
	InputErr           Input = ""
)

var (
	AllInputs = []Input{
		InputAccept,
		InputCancel,
		InputMoveUp,
		InputMoveDown,
		InputMoveLeft,
		InputMoveRight,
		InputAttack,
		InputAttackUp,
		InputAttackDown,
		InputAttackLeft,
		InputAttackRight,
		InputDodge,
		InputInteract,
		InputUseItem,
		InputUseConsumable,
		InputMap,
		InputMenu,
	}
)
