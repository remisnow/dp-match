package service

import (
	"github.com/mediocregopher/radix/v3"
	"log"
	"match/config"
	"match/lib/redis"
	"match/protocol"
	"sync"
)

var roomServiceMgr *rsManager

type rsManager struct {
	services map[string]*RoomService
	lock     sync.RWMutex //内部加锁方法 禁止内部嵌套调用 防止死锁
}

func GetRoomServiceMgr() *rsManager {
	return roomServiceMgr
}

func CreateRoomServiceMgr() {
	roomServiceMgr = &rsManager{}
	roomServiceMgr.initial()
}

func (rsm *rsManager) initial() {
	rsm.services = make(map[string]*RoomService)
	rsm.initialOldService()
}
func (rsm *rsManager) initialOldService() {
	log.Println("initialOldService-------------------------------------------")
	cli := redis.GetClient("match")
	var list []string
	err := cli.Do(radix.FlatCmd(&list, "Smembers", "Match_ServiceList"))
	if err != nil {
		log.Println("initialOldService error", err)
	}
	if len(list) > 0 {
		for _, s := range list {
			var service []string
			err1 := cli.Do(radix.FlatCmd(&service, "Hmget", s, "serviceName", "baseData"))
			if err != nil {
				log.Println("initialOldService error", err1)
			}
			if len(service) > 0 {
				rsm.AddNewRoomService(service[0], service[1])
			}
		}
	}
}
func (rsm *rsManager) AddNewRoomService(name string, data string) {
	rsm.lock.Lock()
	defer rsm.lock.Unlock()
	rsm.services[name] = CreteRoomService(name, data)
	rsm.Save(name)

}

func (rsm *rsManager) GetRoomService(name string) *RoomService {
	rsm.lock.Lock()
	defer rsm.lock.Unlock()
	if rsm.services[name] != nil {
		return rsm.services[name]
	}
	return nil
}

func (rsm *rsManager) GetRoomServiceList(list *[]*RoomService, fn func(rs *RoomService) bool) {
	rsm.lock.Lock()
	defer rsm.lock.Unlock()
	for _, r := range rsm.services {
		if fn(r) {
			*list = append(*list, r)
		}
	}
}

func (rsm *rsManager) OnUpdateRoomService(pb *protocol.UpdateRoomService) {
	switch pb.State {
	case protocol.RoomServiceState_reload:
	case protocol.RoomServiceState_load:
		rsm.AddNewRoomService(pb.ServiceName, pb.Data)
	case protocol.RoomServiceState_heart:
		rsm.GetRoomService(pb.ServiceName).UpdateRoomService(pb.Data)
	case protocol.RoomServiceState_busy:
		rsm.GetRoomService(pb.ServiceName).SetRoomServiceError()
	default:
		break
	}
}

func (rsm *rsManager) Save(serviceName string) {
	cli := redis.GetClient("match")

	err := cli.Do(radix.FlatCmd(nil, "SADD", "Match_ServiceList", config.MatchRedisHead+"RoomService-"+rsm.services[serviceName].serviceName))

	if err != nil {
		log.Println("rsManager.save error", err)
	}
}
