package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/maladroitthief/entree/domain/settings"
)

type SettingsModel struct {
	Window WindowModel `json:"WindowSettings"`
}

type WindowModel struct {
	Width  int    `json:"WindowWidth"`
	Height int    `json:"WindowHeight"`
	Title  string `json:"WindowTitle"`
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

func (r *SettingsJsonRepository) GetWindowSettings() (settings.Window, error) {
  s, err := r.GetSettings()
  if err != nil {
    return settings.Window{}, err
  }
  
  return s.Window, nil
}

func (r *SettingsJsonRepository) SetSettings(s settings.Settings) error {
	settingsModel := r.marshalSettings(s)

	jsonContent, err := json.Marshal(settingsModel)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(r.filePath, jsonContent, 0644)
}

func (r *SettingsJsonRepository) SetWindowSettings(w settings.Window) error {
  s, err := r.GetSettings()
  if err != nil {
    return err
  }
  
  s.Window = w
  return r.SetSettings(s)
}

func (r *SettingsJsonRepository) unmarshalSettings(m SettingsModel) settings.Settings {
  return settings.Settings{
    Window: r.unmarshalWindowSettings(m.Window),
  }
}

func (r *SettingsJsonRepository) marshalSettings(s settings.Settings) SettingsModel {
  return SettingsModel{
    Window: r.marshalWindowSettings(s.Window),
  }
}

func (r *SettingsJsonRepository) unmarshalWindowSettings(m WindowModel) settings.Window {
  w := settings.Window{
		Width:  m.Width,
		Height: m.Height,
		Title:  m.Title,
	}
  err := w.Validate()
  if err != nil {
    return settings.WindowDefaults()
  }

  return w
}

func (r *SettingsJsonRepository) marshalWindowSettings(s settings.Window) WindowModel {
  err := s.Validate()
  if err != nil {
    return r.marshalWindowSettings(settings.WindowDefaults())
  }

	return WindowModel{
		Width:  s.Width,
		Height: s.Height,
		Title:  s.Title,
	}
}
