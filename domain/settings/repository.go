package settings

type Repository interface {
  GetSettings() (Settings, error)
}
