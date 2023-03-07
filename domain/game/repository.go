package game

type Repository interface {
  GetCursor() (x, y int)
  SetCursor(x, y int)
  GetInputs() []string
  SetInputs([]string)
}
