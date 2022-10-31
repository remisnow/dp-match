package service

import (
	"github.com/mediocregopher/radix/v3"
	"match/config"
	"match/lib/log"
	"match/lib/redis"
	"sync"
)

const (
	RssReady = 0
	RssError = 1
)

type RoomService struct {
	lock        sync.RWMutex
	serviceName string
	state       int
	address     string
	StreamID    string
	baseData    string
	RoomContainer
}

func CreteRoomService(name string, data string) *RoomService {
	r := &RoomService{}
	r.initialRS(name, data)
	r.Save()
	log.Println("CreteRoomService ", name, data)
	return r
}

func (r *RoomService) initialRS(name string, data string) {
	r.InitialRC()
	initData := &roomServiceInitData{}
	err := Json2Obj(data, initData)
	if err != nil {
		log.Println(err, name, data)
	}
	log.Println("initialRS")
	r.serviceName = name
	r.address = initData.Address
	r.StreamID = config.MatchStreamHead + name
	r.state = RssReady
	r.baseData = data
}

func (r *RoomService) GetAddress() string {
	return r.address
}

func (r *RoomService) SetRoomServiceError() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.state = RssError
	r.Save()
}

func (r *RoomService) CanCreateRoom() bool {
	return r.state == RssReady
}

func (r *RoomService) UpdateRoomService(data string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	updateData := &roomServiceUpdate{}
	err := Json2Obj(data, updateData)
	if err != nil {
		log.Println(err, r.serviceName, data)
	}
	//更新相关的逻辑
	r.Save()
}

func (r *RoomService) Save(params ...interface{}) {
	cli := redis.GetClient("match")
	var err error
	if len(params) > 0 {
		err = cli.Do(radix.FlatCmd(nil, "HMSET", config.MatchRedisHead+"RoomService-"+r.serviceName, params...))
	} else {
		err = cli.Do(radix.FlatCmd(nil, "HMSET", config.MatchRedisHead+"RoomService-"+r.serviceName,
			"serviceName", r.serviceName,
			"state", r.state,
			"address", r.address,
			"StreamID", r.StreamID,
			"baseData", r.baseData,
		))
	}
	if err != nil {
		log.Println("RoomService.save error", err)
	}
}
