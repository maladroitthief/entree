package settings

type Settings struct {
  WindowSettings WindowSettings
  InputSettings InputSettings
}

func SettingsDefaults() Settings {
	return Settings{
    WindowSettings: DefaultWindowSettings(),
    InputSettings: DefaultInputSettings(),
  }
}
