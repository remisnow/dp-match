package file

import (
	"encoding/json"
	"unsafe"
)

/**
 * @Author: yaoqiang
 * @Description: 读取json文件
 * @Date: 2020/8/28 5:09 下午
 * @param null
 * @return:
 **/
func LoadJsonToObject(filename string, t interface{}) error {

	buf, e := loadFile(filename)

	if buf == nil {
		return e
	}

	if e != nil {
		return e
	}

	err := json.Unmarshal(buf, &t)

	if err != nil {
		return err
	}

	return nil
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
