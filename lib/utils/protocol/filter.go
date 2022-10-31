/*
*

	@author: yaoqiang
	@date: 2021/11/10
	@note:

*
*/
package protocol

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/golang/protobuf/proto"
	"match/lib/log"
	"match/lib/result"
)

var ignoreVerifyTokenUri = make(map[string]int)

/**
 * @Author: yaoqiang
 * @Description: 注册忽略token验证的uri
 * @Date: 2021/11/10 下午7:44
 * @Param:
 * @return:
 **/
func RegisterIgnoreVerifyToken(uri string) {
	ignoreVerifyTokenUri[uri] = 0
}

/**
 * @Author: yaoqiang
 * @Description: 初始化过滤器
 * @Date: 2021/11/10 下午7:53
 * @Param:
 * @return:
 **/
func InitFilter() {
	//tk.Initial()
	//注册消息解析过滤器
	beego.InsertFilter("/*", beego.BeforeExec, FilterClientMessage)
}

func FilterClientMessage(ctx *context.Context) {
	req, rsp := parseToucanRequest(ctx)
	ctx.Input.SetData("ToucanRequest", req)
	ctx.Input.SetData("ToucanResponse", rsp)
	if req != nil && result.Success == rsp.RetCode {
		//if _, ok := ignoreVerifyTokenUri[ctx.Request.RequestURI]; ok {
		//	//不需要验证token
		//	return
		//}
		//验证token
		token, code := VerifyToken(req.Token)
		if code == result.Success {
			ctx.Input.SetParam("userId", token)
			return
		}
		//token验证失败
		rsp.RetCode = code
		rsp.RetMsg = "token is wrong"
	}

	rsp.AddMessage(&CommonError{RetCode: rsp.RetCode, RetMsg: rsp.RetMsg})

	//解析或者验证失败
	b, _ := proto.Marshal(rsp)
	ctx.ResponseWriter.Write(b)
	log.WithField("Msg", rsp.String()).Info("send message")
}
