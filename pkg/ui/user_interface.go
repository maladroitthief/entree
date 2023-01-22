package ui

import (
	"github.com/maladroitthief/entree/pkg/ui/input"
	"github.com/maladroitthief/entree/pkg/ui/scene"
)

type UserInterface struct {
	input *input.Input
	sm    *scene.SceneManager
}

func NewUserInterface() *UserInterface {
	ui := &UserInterface{}

	ui.input = input.NewInput()
	ui.sm = scene.NewSceneManager()

	return ui
}

func (ui *UserInterface) Update() error {
	if ui.sm == nil {
		ui.sm = scene.NewSceneManager()
	}

	ui.input.Update()
	err := ui.sm.Update(ui.input)

	return err
}
