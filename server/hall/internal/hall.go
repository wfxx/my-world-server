package internal

import (
	"github.com/name5566/leaf/gate"
	g "github.com/name5566/leaf/go"
	"github.com/name5566/leaf/timer"
)

var (
	rooms      = make(map[uint]*Room)
	lastRoomID = uint(0)
)

type Hall struct {
	gate.Agent
	*g.LinearContext
	saveDBTimer *timer.Timer
}

func CreateRooms(roomID uint) *Room {
	room := newRoom(roomID)
	rooms[roomID] = room
	return room
}

func GetRoom(roomID uint) *Room {
	room, ok := rooms[roomID]
	if !ok {
		room = CreateRooms(roomID)
	}
	return room
}

func InitRooms() {
	for i := 0; i < 10; i++ {
		CreateRooms(uint(i))

	}

	//test

}
