package settings

type Settings struct {
  Window Window
}

func SettingsDefaults() Settings {
	return Settings{
    Window: WindowDefaults(),
  }
}
