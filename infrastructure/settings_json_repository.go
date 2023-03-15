package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/maladroitthief/entree/domain/settings"
)

type SettingsModel struct {
	Window WindowModel `json:"WindowSettings"`
	Input  InputModel  `json:"InputSettings"`
}

type SettingsJsonRepository struct {
	filePath string
}

func NewSettingsJsonRepository(filePath string) *SettingsJsonRepository {
	r := &SettingsJsonRepository{
		filePath: filePath,
	}
	r.initializeIfNeeded()

	return r
}

func (r *SettingsJsonRepository) initializeIfNeeded() error {
	_, err := os.Stat(r.filePath)
	if err == nil {
		return nil
	}

	return r.SetSettings(settings.SettingsDefaults())
}

func (r *SettingsJsonRepository) GetSettings() (settings.Settings, error) {
	err := r.initializeIfNeeded()
	if err != nil {
		return settings.Settings{}, err
	}

	jsonContent, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return settings.Settings{}, err
	}

	s := SettingsModel{}
	err = json.Unmarshal(jsonContent, &s)
	if err != nil {
		return settings.Settings{}, err
	}

	return r.unmarshalSettings(s), nil
}

func (r *SettingsJsonRepository) SetSettings(s settings.Settings) error {
	settingsModel := r.marshalSettings(s)

	jsonContent, err := json.Marshal(settingsModel)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(r.filePath, jsonContent, 0644)
}


func (r *SettingsJsonRepository) unmarshalSettings(m SettingsModel) settings.Settings {
	return settings.Settings{
		WindowSettings: r.unmarshalWindowSettings(m.Window),
    InputSettings: r.unmarshalInputSettings(m.Input),
	}
}

func (r *SettingsJsonRepository) marshalSettings(s settings.Settings) SettingsModel {
	return SettingsModel{
		Window: r.marshalWindowSettings(s.WindowSettings),
    Input: r.marshalInputSettings(s.InputSettings),
	}
}

