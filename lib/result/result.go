/*
*

	@author: yaoqiang
	@date: 2021/12/8
	@note:

*
*/
package result

import "fmt"

type ResultInterface interface {
	GetRetCode() int32
	GetRetMsg() string
}

type Result struct {
	RetCode int32
	RetMsg  string
}

func (r *Result) String() string {
	return fmt.Sprintf("RetCode=%d, RetMsg=%s", r.RetCode, r.RetMsg)
}

const SuccessMsg = "success"

const (
	Success    int32 = 0    //成功
	Unknown    int32 = 1001 //未知错误
	Param      int32 = 1002 //参数错误
	Encoding   int32 = 1003 //序列化失败
	Decoding   int32 = 1004 //反序列化失败
	Internal   int32 = 1005 //内部错误
	Repeatedly int32 = 1006 //重复消息
)
