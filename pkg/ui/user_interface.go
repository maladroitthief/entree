package ui

import (
	"github.com/maladroitthief/entree/pkg/ui/input"
	"github.com/maladroitthief/entree/pkg/ui/scene"
)

type UserInterface struct {
	input        *input.Input
	sceneManager *scene.SceneManager
}

func NewUserInterface() *UserInterface {
  ui := &UserInterface{}

  return ui
}

func (ui *UserInterface) Update() error {
	if ui.sceneManager == nil {
		ui.sceneManager = scene.NewSceneManager()
	}

	ui.input.Update()
	err := ui.sceneManager.Update(ui.input)

	return err
}

func (ui *UserInterface) SetInput(i *input.Input) {
	ui.input = i
}

func (ui *UserInterface) SetSceneManager(sm *scene.SceneManager) {
	ui.sceneManager = sm
}
