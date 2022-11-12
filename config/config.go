package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const (
	ServiceType        = "local"
	BizTag             = "match"
	GameID             = "1"
	RoomServiceSteamID = "stream.match." + GameID //roomservice - match
	MatchStreamHead    = "stream."
	MatchRedisHead     = "Match-"
	DataPath           = "/matchAPP"
	//DataPath = "D:/Users/17444/go/src/match"
)

func GetCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}

// 获取当前执行文件绝对路径（go run）
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
		return abPath
	}
	return ""
}
func Initial(dataPath string) {
	initServiceConfig(dataPath)
}
