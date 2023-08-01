package driver

import "github.com/maladroitthief/entree/pkg/ui"

type WindowModel struct {
	Width  int    `json:"WindowWidth"`
	Height int    `json:"WindowHeight"`
	Title  string `json:"WindowTitle"`
}

func (r *SettingsRepository) GetWindowSettings() (ui.WindowSettings, error) {
	s, err := r.GetSettings()
	if err != nil {
		return ui.WindowSettings{}, err
	}

	return s.WindowSettings, nil
}

func (r *SettingsRepository) SetWindowSettings(w ui.WindowSettings) error {
	s, err := r.GetSettings()
	if err != nil {
		return err
	}

	s.WindowSettings = w
	return r.SetSettings(s)
}

func (r *SettingsRepository) unmarshalWindowSettings(m WindowModel) ui.WindowSettings {
	w := ui.WindowSettings{
		Width:  m.Width,
		Height: m.Height,
		Title:  m.Title,
	}
	err := w.Validate()
	if err != nil {
		return ui.DefaultWindowSettings()
	}

	return w
}

func (r *SettingsRepository) marshalWindowSettings(s ui.WindowSettings) WindowModel {
	err := s.Validate()
	if err != nil {
		return r.marshalWindowSettings(ui.DefaultWindowSettings())
	}

	return WindowModel{
		Width:  s.Width,
		Height: s.Height,
		Title:  s.Title,
	}
}
