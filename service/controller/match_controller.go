package controller

import (
	"github.com/astaxie/beego"
	"match/lib/log"
	"match/lib/result"
	"match/lib/utils/messageHandler"
	pc "match/lib/utils/protocol"
)

type MatchController struct {
	beego.Controller
}

// MatchChan var MatchControllerChan chan *MatchController = make(chan *MatchController)
//var MatchChan = make(chan func())

func (c *MatchController) Match() {

	//log.Println("receive Match")
	//log.Println(c.Ctx.Request.RequestURI)
	//res := make(chan []byte)
	//defer close(res)
	//MatchChan <- func() {
	//	log.Println("run  func")
	//	curRoom := room.GetRoomManager().CreateRoom(1)
	//	log.Println(curRoom.GetTopicId())
	//	//defer close(res)
	//	room.GetRoomManager().WatchEvent(curRoom.GetTopicId(), "createFinish", func(i interface{}) {
	//		//time.Sleep(10 * time.Second)
	//		log.Println("事件createFinish的 回调")
	//		par := i.(map[string]interface{}) // 取出接口中的数据
	//		if data, ok := par["data"].([]byte); ok {
	//			log.Println(data)
	//			res <- data
	//		}
	//		room.GetRoomManager().UnWatchEvent(curRoom.GetTopicId(), "createFinish")
	//	})
	//	log.Println("注册事件完毕")
	//
	//	redis.PubMessage("pubSub_room", &radix.PubSubMessage{
	//		Channel: "rlp",
	//		Message: []byte(c.Ctx.Request.RequestURI),
	//	})
	//	log.Println("写入消息队列")
	//}
	//c.Ctx.ResponseWriter.Write([]byte(string(<-res) + " end"))
	//massage := &protocol.CreateRoom{GID: "appName-serviceName-index", Type: protocol.RoomType_COMMON}
	//data, err := proto.Marshal(massage)
	//if err != nil {
	//	log.Println(err)
	//}
	////log.Println(massage.String())
	//log.Println(data)
	//var receiveM protocol.CreateRoom
	//errR := proto.Unmarshal(data, &receiveM)
	//if errR != nil {
	//	log.Println(errR)
	//}
	//log.Println(receiveM)
	//c.Ctx.ResponseWriter.Write([]byte(receiveM.String()))

}

type httpRsp struct {
	MessageId int32
	DataLen   int32
	Data      []byte
	err       error
}

func (c *MatchController) Common() {
	log.Println("receive Common", c.Ctx.Request.RequestURI)
	account, tReq, tRsp, code, msg := pc.GetParams(c.Ctx)
	log.Println(account)
	if code != result.Success {
		log.WithField("Error", msg).Error("pc.GetParams failed")
		return
	}

	data := tReq.Data
	if data == nil {
		return
	}
	//res := make(chan *httpRsp)
	//defer close(res)

	//MatchChan <- func() {
	rspMessageId, rspData, dataLen, err := messageHandler.GetMessageContainer().HttpCall(c.Ctx, tReq.MessageId, data)
	//res <- &httpRsp{
	//	MessageId: rspMessageId,
	//	DataLen:   dataLen,
	//	Data:      rspData,
	//	err:       err,
	//}
	//}

	//hRsp := <-res
	hRsp := &httpRsp{
		MessageId: rspMessageId,
		DataLen:   dataLen,
		Data:      rspData,
		err:       err,
	}
	if hRsp.err != nil {
		tRsp.RetCode = result.Decoding
		tRsp.RetMsg = hRsp.err.Error()
	}
	tRsp.RetCode = result.Success
	tRsp.RetMsg = ""

	if tRsp.Bodies == nil {
		tRsp.Bodies = make([]*pc.ToucanBody, 0)
	}
	tRsp.Bodies = append(tRsp.Bodies, &pc.ToucanBody{
		MessageId: hRsp.MessageId,
		DataLen:   hRsp.DataLen,
		Data:      hRsp.Data,
	})

	pc.ResponseMessage(c.Ctx, tRsp)

}

func (c *MatchController) Leave() {
	log.Println("receive Leave")
	log.Println(c)
	c.Ctx.ResponseWriter.Write([]byte(c.Ctx.Request.RequestURI + " end"))

}
