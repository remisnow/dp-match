package service

import (
	"match/lib/utils/messageHandler"
	"match/protocol"
)

func Initial(container *messageHandler.MessageContainer) {
	//client-match
	protocol.InitMatchReqHandler(container, OnMatch)

	//redis-match
	protocol.InitUpdateRoomServiceHandler(container, OnUpdateRoomService)
	protocol.InitUserEnterRoomHandler(container, OnUserEnterRoom)
	protocol.InitUserLeaveRoomHandler(container, OnUserLeaveRoom)
	protocol.InitCreateRoomBackHandler(container, OnCreateRoomBack)
	protocol.InitUpdateRoomHandler(container, OnUpdateRoom)
}
