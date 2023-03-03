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


func (s JsonRepository) SaveSettings() {
  
}

func (s JsonRepository) LoadSettings() {

}
