package service

import (
	"encoding/json"
	"unsafe"
)

// 根据功能随时添加 message 中的json转成struct

// 房间服务启动的时候的详细信息
type roomServiceInitData struct {
	Address string
}

// 房间服务更新时的详细信息
type roomServiceUpdate struct {
}

// 房间数据更新时的额外信息
type roomUpdate struct {
}

func Json2Obj(str string, t interface{}) error {
	b := str2bytes(str)
	err := json.Unmarshal(b, t)
	return err
}

func Obj2Json(t interface{}) (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return bytes2str(b), err
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
