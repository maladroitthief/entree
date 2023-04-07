package infrastructure

import "github.com/maladroitthief/entree/domain/settings"

type InputModel struct {
	Keyboard map[settings.Input]string `json:"Keyboard"`
}

func (r *SettingsJsonRepository) GetInputSettings() (settings.InputSettings, error) {
	s, err := r.GetSettings()
	if err != nil {
		return settings.InputSettings{}, err
	}

	return s.InputSettings, nil
}

func (r *SettingsJsonRepository) SetInputSettings(i settings.InputSettings) error {
	s, err := r.GetSettings()
	if err != nil {
		return err
	}

	s.InputSettings = i
	return r.SetSettings(s)
}

func (r *SettingsJsonRepository) unmarshalInputSettings(m InputModel) settings.InputSettings {
	i := settings.InputSettings{
		Keyboard: m.Keyboard,
	}

	err := i.Validate()
	if err != nil {
		return settings.DefaultInputSettings()
	}

	return i
}

func (r *SettingsJsonRepository) marshalInputSettings(s settings.InputSettings) InputModel {
	err := s.Validate()
	if err != nil {
		return r.marshalInputSettings(settings.DefaultInputSettings())
	}

	return InputModel{
		Keyboard: s.Keyboard,
	}
}
