/*
*

	@author: yaoqiang
	@date: 2021/10/22
	@note:

*
*/
package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

// type Fields map[string]interface{}
type Fields logrus.Fields

//type myFields struct {
//	src logrus.Fields
//}

type Level logrus.Level

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type Logger struct {
	src *logrus.Logger
}

// 想日志条目添加一个字段,
// 请注意, 在调用Debug, Print, Info, Warn, Error, Fatal或Panic之前, 它不会记录日志. 它只会创建一个日志条目.
// 如果需要多个字段, 请使用'WithFields'.
func (l *Logger) WithField(key string, value interface{}) *Entry {
	return &Entry{src: l.src.WithField(key, value)}
}

// 想日志条目添加字段的结构. 它所做的就是为每个Field调用WithField.
func (l *Logger) WithFields(fields Fields) *Entry {
	return &Entry{src: l.src.WithFields(logrus.Fields(fields))}
}

// 将一个错误作为单个字段添加到日志条目中. 它所做的就是为给定的'error'调用'WithError'.
func (l *Logger) WithError(err error) *Entry {
	return &Entry{src: l.src.WithError(err)}
}

// 添加Context到Entry
func (l *Logger) WithContext(ctx context.Context) *Entry {
	return &Entry{src: l.src.WithContext(ctx)}
}

// 覆盖日志条目的时间.
func (l *Logger) WithTime(t time.Time) *Entry {
	return &Entry{src: l.src.WithTime(t)}
}

func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	l.src.Logf(logrus.Level(level), format, args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.src.Tracef(format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.src.Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.src.Infof(format, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.src.Printf(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.src.Warnf(format, args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.src.Warningf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.src.Errorf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.src.Fatalf(format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.src.Panicf(format, args...)
}

func (l *Logger) Log(level Level, args ...interface{}) {
	l.src.Log(logrus.Level(level), args...)
}

func (l *Logger) Trace(args ...interface{}) {
	l.src.Trace(args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.src.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.src.Info(args...)
}

func (l *Logger) Print(args ...interface{}) {
	l.src.Print(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.src.Warn(args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.src.Warning(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.src.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.src.Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.src.Panic(args...)
}

func (l *Logger) Logln(level Level, args ...interface{}) {
	l.src.Logln(logrus.Level(level), args...)
}

func (l *Logger) Traceln(args ...interface{}) {
	l.src.Traceln(args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.src.Debugln(args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.src.Infoln(args...)
}

func (l *Logger) Println(args ...interface{}) {
	l.src.Println(args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.src.Warnln(args...)
}

func (l *Logger) Warningln(args ...interface{}) {
	l.src.Warningln(args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.src.Errorln(args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.src.Fatalln(args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.src.Panicln(args...)
}
