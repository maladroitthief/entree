package level

import (
	"errors"
)

// Room layout
//
// 00A0B0C00
// 000000000
// L0000000D
// 000000000
// K0000000E
// 000000000
// J0000000F
// 000000000
// 00I0H0G00

type Position int

const (
	RoomHeight   = 9
	RoomWidth    = 9
	EmptySpace   = '0'
	SolidBlock   = '1'
	SolidBlock50 = '2'
)

var (
	ErrLayoutHeightMismatch = errors.New("layout size does not match height")
	ErrLayoutWidthMismatch  = errors.New("layout size does not match width")
	ErrLayoutNotSquare      = errors.New("layout is not square")
)

type RoomFactory interface {
	Entrance() Room
	Exit() Room
	PathRoom(paths ...Direction) Room
	Room() Room
}

type Room interface {
	Load()
	Layout() [][]rune
}

type room struct {
	template string
	layout   [][]rune
}

func SampleRoom() Room {
	r := &room{
		template: "" +
			"110000000" +
			"110000000" +
			"110000000" +
			"110000000" +
			"110000000" +
			"110000000" +
			"110000000" +
			"111111111" +
			"111111111",
	}

	r.Load()

	return r
}

func (r *room) Load() {
	for i := 0; i < RoomHeight; i++ {
		copy(
			r.layout[i][:],
			[]rune(r.template[i*RoomWidth:(i*RoomWidth)+RoomWidth]),
		)
	}
}

func (r *room) Layout() [][]rune {
	return r.layout
}
