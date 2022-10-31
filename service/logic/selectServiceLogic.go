package logic

import (
	"match/service"
)

func SelectRoomService() *service.RoomService {
	rsMgr := service.GetRoomServiceMgr()
	var list []*service.RoomService
	rsMgr.GetRoomServiceList(&list, func(rs *service.RoomService) bool {
		return rs.CanCreateRoom()
	})
	if len(list) <= 0 {
		return nil
	}
	if len(list) == 1 {
		return list[0]
	}
	// todo 根据房间数量和服务压力做负载均衡
	return list[0]

}

// 注册选择逻辑（优先考虑负债均衡）

func InitSelectServiceLogic() {
	service.RegisterSelectLogic("commonSelect", SelectRoomService)
}
