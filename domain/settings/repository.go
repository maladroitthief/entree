package settings

type Repository interface {
  GetSettings() (Settings, error)
  SetSettings(Settings) error
  GetWindowSettings() (WindowSettings, error)
  SetWindowSettings(WindowSettings) error
  GetInputSettings() (InputSettings, error)
  SetInputSettings(InputSettings) error
}
