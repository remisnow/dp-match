package service

import (
	"github.com/mediocregopher/radix/v3"
	"match/config"
	"match/lib/base"
	"match/lib/log"
	"match/lib/redis"
	utilTime "match/lib/utils/time"
	"strconv"
	"time"
)

const (
	PlayerStateWaitEnter = 0 //等待进入房间
	PlayerStateEntering  = 1 //正在进入房间
	PlayerStateInRoom    = 2 //房间内
	PlayerStateOutRoom   = 3 //不在房间
)

type Player struct {
	base.BaseObj
	GId          string
	UserId       int64
	roomId       int64
	IsOnline     int
	Level        int
	Class        int
	Sex          int
	State        int
	timer        utilTime.TimeStopChan
	isGamePlayer bool
}

func (p *Player) initPlayer(playerId int, userId int64) {
	GId := config.MatchRedisHead + "Player-" + strconv.Itoa(playerId)
	p.InitBaseObj(GId)
	p.GId = GId
	p.UserId = userId
	p.IsOnline = 1
	p.Level = 1
	p.Class = 1
	p.Sex = int(userId) % 2
	p.State = PlayerStateOutRoom
	p.isGamePlayer = true
	p.Save()
	//查redis 或者其他方法从斯哥获取 相关信息
}

// OnUserWaitEnterRoom 等待房间创建好
func (p *Player) OnUserWaitEnterRoom(id int64) {
	p.State = PlayerStateWaitEnter
	p.roomId = id
	p.Save("State", p.State, "roomId", p.roomId)
	p.timer = utilTime.TimeTimer(60*time.Second, func() {
		if p == nil {
			return
		}
		if p.State == PlayerStateWaitEnter {
			TimerChan <- func() {
				if p == nil {
					return
				}
				GetRoomManager().GetRoom(p.roomId).RemovePlayer(p)
				p.roomId = 0
				p.State = PlayerStateOutRoom
			}
		}
	})

}

func (p *Player) OnUserEnteringRoom(id int64) {
	if p.State == PlayerStateWaitEnter {
		if id != p.roomId {
			log.Error("OnUserEnteringRoom wrong roomId = ", p.roomId, "id = ", id)
			return
		}
		p.timer <- true
	}
	p.roomId = id
	p.State = PlayerStateEntering
	p.Save("State", p.State, "roomId", p.roomId)
	p.timer = utilTime.TimeTimer(60*time.Second, func() {
		if p == nil {
			return
		}
		if p.State == PlayerStateEntering {
			TimerChan <- func() {
				if p == nil {
					return
				}
				GetRoomManager().GetRoom(p.roomId).RemovePlayer(p)
				p.roomId = 0
				p.State = PlayerStateOutRoom
			}
		}
	})
}
func (p *Player) OnUserEnterRoom() {
	if p.State == PlayerStateEntering {
		p.timer <- true
	}
	p.State = PlayerStateInRoom
	p.Save("State", p.State)

}

func (p *Player) OnUserLeaveRoom() {
	if p.roomId > 0 {
		p.roomId = 0
		p.State = PlayerStateOutRoom
		p.Save("State", p.State, "roomId", p.roomId)
	}

}

func (p *Player) SetState(state int) {
	p.State = state
}

func (p *Player) GetState() int {
	return p.State
}

func (p *Player) SetRoomId(id int64) {
	p.roomId = id
}

func (p *Player) GetRoomId() int64 {
	return p.roomId
}

type ticketData struct {
	UserId int64 `json:"userId"`
	RoomId int64 `json:"roomId"`
	Player bool  `json:"player"`
	Data   struct {
	} `json:"data"`
	Rand int32 `json:"rand"`
}

func (p *Player) GetTicketId(rand int32) (ticketId string) {
	ticketId = p.GId + strconv.Itoa(int(p.roomId)) + strconv.Itoa(int(rand))
	cli := redis.GetClient("match")
	str, strErr := Obj2Json(&ticketData{
		UserId: p.UserId,
		RoomId: p.roomId,
		Player: p.isGamePlayer,
		Data:   struct{}{},
		Rand:   rand,
	})
	if strErr != nil {
		log.Error("GetTicketId", strErr)
		return ""
	}
	err := cli.Do(radix.FlatCmd(nil, "set", ticketId, str))
	if err != nil {
		log.Println("GetTicketId", err)
	}
	return ticketId
}

func (p *Player) Save(params ...interface{}) {
	cli := redis.GetClient("match")
	var err error
	if len(params) > 0 {
		err = cli.Do(radix.FlatCmd(nil, "HMSET", p.GId, params...))
	} else {
		err = cli.Do(radix.FlatCmd(nil, "HMSET", p.GId,
			"UserId", p.UserId,
			"IsOnline", p.IsOnline,
			"Level", p.Level,
			"Class", p.Class,
			"Sex", p.Sex,
			"State", p.State,
		))
	}
	if err != nil {
		log.Println("Player.save error", err)
	}
}
