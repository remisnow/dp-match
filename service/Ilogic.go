package service

import "log"

//==========================================================================================================

var rSSelectLogic = make(map[string]func() *RoomService)

func RegisterSelectLogic(name string, f func() *RoomService) {
	log.Println("RegisterSelectLogic", name)
	rSSelectLogic[name] = f
}

func GetRSSelectLogic(name string) func() *RoomService {
	return rSSelectLogic[name]
}

//==========================================================================================================

var roomMatch = make(map[string]func(p *Player) *Room)

func RegisterRoomMatchLogic(name string, f func(p *Player) *Room) {
	log.Println("RegisterRoomMatchLogic", name)
	roomMatch[name] = f
}

func GetRoomMatchLogic(name string) func(p *Player) *Room {
	return roomMatch[name]
}
