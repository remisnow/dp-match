package service

import (
	"match/lib/log"
)

type PlayerContainer struct {
	players map[int64]*Player
	count   int
}

func (pc *PlayerContainer) InitialPC() {
	pc.players = make(map[int64]*Player)
	pc.count = 0
}

func (pc *PlayerContainer) AddPlayer(p *Player) {
	if pc.players[p.UserId] != nil {
		log.Error("addPlayer id=", p.UserId, " multiple registration  ", pc.players[p.UserId])
		return
	}
	pc.players[p.UserId] = p
	pc.count++
}

func (pc *PlayerContainer) RemovePlayer(p *Player) {
	if pc.players[p.UserId] == nil {
		log.Error("removeRoom ", p.UserId, " is remove error not in container")
		return
	}
	delete(pc.players, p.UserId)
	pc.count--
}

func (pc *PlayerContainer) GetPlayer(UserId int64) *Player {
	return pc.players[UserId]
}

func (pc *PlayerContainer) GetPlayerTotal() int {
	return pc.count
}
