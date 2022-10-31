package service

import (
	"match/lib/redis"
	"sync"
)

var playerMgr *playerManager

type playerManager struct {
	PlayerContainer
	lock sync.RWMutex //内部加锁方法 禁止内部嵌套调用 防止死锁
}

/*var getIndex = utils.GetIntEvaluator()*/
var getPlayerIndex = func() int {
	return redis.GetINCRValue("match", "math_player_incr_index")
}

func GetPlayerManager() *playerManager {
	return playerMgr
}

func CreatePlayerManager() {
	playerMgr = &playerManager{}
	playerMgr.initial()
}

func (pm *playerManager) initial() {
	pm.InitialPC()
}
func (pm *playerManager) CreatePlayer(userId int64) *Player {
	playerId := getPlayerIndex()
	curPlayer := &Player{}
	curPlayer.initPlayer(playerId, userId)

	pm.AddPlayer(curPlayer)
	return curPlayer
}
