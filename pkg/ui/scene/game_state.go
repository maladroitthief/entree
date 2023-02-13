package scene

import "github.com/maladroitthief/entree/pkg/ui/input"

type GameState struct {
	SceneManager *SceneManager
	Input        input.InputHandler
}
