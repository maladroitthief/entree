package driver

import (
	"encoding/json"
	"os"

	"github.com/maladroitthief/entree/pkg/ui"
)

type Settings struct {
	WindowSettings ui.WindowSettings
	InputSettings  ui.InputSettings
}

type SettingsModel struct {
	Window WindowModel `json:"WindowSettings"`
	Input  InputModel  `json:"InputSettings"`
}

type SettingsRepository struct {
	filePath string
}

func NewSettingsRepository(filePath string) *SettingsRepository {
	r := &SettingsRepository{
		filePath: filePath,
	}
	r.initializeIfNeeded()

	return r
}

func (r *SettingsRepository) initializeIfNeeded() error {
	_, err := os.Stat(r.filePath)
	if err == nil {
		return nil
	}

	return r.SetSettings(r.settingsDefaults())
}

func (r *SettingsRepository) settingsDefaults() Settings {
	return Settings{
		WindowSettings: ui.DefaultWindowSettings(),
		InputSettings:  ui.DefaultInputSettings(),
	}

}

func (r *SettingsRepository) GetSettings() (Settings, error) {
	err := r.initializeIfNeeded()
	if err != nil {
		return Settings{}, err
	}

	jsonContent, err := os.ReadFile(r.filePath)
	if err != nil {
		return Settings{}, err
	}

	s := SettingsModel{}
	err = json.Unmarshal(jsonContent, &s)
	if err != nil {
		return Settings{}, err
	}

	return r.unmarshalSettings(s), nil
}

func (r *SettingsRepository) SetSettings(s Settings) error {
	settingsModel := r.marshalSettings(s)

	jsonContent, err := json.Marshal(settingsModel)
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, jsonContent, 0644)
}

func (r *SettingsRepository) unmarshalSettings(m SettingsModel) Settings {
	return Settings{
		WindowSettings: r.unmarshalWindowSettings(m.Window),
		InputSettings:  r.unmarshalInputSettings(m.Input),
	}
}

func (r *SettingsRepository) marshalSettings(s Settings) SettingsModel {
	return SettingsModel{
		Window: r.marshalWindowSettings(s.WindowSettings),
		Input:  r.marshalInputSettings(s.InputSettings),
	}
}
