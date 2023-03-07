package infrastructure

type GameMemoryRepository struct {
	CursorX int
	CursorY int
	Inputs  []string
}

func NewGameMemoryRepository() *GameMemoryRepository {
	return &GameMemoryRepository{
    CursorX: 0,
    CursorY: 0,
    Inputs: []string{},
  }
}

func (r *GameMemoryRepository) SetCursor(x, y int) {
  r.CursorX = x
  r.CursorY = y
}

func (r *GameMemoryRepository) GetCursor() (x, y int) {
  return r.CursorX, r.CursorY
}

func (r *GameMemoryRepository) SetInputs(inputs []string) {
  r.Inputs = inputs
}

func (r *GameMemoryRepository) GetInputs() []string {
  return r.Inputs
}
