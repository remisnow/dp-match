package base

import (
	"hash/crc32"
	"match/lib/log"
	"match/lib/utils"
	"strconv"
)

type TopicObj interface {
	GetTopicId() uint32
}

type BaseObj struct {
	_gId     string //默认生成保证单进程内唯一
	_topicId uint32
}

const _baseNamePer = "BaseObj"

var _getIntValue = utils.GetIntEvaluator()

func (o *BaseObj) InitBaseObj(GId string) {
	index := strconv.Itoa(_getIntValue())
	if GId == "" {
		o._gId = _baseNamePer + index
	} else {
		o._gId = GId + "-" + index
	}
	o._topicId = crc32.ChecksumIEEE([]byte(o._gId))
	log.Println(o._topicId)
}

func (o *BaseObj) GetTopicId() uint32 {
	return o._topicId
}
