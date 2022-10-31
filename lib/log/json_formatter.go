/*
*

	@author: yaoqiang
	@date: 2021/10/23
	@note:

*
*/
package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

// JSONFormatter formats logs into parsable json
type JSONFormatter struct{}

// Format renders a single log entry
func (f *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	data[logrus.FieldKeyTime] = entry.Time.Local().Format("2006-01-02 15:04:05.000")
	if entry.HasCaller() {
		if f := getCaller(); f != nil {
			i := strings.LastIndex(f.Function, ".")
			data[logrus.FieldKeyFunc] = fmt.Sprintf("%s:%d:%s", filepath.Base(f.File), f.Line, f.Function[i+1:])
		}
	}
	data[logrus.FieldKeyLevel] = strings.ToUpper(entry.Level.String())
	data[logrus.FieldKeyMsg] = entry.Message

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}

/**
 * @Author: yaoqiang
 * @Description: 查找调用函数
 * @Date: 2021/10/23 下午8:34
 * @Param:
 * @return:
 **/
func callerPrettyfier(frame *runtime.Frame) (function string, file string) {
	prevFunc, prevFile := "", ""
	for skip := initCallerDepth; true; skip++ {
		pc, codePath, codeLine, ok := runtime.Caller(skip)
		if !ok {
			//函数栈用尽了
			file = prevFile
			function = prevFunc
			return function, file
		} else {
			prevFile = fmt.Sprintf("%s:%d", codePath, codeLine)
			prevFunc = runtime.FuncForPC(pc).Name()
			file = prevFile
			function = prevFunc
			if !strings.Contains(prevFile, "log.go") && !strings.Contains(prevFile, "logger.go") {
				//找到包外的函数了
				return function, file
			}
		}
	}
	return function, file
}
