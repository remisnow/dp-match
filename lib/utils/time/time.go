/*
*

	@author: yaoqiang
	@date: 2020/9/11
	@note:

*
*/
package utilTime

import (
	"time"
)

const (
	StandardFormat  = "2006-01-02 15:04:05"                   //时间格式
	SecondsOfWeek   = int64(time.Hour * 24 * 7 / time.Second) //一周多少秒
	SecondsOfDay    = int64(time.Hour * 24 / time.Second)     //一天多少秒
	SecondsOfHour   = int64(time.Hour / time.Second)          //一小时多少秒
	SecondsOfMinute = int64(time.Minute / time.Second)        //一分钟多少秒
)

// 当前时间戳(秒)
func GetCurrentSecond() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

// 当前时间戳(毫秒)
func GetCurrentMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 当前时间戳(微妙)
func GetCurrentMicrosecond() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// 是否同一天
func IsSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

// 同一周
func SameWeek(t1, t2 time.Time) bool {
	if t1.Unix() == t2.Unix() {
		return true
	}
	start := t1
	end := t2
	if end.Unix() < start.Unix() {
		start = t2
		end = t1
	}
	d1 := int32(start.Unix() / int64(time.Hour*24))
	d2 := int32(end.Unix() / int64(time.Hour*24))
	dbl := d2 - d1
	dow := int32(end.Weekday())
	if dow == 0 {
		dow = 7
	}
	return dbl < 7 && dbl < dow
}

// 计算给定时间的周数,1970-01-01所在周为第0周
// 1970-01-01为周4,可以自己设定从周几[0:6]的几点几分几秒开始计算周数量
func CalcWeekNumber(unixTimestamp int64, startWeek time.Weekday, hour, minute, second int64) int64 {
	totalSecond := hour*60*60 + minute*60 + second
	return (unixTimestamp + int64(time.Thursday-startWeek)*SecondsOfWeek - totalSecond) / SecondsOfWeek
}

var startDay, _ = time.ParseInLocation(StandardFormat, "2020-01-01 00:00:00", time.Local)

// 计算当前天数, 2020-01-01为第一天
func CalcCurrentDayNumber() int64 {
	return CalcDayNumber(time.Now().Unix())
}

// 计算给定时间的天数, 2020-01-01为第一天
func CalcDayNumber(t int64) int64 {
	return (t-startDay.Unix())/SecondsOfDay + 1
}

type TimeStopChan chan bool

func TimeTimer(d time.Duration, fun func()) chan bool {
	timer := time.NewTimer(d)

	stopChan := make(chan bool)
	go func(timer *time.Timer) {
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				fun()
			case stop := <-stopChan:
				if stop {
					close(stopChan)
					return
				}
			}
		}
	}(timer)

	return stopChan
}
