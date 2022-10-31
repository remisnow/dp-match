package service

import (
	"github.com/mediocregopher/radix/v3"
	"log"
	"match/config"
	"match/lib/base"
	"match/lib/event"
	"match/lib/redis"
	"strconv"
)

type Room struct {
	base.BaseObj
	event.Event
	PlayerContainer
	GId       string
	RoomId    int64
	service   *RoomService
	roomState int
	roomType  int
	maxCount  int
}

func (r *Room) initRoom(roomId int, roomType int, service *RoomService) {
	GId := config.MatchRedisHead + "Room-" + strconv.Itoa(roomId)
	r.InitBaseObj(GId)
	r.InitialPC()
	r.GId = GId
	r.RoomId = int64(roomId)
	r.roomType = roomType
	r.service = service
	r.maxCount = config.ServiceConfig.RoomPlayerMaxCount
	r.Save()
}

func (r *Room) GetService() *RoomService {
	return r.service
}

func (r *Room) UserWaitEnterRoom(p *Player) {
	r.AddPlayer(p)
	p.OnUserWaitEnterRoom(r.RoomId)
}

func (r *Room) UserEnterRoom(p *Player) {
	if r.GetPlayer(p.UserId) == nil {
		r.AddPlayer(p)
	}
	p.OnUserEnteringRoom(r.RoomId)
	//生成 resp.TicketId
	//写入redis
	//更新房间状态
	//更新redis玩家状态
	//设置玩家过期时间 房间过期时间
}

func (r *Room) Save(params ...interface{}) {
	cli := redis.GetClient("match")
	var err error
	if len(params) > 0 {
		err = cli.Do(radix.FlatCmd(nil, "HMSET", r.GId, params...))
	} else {
		err = cli.Do(radix.FlatCmd(nil, "HMSET", r.GId,
			"RoomId", r.RoomId,
			"service", r.service.serviceName,
			"roomState", r.roomState,
			"roomType", r.roomType,
			"playerCount", r.count,
		))
	}
	if err != nil {
		log.Println("Room.save error", err)
	}
}
func (r *Room) GetLeftPosition() int {
	log.Println("GetLeftPosition maxCount", r.maxCount, r.GetPlayerTotal())
	return r.maxCount - r.GetPlayerTotal()
}

//func Room() *Room { // 实例化对象
//	return &Room{}
//}
