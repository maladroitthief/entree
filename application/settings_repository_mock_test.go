package application_test

import "github.com/maladroitthief/entree/domain/settings"

type settingsRepository struct {
}

func (r *settingsRepository) GetSettings() (settings.Settings, error) {
	return settings.Settings{}, nil
}

func (r *settingsRepository) SetSettings(settings.Settings) error {
	return nil
}

func (r *settingsRepository) GetWindowSettings() (settings.WindowSettings, error) {
	return settings.WindowSettings{}, nil
}

func (r *settingsRepository) SetWindowSettings(settings.WindowSettings) error {
	return nil
}

func (r *settingsRepository) GetInputSettings() (settings.InputSettings, error) {
	return settings.InputSettings{}, nil
}

func (r *settingsRepository) SetInputSettings(settings.InputSettings) error {
	return nil
}
