package driver

import (
	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/entree/pkg/ui"
)

type InputModel struct {
	Keyboard map[core.Input]string `json:"Keyboard"`
}

func (r *SettingsRepository) GetInputSettings() (ui.InputSettings, error) {
	s, err := r.GetSettings()
	if err != nil {
		return ui.InputSettings{}, err
	}

	return s.InputSettings, nil
}

func (r *SettingsRepository) SetInputSettings(i ui.InputSettings) error {
	s, err := r.GetSettings()
	if err != nil {
		return err
	}

	s.InputSettings = i
	return r.SetSettings(s)
}

func (r *SettingsRepository) unmarshalInputSettings(m InputModel) ui.InputSettings {
	i := ui.InputSettings{
		Keyboard: m.Keyboard,
	}

	err := i.Validate()
	if err != nil {
		return ui.DefaultInputSettings()
	}

	return i
}

func (r *SettingsRepository) marshalInputSettings(s ui.InputSettings) InputModel {
	err := s.Validate()
	if err != nil {
		return r.marshalInputSettings(ui.DefaultInputSettings())
	}

	return InputModel{
		Keyboard: s.Keyboard,
	}
}
