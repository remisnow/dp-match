package router

import (
	"github.com/astaxie/beego"
	"match/lib/utils/protocol"
	"match/service/controller"
)

func Init() {
	beego.Router("/match/*", &controller.MatchController{}, "*:Match")
	beego.Router("/leave/*", &controller.MatchController{}, "*:Leave")
	beego.Router("/common/*", &controller.MatchController{}, "*:Common")
	protocol.InitFilter()
}
