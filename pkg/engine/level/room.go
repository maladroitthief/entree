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
			"00e00e000" +
			"e0e5000e0" +
			"00ee0e500" +
			"0500@00e0" +
			"00e0000e0" +
			"0e00e0500" +
			"00e000e00" +
			"111111111" +
			"111111111",
	}
}

func (rf *roomFactory) Exit() Room {
	return Room{
		layout: "" +
			"111111111" +
			"111111111" +
			"000555000" +
			"005000500" +
			"005010500" +
			"0050e0500" +
			"000555000" +
			"000000000" +
			"000000000",
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
