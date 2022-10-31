package service

import (
	"match/lib/log"
	"sync"
)

type RoomContainer struct {
	rooms map[int64]*Room
	count int
	lock  sync.RWMutex //内部枷锁禁止内部嵌套调用防止死锁
}

/*
	func NewGetRoomContainer() *RoomContainer {
		rc := &RoomContainer{}
		rc.InitialRC()
		return rc
	}
*/
func (c *RoomContainer) InitialRC() {
	c.rooms = make(map[int64]*Room)
	c.count = 0
}

func (c *RoomContainer) SetRoom(r *Room) *Room {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.rooms[r.RoomId] != nil {
		log.Error("setRoom " + r.GId + "multiple registration")
		return r
	}
	c.rooms[r.RoomId] = r
	c.count++
	return r
}

func (c *RoomContainer) RemoveRoom(r *Room) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.rooms[r.RoomId] == nil {
		log.Error("removeRoom " + r.GId + " is remove error not in container")
		return
	}
	delete(c.rooms, r.RoomId)
	c.count--
}

func (c *RoomContainer) GetRoom(roomId int64) *Room {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.rooms[roomId] == nil {
		log.Error("getRoom ", roomId, " is not in container")
	}
	return c.rooms[roomId]
}

func (c *RoomContainer) GetRoomList(list *[]*Room, fn func(r *Room) bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	for _, r := range c.rooms {
		if fn(r) {
			*list = append(*list, r)
		}
	}
}

func (c *RoomContainer) Count() int {
	return c.count
}
