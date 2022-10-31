package service

func MatchRoom(p *Player, gameId int, gameMode int) *Room {
	var r *Room
	switch gameMode {
	case 1:
		r = GetRoomMatchLogic("commonMatch")(p)
	default:
		r = GetRoomMatchLogic("commonMatch")(p)
	}
	return r
}

func SelectRoomService(gameId int, gameMode int) *RoomService {
	var rs *RoomService
	switch gameMode {
	case 1:
		rs = GetRSSelectLogic("commonSelect")()
	default:
		rs = GetRSSelectLogic("commonSelect")()
	}
	return rs
}
