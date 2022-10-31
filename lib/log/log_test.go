/*
*

	@author: yaoqiang
	@date: 2021/10/22
	@note:

*
*/
package log

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestLogConfig(t *testing.T) {
	config := &Config{
		Loggers:       make(map[string]*LoggerConfig),
		DefaultLogger: "test1",
	}
	loggerConfig1 := &LoggerConfig{
		OutFile: "./logfile/test1.log",
		Level:   6,
	}
	loggerConfig2 := &LoggerConfig{
		OutFile: "./logfile/test2.log",
		Level:   6,
	}
	config.Loggers["test1"] = loggerConfig1
	config.Loggers["test2"] = loggerConfig2

	b, _ := json.Marshal(config)
	fmt.Println(string(b))
}

func TestLog(t *testing.T) {
	err := Init("log.json")
	if err != nil {
		panic(err)
	}

	Debug("say 1")
	Debugf("say 2")
	Debugln("say 3")
	WithLogger("test1").Debug("say 4")
	WithLogger("test1").Debugf("say 5")
	WithLogger("test1").Debugln("say 6")

	for i := 7; true; i++ {
		WithFields(
			Fields{
				"field1": "abc",
				"field2": 123,
			}).Debug("say", i)
		time.Sleep(time.Millisecond * 10)
	}
}
