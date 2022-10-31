package service

import (
	"match/lib/event"
	"match/lib/redis"
	"sync"
)

var roomMgr *roomManager

type roomManager struct {
	event.Event
	RoomContainer
	FreeRooms RoomContainer
	lock      sync.RWMutex //内部加锁方法 禁止内部嵌套调用 防止死锁
}

// var getIndex = utils.GetIntEvaluator()
var getIndex = func() int {
	return redis.GetINCRValue("match", "math_room_incr_index")
}

func GetRoomManager() *roomManager {
	return roomMgr
}

func CreateRoomManager() {
	roomMgr = &roomManager{}
	roomMgr.initial()
}

func (rm *roomManager) initial() {
	rm.InitialRC()
	rm.FreeRooms.InitialRC()
}
func (rm *roomManager) CreateRoom(roomType int, service *RoomService) *Room {
	curRoom := &Room{}
	var roomId = getIndex()
	curRoom.initRoom(roomId, roomType, service)

	rm.SetRoom(curRoom)
	rm.FreeRooms.SetRoom(curRoom)
	service.SetRoom(curRoom)
	return curRoom
}
