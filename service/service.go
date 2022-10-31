package service

import (
	"github.com/mediocregopher/radix/v3"
	"match/config"
	"match/lib/log"
	"match/lib/redis"
	"match/lib/utils/messageHandler"
	"strconv"
	"strings"
	"time"
)

var MatchChan chan func()
var TimerChan chan func()
var RoomServicePush chan []radix.StreamEntry

func InitLister() {
	MatchChan = make(chan func())
	TimerChan = make(chan func())
	RoomServicePush = make(chan []radix.StreamEntry, 100)
	redis.CreateConsumerGroup("match", config.RoomServiceSteamID, "matchService", "0")
	curStream := map[string]*radix.StreamEntryID{config.RoomServiceSteamID: GetReadStreamKeyID(config.RoomServiceSteamID)}
	curOpts := redis.CreateStreamReaderOpts(curStream, "matchService", "MatchConsumer", true, redis.StreamConsumeMaxCount)
	redis.CreateStreamBlockByGroupConsumer("match", curOpts, RoomServicePush)
	go func() {
		for {
			select {
			case msg := <-RoomServicePush:

				for m := 0; m < len(msg); m++ {
					//id := msg[m].ID
					for key, value := range msg[m].Fields {
						mId, _ := strconv.Atoi(key)
						messageHandler.GetMessageContainer().RedisCall(int32(mId), []byte(value))
					}
					//SaveReadStreamKeyID("stream.match."+config.GameID, id)
				}
				//room.GetRoomManager().CallEvent("createFinish", pa)
			case MatchFun := <-MatchChan:
				log.Printf("controller.MatchChan %q", MatchFun)
				MatchFun()

			case TimerFun := <-TimerChan:
				log.Printf("TimerFun %q", TimerFun)
				TimerFun()
			default:
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()
}

func SaveReadStreamKeyID(streamKey string, ID radix.StreamEntryID) error {
	cli := redis.GetClient("match")
	err := cli.Do(radix.FlatCmd(nil, "set", "match_save_"+streamKey, ID.String()))
	if err != nil {
		log.Error("SaveReadStreamKeyID", err)
	}
	return err
}

func GetReadStreamKeyID(streamKey string) *radix.StreamEntryID {
	cli := redis.GetClient("match")
	var idStr string
	err := cli.Do(radix.FlatCmd(&idStr, "get", "match_save_"+streamKey))
	if err != nil {
		log.Error("GetReadStreamKeyID", err)
	}
	if idStr == "" {
		return &radix.StreamEntryID{}
	}
	ss := strings.Split(idStr, "-")
	time, err1 := strconv.ParseUint(ss[0], 10, 64)
	seq, err2 := strconv.ParseUint(ss[1], 10, 64)
	if err1 != nil || err2 != nil {
		log.Error("GetReadStreamKeyID  err1= ", err1, "err1 = ", err2)
		return nil
	}
	return &radix.StreamEntryID{
		Time: time,
		Seq:  seq,
	}
}

// stream.serviceName
// stream.match.gameid
