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

type Entry struct {
	src *logrus.Entry
}

// Returns the string representation from the reader and ultimately the
// formatter.
func (entry *Entry) String() (string, error) {
	return entry.src.String()
}

// Add an error as single field (using the key defined in ErrorKey) to the Entry.
func (entry *Entry) WithError(err error) *Entry {
	return &Entry{src: entry.src.WithError(err)}
}

// Add a context to the Entry.
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	return &Entry{src: entry.src.WithContext(ctx)}
}

// Add a single field to the Entry.
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return &Entry{src: entry.src.WithField(key, value)}
}

// Add a map of fields to the Entry.
func (entry *Entry) WithFields(fields Fields) *Entry {
	return &Entry{src: entry.src.WithFields(logrus.Fields(fields))}
}

// Overrides the time of the Entry.
func (entry *Entry) WithTime(t time.Time) *Entry {
	return &Entry{src: entry.src.WithTime(t)}
}

func (entry Entry) HasCaller() (has bool) {
	return entry.src.HasCaller()
}

func (entry *Entry) Log(level Level, args ...interface{}) {
	entry.src.Log(logrus.Level(level), args...)
}

func (entry *Entry) Trace(args ...interface{}) {
	entry.src.Trace(args...)
}

func (entry *Entry) Debug(args ...interface{}) {
	entry.src.Debug(args...)
}

func (entry *Entry) Print(args ...interface{}) {
	entry.src.Print(args...)
}

func (entry *Entry) Info(args ...interface{}) {
	entry.src.Info(args...)
}

func (entry *Entry) Warn(args ...interface{}) {
	entry.src.Warn(args...)
}

func (entry *Entry) Warning(args ...interface{}) {
	entry.src.Warning(args...)
}

func (entry *Entry) Error(args ...interface{}) {
	entry.src.Error(args...)
}

func (entry *Entry) Fatal(args ...interface{}) {
	entry.src.Fatal(args...)
}

func (entry *Entry) Panic(args ...interface{}) {
	entry.src.Panic(args...)
}

// Entry Printf family functions
func (entry *Entry) Logf(level Level, format string, args ...interface{}) {
	entry.src.Logf(logrus.Level(level), format, args...)
}

func (entry *Entry) Tracef(format string, args ...interface{}) {
	entry.src.Tracef(format, args...)
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	entry.src.Debugf(format, args...)
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.src.Infof(format, args...)
}

func (entry *Entry) Printf(format string, args ...interface{}) {
	entry.src.Printf(format, args...)
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.src.Warnf(format, args...)
}

func (entry *Entry) Warningf(format string, args ...interface{}) {
	entry.src.Warningf(format, args...)
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	entry.src.Errorf(format, args...)
}

func (entry *Entry) Fatalf(format string, args ...interface{}) {
	entry.src.Fatalf(format, args...)
}

func (entry *Entry) Panicf(format string, args ...interface{}) {
	entry.src.Panicf(format, args...)
}

// Entry Println family functions
func (entry *Entry) Logln(level Level, args ...interface{}) {
	entry.src.Logln(logrus.Level(level), args...)
}

func (entry *Entry) Traceln(args ...interface{}) {
	entry.src.Traceln(args...)
}

func (entry *Entry) Debugln(args ...interface{}) {
	entry.src.Debugln(args...)
}

func (entry *Entry) Infoln(args ...interface{}) {
	entry.src.Infoln(args...)
}

func (entry *Entry) Println(args ...interface{}) {
	entry.src.Println(args...)
}

func (entry *Entry) Warnln(args ...interface{}) {
	entry.src.Warnln(args...)
}

func (entry *Entry) Warningln(args ...interface{}) {
	entry.src.Warningln(args...)
}

func (entry *Entry) Errorln(args ...interface{}) {
	entry.src.Errorln(args...)
}

func (entry *Entry) Fatalln(args ...interface{}) {
	entry.src.Fatalln(args...)
}

func (entry *Entry) Panicln(args ...interface{}) {
	entry.src.Panicln(args...)
}
