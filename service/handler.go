package service

import (
	"github.com/astaxie/beego/context"
	"match/config"
	"match/lib/log"
	"match/lib/result"
	"match/lib/utils/messageHandler"
	"match/protocol"
	"math/rand"
	"strconv"
)

func OnMatch(ctx *context.Context, req *protocol.MatchReq) *protocol.MatchResp {
	resp := &protocol.MatchResp{}

	userIdStr := ctx.Input.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		resp.RetCode = result.Internal
		return resp
	}
	//gameId, err1 := strconv.Atoi(ctx.Input.Param("gameId"))

	gameId := 1 //在input里面取 或者在 配置文件取
	gameMode := int(req.GameMode)

	//检查玩家是否存在 内存 redis
	p := GetPlayerManager().GetPlayer(userId)
	if p == nil {
		// 不存在创建玩家
		p = GetPlayerManager().CreatePlayer(userId)
	}
	if p.GetState() != PlayerStateOutRoom {
		resp.RetCode = result.Repeatedly
		return resp
	}

	res := make(chan *protocol.MatchResp)
	defer close(res)
	MatchChan <- func() {

		curRoom := MatchRoom(p, gameId, gameMode)
		//找到房间
		if curRoom != nil {
			curRoom.UserEnterRoom(p)
			curRand := rand.Int31()
			res <- &protocol.MatchResp{
				RetCode:  result.Success,
				Address:  curRoom.GetService().GetAddress(),
				TicketId: p.GetTicketId(curRand),
				Rand:     curRand}
			return
		}

		//创建一个房间
		curRoomService := SelectRoomService(gameId, gameMode)
		log.Println("curRoomService =", curRoomService)
		curRoom = GetRoomManager().CreateRoom(gameMode, curRoomService)
		//设置玩家等待进入房间的状态 房间过期时间
		log.Println(p.GetTopicId())
		curRoom.WatchEvent(p.GetTopicId(), "createFinish", func(i interface{}) {
			log.Println("事件createFinish的 回调")
			GetRoomManager().UnWatchEvent(curRoom.GetTopicId(), "createFinish")
			par := i.(*protocol.CreateRoomBack) // 取出接口中的数据
			if par.Error > 0 {
				//清理玩家的状态
				res <- &protocol.MatchResp{
					RetCode:  result.Unknown,
					Address:  "",
					TicketId: ""}
				return
			}

			curRoom.UserEnterRoom(p)
			curRand := rand.Int31()

			res <- &protocol.MatchResp{
				RetCode:  result.Success,
				Address:  curRoom.GetService().GetAddress(),
				TicketId: p.GetTicketId(curRand),
				Rand:     curRand}

		})
		messageHandler.GetMessageContainer().CallRedis("match", config.MatchStreamHead+curRoomService.serviceName, &protocol.CreateRoom{
			GameId:   int32(gameId),
			GameMode: int32(gameMode),
			RoomId:   curRoom.RoomId,
			Data:     "",
		})

	}
	return <-res
}

func OnUpdateRoomService(message *protocol.UpdateRoomService) *protocol.UpdateRoomService {
	log.Println("OnUpdateRoomService", message)
	GetRoomServiceMgr().OnUpdateRoomService(message)
	return message
}

func OnUserEnterRoom(message *protocol.UserEnterRoom) *protocol.UserEnterRoom {
	if GetRoomManager().GetRoom(message.RoomId).GetPlayer(message.UserId) == nil {
		log.Println("OnUserEnterRoom error", message)
	}
	GetPlayerManager().GetPlayer(message.UserId).OnUserEnterRoom()
	return message
}

func OnUserLeaveRoom(message *protocol.UserLeaveRoom) *protocol.UserLeaveRoom {
	if GetRoomManager().GetRoom(message.RoomId).GetPlayer(message.UserId) == nil {
		log.Println("OnUserLeaveRoom error", message)
	}
	log.Println("OnUserLeaveRoom ", message)
	curRoom := GetRoomManager().GetRoom(message.RoomId)
	curRoom.RemovePlayer(curRoom.GetPlayer(message.UserId))
	GetPlayerManager().GetPlayer(message.UserId).OnUserLeaveRoom()
	return message
}

func OnCreateRoomBack(message *protocol.CreateRoomBack) *protocol.CreateRoomBack {
	if message.Error > 0 {
		//删除Match 服创建好的房间
	}
	//更新room数据
	GetRoomManager().GetRoom(message.RoomId).CallEvent("createFinish", message)
	return message
}

func OnUpdateRoom(message *protocol.UpdateRoom) *protocol.UpdateRoom {
	return message
}
