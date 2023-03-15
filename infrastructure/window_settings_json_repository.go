package infrastructure

import "github.com/maladroitthief/entree/domain/settings"

type WindowModel struct {
	Width  int    `json:"WindowWidth"`
	Height int    `json:"WindowHeight"`
	Title  string `json:"WindowTitle"`
}

func (r *SettingsJsonRepository) GetWindowSettings() (settings.WindowSettings, error) {
	s, err := r.GetSettings()
	if err != nil {
		return settings.WindowSettings{}, err
	}

	return s.WindowSettings, nil
}

func (r *SettingsJsonRepository) SetWindowSettings(w settings.WindowSettings) error {
	s, err := r.GetSettings()
	if err != nil {
		return err
	}

	s.WindowSettings = w
	return r.SetSettings(s)
}

func (r *SettingsJsonRepository) unmarshalWindowSettings(m WindowModel) settings.WindowSettings {
	w := settings.WindowSettings{
		Width:  m.Width,
		Height: m.Height,
		Title:  m.Title,
	}
	err := w.Validate()
	if err != nil {
		return settings.DefaultWindowSettings()
	}

	return w
}

func (r *SettingsJsonRepository) marshalWindowSettings(s settings.WindowSettings) WindowModel {
	err := s.Validate()
	if err != nil {
		return r.marshalWindowSettings(settings.DefaultWindowSettings())
	}

	return WindowModel{
		Width:  s.Width,
		Height: s.Height,
		Title:  s.Title,
	}
}
