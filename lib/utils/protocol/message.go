/*
*

	@author: yaoqiang
	@date: 2021/11/9
	@note:

*
*/
package protocol

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/golang/protobuf/proto"
	"hash/crc32"
	"match/lib/log"
	"match/lib/result"
	"reflect"
	"time"
)

/**
 * @Author: yaoqiang
 * @Description: 根据消息名称生成消息id
 * @Date: 2021/11/10 下午2:28
 * @Param:
 * @return:
 **/
func MessageId1(msgName string) int32 {
	return int32(crc32.ChecksumIEEE([]byte(msgName)))
}

/**
 * @Author: yaoqiang
 * @Description: 消息id
 * @Date: 2021/11/10 下午2:28
 * @Param:
 * @return:
 **/
func MessageId(msg proto.Message) int32 {
	name := reflect.TypeOf(msg).Elem().Name()
	log.Println("MessageId ", name)

	return MessageId1(name)
}

func parseToucanRequest(ctx *context.Context) (req *ToucanRequest, rsp *ToucanResponse) {
	rsp = &ToucanResponse{RetCode: result.Success, RetMsg: result.SuccessMsg}
	log.Println("parseToucanRequest ")
	//if !ctx.Input.IsPost() {
	//	rsp.RetCode = result.Method
	//	rsp.RetMsg = "it must be a post method"
	//	return
	//}
	b := ctx.Input.RequestBody
	ctx.Request.Body.Close()

	req = &ToucanRequest{}
	err := proto.Unmarshal(b, req)
	if err != nil {
		log.Error(err)
		rsp.RetCode = result.Decoding
		rsp.RetMsg = "deserialize ToucanRequest failed"
		return nil, rsp
	}
	if req.CliInfo == nil {
		log.Error("input CliInfo is nil")
		rsp.RetCode = result.Param
		rsp.RetMsg = "input CliInfo is nil"
		return nil, rsp
	}

	rsp.CliTime = req.CliTime

	if int(req.DataLen) != len(req.Data) {
		log.Error("input DataLen not equal to length of Data")
		rsp.RetCode = result.Param
		rsp.RetMsg = "input DataLen not equal to length of Data"
		return nil, rsp
	}

	code, msg := req.CliInfo.IsValid()
	if code != result.Success {
		log.WithField("Error", msg).Error("ClientInfo invalid")
		rsp.RetCode = int32(code)
		rsp.RetMsg = msg
		return nil, rsp
	}

	log.WithFields(map[string]interface{}{
		"ClientInfo": req.CliInfo.String(),
		"Token":      req.Token,
		"CliTime":    req.CliTime,
		"MessageId":  req.MessageId,
		"DataLen":    req.DataLen,
	}).Info("message header")
	return req, rsp
}

/**
 * @Author: yaoqiang
 * @Description: 解析消息
 * @Date: 2021/11/9 下午4:04
 * @Param:
 * @return:
 **/
func ParseMessage(bs []byte, req proto.Message) (code int32, msg string) {
	err := proto.Unmarshal(bs, req)
	if err != nil {
		log.Error(err)
		return result.Decoding, fmt.Sprintf("deserialize %s failed", proto.MessageName(req))
	}

	log.WithField("MessageName", proto.MessageName(req)).Info(req.String())
	return result.Success, result.SuccessMsg
}

/**
 * @Author: yaoqiang
 * @Description: 消息回复
 * @Date: 2021/11/9 下午4:04
 * @Param:
 * @return:
 **/
func ResponseMessage(ctx *context.Context, rsp *ToucanResponse) {
	if rsp == nil {
		return
	}
	rsp.SvrTime = time.Now().UnixNano() / int64(time.Millisecond)

	b, err := proto.Marshal(rsp)
	if err != nil {
		log.WithField("Error", err).Error("ResponseMessage.proto.Marshal ToucanResponse failed")
		return
	}

	ctx.ResponseWriter.Write(b)
	log.WithField("Message", rsp.String()).Info("send message success")
}

/**
 * @Author: yaoqiang
 * @Description: 添加消息body
 * @Date: 2021/11/10 下午8:15
 * @Param:
 * @return:
 **/
func (r *ToucanResponse) AddMessage(m proto.Message) {
	if m == nil {
		return
	}
	if r.Bodies == nil {
		r.Bodies = make([]*ToucanBody, 0)
	}
	data, err := proto.Marshal(m)
	if err != nil {
		log.WithField("MessageName", proto.MessageName(m)).WithField("Error", err.Error()).Error("proto.Marshal failed")
		return
	}
	r.Bodies = append(r.Bodies, &ToucanBody{
		MessageId: MessageId(m),
		DataLen:   int32(len(data)),
		Data:      data,
	})
}

func GetParams(ctx *context.Context) (account string, req *ToucanRequest, rsp *ToucanResponse, code int32, msg string) {
	account = ctx.Input.Param("Account")
	in := ctx.Input.GetData("ToucanRequest")
	if in == nil {
		log.Error("input ToucanRequest is nil")
		return account, nil, nil, result.Unknown, "data ToucanRequest is nil"
	}
	out := ctx.Input.GetData("ToucanResponse")
	if out == nil {
		log.Error("input ToucanResponse is nil")
		return account, nil, nil, result.Unknown, "data ToucanResponse is nil"
	}
	var ok bool
	req, ok = in.(*ToucanRequest)
	if !ok {
		log.Error("data request is not ToucanRequest")
		return account, nil, nil, result.Unknown, "data request is not ToucanRequest"
	}
	rsp, ok = out.(*ToucanResponse)
	if !ok {
		log.Error("data request is not ToucanResponse")
		return account, req, nil, result.Unknown, "data request is not ToucanResponse"
	}
	return account, req, rsp, result.Success, result.SuccessMsg
}
