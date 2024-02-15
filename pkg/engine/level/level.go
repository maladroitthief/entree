package level

import (
	"math/rand"

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
	x            int
	y            int
	size         int
}

func NewLevel(rf RoomFactory, bf BlockFactory, player core.Entity, x, y, size int) *Level {
	l := &Level{
		roomFactory:  rf,
		blockFactory: bf,
		player:       player,
		x:            x,
		y:            y,
		size:         size,
	}
	return l
}

func (l *Level) GenerateRooms() {
	l.Rooms = make([][]Room, l.x)
	for i := range l.Rooms {
		l.Rooms[i] = make([]Room, l.y)
	}

	currentX := rand.Intn(l.x)
	currentY := 0
	l.Rooms[currentX][currentY] = l.roomFactory.Exit()
	l.addPathRooms(currentX, currentY)
	l.fillRemainingRooms()
}

func (l *Level) addPathRooms(x, y int) {
	nextX, nextY := x, y
	if nextX < 0 {
		nextX = 0
	}

	if nextX >= l.x {
		nextX = l.x - 1
	}

	if y >= l.y-1 {
		l.Rooms[nextX][l.y-1] = l.roomFactory.Entrance()
		return
	}

	if l.Rooms[nextX][nextY].layout == "" {
		l.Rooms[nextX][nextY] = l.roomFactory.PathRoom()
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
	for x := range l.Rooms {
		for y := range l.Rooms[x] {
			if l.Rooms[x][y].layout == "" {
				l.Rooms[x][y] = l.roomFactory.Room()
			}
		}
	}
}

func (l *Level) Render(e *core.ECS) {
	for x := 0; x < len(l.Rooms); x++ {
		for y := 0; y < len(l.Rooms[x]); y++ {
			for i, block := range l.Rooms[x][y].layout {
				switch block {
				case Player:
					l.blockFactory.AddPlayer(l.player, xPosition(x, l.size, i), yPosition(y, l.size, i))
				case EmptySpace:
				case Solid:
					l.blockFactory.AddSolid(xPosition(x, l.size, i), yPosition(y, l.size, i))
				case Solid50:
					l.blockFactory.AddSolid50(xPosition(x, l.size, i), yPosition(y, l.size, i))
				case Obstacle:
					l.blockFactory.AddObstacle(xPosition(x, l.size, i), yPosition(y, l.size, i))
				case Enemy:
					l.blockFactory.AddEnemy(xPosition(x, l.size, i), yPosition(y, l.size, i))
				}
			}
		}
	}
}

func xPosition(x, size, blockIndex int) float64 {
	return float64(
		(x*RoomWidth*size)+(blockIndex%9)*size,
	) + float64(size)/2
}

func yPosition(y, size, blockIndex int) float64 {
	return float64(
		(y*RoomHeight*size)+(blockIndex/9)*size,
	) + float64(size)/2
}
