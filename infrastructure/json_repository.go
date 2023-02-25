package infrastructure

type JsonRepository struct {
	filePath string
}

func NewJsonRepository(filePath string) JsonRepository {
	s := JsonRepository{
		filePath: filePath,
	}

	return s
}

type SettingsModel struct {
  Test string `json:"Test"`
}

func (s JsonRepository) SaveSettings() {
  
}

func (s JsonRepository) LoadSettings() {

}
