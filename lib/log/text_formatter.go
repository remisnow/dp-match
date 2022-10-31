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
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	//bufferPool *sync.Pool
	//
	//// qualified package name, cached at first use
	//logrusPackage string

	// Positions in the call stack when tracing to report the calling method
	minimumCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

const (
	maximumCallerDepth int = 25
	knownLogFrames     int = 8
)

type TextFormatter struct{}

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05.000")
	var file string
	var line int
	var function string
	if entry.HasCaller() {
		if f := getCaller(); f != nil {
			file = filepath.Base(f.File)
			line = f.Line
			i := strings.LastIndex(f.Function, ".")
			function = f.Function[i+1:]
		}
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(fmt.Sprintf("%s [%s:%d:%s] [%s] : %s", timestamp, file, line, function, strings.ToUpper(entry.Level.String()), entry.Message))
	if entry.Data != nil && len(entry.Data) > 0 {
		b.WriteString(" :")
		i := 0
		len := len(entry.Data)
		for k, v := range entry.Data {
			b.WriteString(fmt.Sprintf(" %s=%v", k, v))
			if i < len-1 {
				b.WriteString(",")
			}
		}
	}
	b.WriteString("\n")
	return b.Bytes(), nil
}

func getCaller() *runtime.Frame {

	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, 2)
		_ = runtime.Callers(0, pcs)
		// now that we have the cache, we can skip a minimum count of known-logrus functions
		// XXX this is dubious, the number of frames may vary
		minimumCallerDepth = knownLogFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		// If the caller isn't in log.go or logger.go, we're done
		if !strings.Contains(f.File, "log.go") && !strings.Contains(f.File, "logger.go") {
			return &f
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}
