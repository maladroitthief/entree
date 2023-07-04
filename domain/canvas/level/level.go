package level

type Direction int

const (
	DefaultLevelWidth  = 9
	DefaultLevelHeight = 9

	North Direction = iota
	South
	East
	West
)

type Level struct {
	StartFrom Direction
	Rooms     [][]*Room
}

func NewLevel() *Level {
	l := &Level{}

	return l
}

func (l *Level) GenerateRooms() {

}
