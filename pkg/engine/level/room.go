package level

import (
	"errors"
)

type Position int

const (
	RoomHeight   = 9
	RoomWidth    = 9
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
			"000000000" +
			"000000000" +
			"0000P0000" +
			"000000000" +
			"000000000" +
			"000000000" +
			"111111111" +
			"111111111",
	}
}

func (rf *roomFactory) Exit() Room {
	return Room{
		layout: "" +
			"111111111" +
			"111111111" +
			"000000000" +
			"000000000" +
			"000000000" +
			"000000000" +
			"000000000" +
			"000000000" +
			"000000000",
	}
}

func (rf *roomFactory) Room() Room {
	return Room{
		layout: "" +
			"110000011" +
			"110000011" +
			"000000000" +
			"000010000" +
			"000111000" +
			"000010000" +
			"000000000" +
			"110000011" +
			"110000011",
	}
}

func (rf *roomFactory) PathRoom(paths ...Direction) Room {
	return Room{
		layout: "" +
			"110000011" +
			"110000011" +
			"000000000" +
			"000000000" +
			"000000000" +
			"000000000" +
			"000000000" +
			"110000011" +
			"110000011",
	}
}
