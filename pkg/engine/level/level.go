package level

import (
	"math/rand"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

var (
	pathDirections = [3]Direction{North, West, East}
)

type Level struct {
	Rooms        [][]Room
	roomFactory  RoomFactory
	blockFactory BlockFactory
	player       core.Entity
	width        int
	height       int
	size         int
}

func NewLevel(rf RoomFactory, bf BlockFactory, player core.Entity, width, height, size int) *Level {
	l := &Level{
		roomFactory:  rf,
		blockFactory: bf,
		player:       player,
		width:        width,
		height:       height,
		size:         size,
	}
	return l
}

func (l *Level) SetSize(width, height int) {
	l.width = width
	l.height = height
}

func (l *Level) Size() (width, height int) {
	return l.width, l.height
}

func (l *Level) GenerateRooms() {
	l.Rooms = make([][]Room, l.height)
	for i := range l.Rooms {
		l.Rooms[i] = make([]Room, l.width)
	}

	currentX := rand.Intn(l.width)
	currentY := 0
	l.Rooms[currentY][currentX] = l.roomFactory.Exit()
	l.addPathRooms(currentX, currentY)
	l.fillRemainingRooms()
}

func (l *Level) addPathRooms(x, y int) {
	nextX, nextY := x, y
	if nextX < 0 {
		nextX = 0
	}

	if nextX >= l.width {
		nextX = l.width - 1
	}

	if y >= l.height-1 {
		l.Rooms[l.height-1][nextX] = l.roomFactory.Entrance()
		return
	}

	if l.Rooms[nextY][nextX].layout == "" {
		l.Rooms[nextY][nextX] = l.roomFactory.PathRoom()
	}

	nextDirection := pathDirections[rand.Intn(3)]
	switch nextDirection {
	case North:
		l.addPathRooms(nextX, nextY+1)
	case West:
		l.addPathRooms(nextX-1, nextY)
	case East:
		l.addPathRooms(nextX+1, nextY)
	}
}

func (l *Level) fillRemainingRooms() {
	for y := range l.Rooms {
		for x := range l.Rooms[y] {
			if l.Rooms[y][x].layout == "" {
				l.Rooms[y][x] = l.roomFactory.Room()
			}
		}
	}
}

func (l *Level) Render(e *core.ECS) {
	l.SetOutOfBounds(e)
	for y := 0; y < len(l.Rooms); y++ {
		for x := 0; x < len(l.Rooms[y]); x++ {
			for i, block := range l.Rooms[y][x].layout {
				switch block {
				case Player:
					l.blockFactory.AddPlayer(e, l.player, xPosition(x, i), yPosition(y, i))
				case EmptySpace:
				case Solid:
					l.blockFactory.AddSolid(e, xPosition(x, i), yPosition(y, i))
				case Solid50:
					l.blockFactory.AddSolid50(e, xPosition(x, i), yPosition(y, i))
				case Obstacle:
					l.blockFactory.AddObstacle(e, xPosition(x, i), yPosition(y, i))
				case Enemy:
					l.blockFactory.AddEnemy(e, xPosition(x, i), yPosition(y, i))
				}
			}
		}
	}
}

func (l *Level) SetOutOfBounds(e *core.ECS) {
	entities := [4]core.Entity{}
	size := float64(l.size)
	xSize := float64(l.width) * size * RoomWidth
	ySize := float64(l.height) * size * RoomHeight

	positions := [4]data.Vector{
		{X: xSize / 2, Y: -size / 2},
		{X: xSize / 2, Y: ySize + size/2},
		{X: xSize + size/2, Y: ySize / 2},
		{X: -size / 2, Y: ySize / 2},
	}
	sizes := [4]data.Vector{
		{X: xSize, Y: size},
		{X: xSize, Y: size},
		{X: size, Y: ySize},
		{X: size, Y: ySize},
	}

	dimensions := [4]core.Dimension{}
	for i := 0; i < len(entities); i++ {
		dimensions[i] = e.NewDimension(positions[i], sizes[i])
		collider := e.NewCollider()
		collider.ColliderType = core.Immovable
		entity := e.NewEntity()
		entity = e.BindDimension(entity, dimensions[i])
		entity = e.BindCollider(entity, collider)
	}
}

func xPosition(x, blockIndex int) float64 {
	return float64(
		(x*RoomWidth*BlockSize)+(blockIndex%9)*BlockSize,
	) + BlockSize/2
}

func yPosition(y, blockIndex int) float64 {
	return float64(
		(y*RoomHeight*BlockSize)+(blockIndex/9)*BlockSize,
	) + BlockSize/2
}
