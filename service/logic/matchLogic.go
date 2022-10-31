package logic

import (
	"log"
	"match/service"
)

func MatchRoom(p *service.Player) *service.Room {
	rMgr := service.GetRoomManager()
	var list []*service.Room
	rMgr.FreeRooms.GetRoomList(&list, func(r *service.Room) bool {
		return r.GetLeftPosition() > 0
	})
	log.Println("MatchRoom len(list)", len(list))
	if len(list) > 0 {
		return list[0]
	}
	return nil
}

//注册匹配逻辑

func InitRoomMatchLogic() {
	service.RegisterRoomMatchLogic("commonMatch", MatchRoom)
}
