package level

import (
	"errors"
)

type Position int

const (
	RoomHeight = 9
	RoomWidth  = 9
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

func NewRoomFactory() RoomFactory {
	rf := &roomFactory{}

	return rf
}

type Room struct {
	layout string
}

type roomFactory struct {
}

func (rf *roomFactory) Entrance() Room {
	return Room{
		layout: "" +
			"000000000" +
			"000500000" +
			"000000500" +
			"0500@0000" +
			"000000000" +
			"000000500" +
			"000000000" +
			"111111111" +
			"111111111",
	}
}

func (rf *roomFactory) Exit() Room {
	return Room{
		layout: "" +
			"111111111" +
			"011111110" +
			"010050010" +
			"010101010" +
			"5101e1015" +
			"010111010" +
			"010000010" +
			"000505000" +
			"000050000",
	}
}

func (rf *roomFactory) Room() Room {
	return Room{
		layout: "" +
			"110502011" +
			"110005021" +
			"000502050" +
			"005000000" +
			"002000500" +
			"050020020" +
			"000000502" +
			"120502011" +
			"112000011",
	}
}

func (rf *roomFactory) PathRoom(paths ...Direction) Room {
	return Room{
		layout: "" +
			"110000011" +
			"110050011" +
			"005000050" +
			"005000500" +
			"000050000" +
			"005000500" +
			"000000050" +
			"115005011" +
			"110000011",
	}
}
