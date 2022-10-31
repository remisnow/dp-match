package main

import (
	"github.com/astaxie/beego"
	"match/config"
	"match/lib/log"
	"match/lib/redis"
	"match/lib/utils/messageHandler"
	"match/service"
	"match/service/logic"
	"match/service/router"
)

//type player struct {
//	base.BaseObj
//}

func main() {
	//初始化log模块
	err := log.InitWithBizTag(config.DataPath+"/conf/log.json", config.BizTag)
	if err != nil {
		panic(err)
	}
	log.Println("start server")

	//初始化配置数据
	config.Initial(config.DataPath + "/conf/service.json")

	//初始化redis模块
	redis.Initial(config.DataPath + "/conf/redis.json")

	//初始化选择逻辑
	logic.InitSelectServiceLogic()

	//初始化匹配逻辑
	logic.InitRoomMatchLogic()

	//初始化房间服务管理对象
	service.CreateRoomServiceMgr()
	//curRoomService = roomService.GetRoomServiceMgr().

	//初始化房间管理对象
	service.CreateRoomManager()

	//初始化玩家管理对象
	service.CreatePlayerManager()

	//初始化MessageContainer
	messageHandler.InitMessageContainer()

	//初始化消息handler
	service.Initial(messageHandler.GetMessageContainer())

	//初始化消息接收进程
	service.InitLister()

	router.Init()
	beego.Run()
	log.Println("server exit")
}
