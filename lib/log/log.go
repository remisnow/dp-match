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
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"match/lib/file"
	"os"
	"strings"
	"time"
)

var loggerMap map[string]*Logger

const (
	//调用函数起始查找层级
	initCallerDepth = 8
)

/**
 * @Author: yaoqiang
 * @Description: 初始化
 * @Date: 2021/10/22 上午11:25
 * @Param:
 * @return:
 **/
func Init(filename string) error {
	config := &Config{}
	err := file.LoadJsonToObject(filename, config)
	if err != nil {
		logrus.WithField("filename", filename).Error(err)
		return err
	}

	return initWithConfig(config)
}

/**
 * @Author: yaoqiang
 * @Description: 初始化
 * @Date: 2022/1/8 下午2:20
 * @Param:
 * @return:
 **/
func InitWithBizTag(filename, bizTag string) error {
	config := &Config{}
	err := file.LoadJsonToObject(filename, config)
	if err != nil {
		logrus.WithField("filename", filename).Error(err)
		return err
	}

	for _, loggerConfig := range config.Loggers {
		if strings.HasSuffix(loggerConfig.OutFile, "/") {
			loggerConfig.OutFile = fmt.Sprintf("%s%s.log", loggerConfig.OutFile, bizTag)
		} else {
			loggerConfig.OutFile = fmt.Sprintf("%s/%s.log", loggerConfig.OutFile, bizTag)
		}
	}

	return initWithConfig(config)
}

func initWithConfig(config *Config) error {
	loggerMap = make(map[string]*Logger)
	for name, loggerConfig := range config.Loggers {
		logger, err := NewLogger(loggerConfig)
		if err != nil {
			logrus.WithField("LoggerName", name).Error(err)
			return err
		}
		loggerMap[name] = logger
		if config.DefaultLogger == name {
			loggerMap["default"] = logger
		}
	}
	if _, ok := loggerMap["default"]; !ok {
		//如果没有默认logger则设置为logrus的标准logger
		loggerMap["default"] = &Logger{src: logrus.StandardLogger()}
	}

	return nil
}

/**
 * @Author: yaoqiang
 * @Description: 创建logger
 * @Date: 2021/10/22 上午11:22
 * @Param:
 * @return:
 **/
func NewLogger(config *LoggerConfig) (*Logger, error) {
	logger := logrus.New()

	logFile, err := rotatelogs.New(
		config.OutFile+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(config.OutFile),
		rotatelogs.WithRotationSize(config.RotationSize),
		rotatelogs.WithRotationTime(time.Duration(config.RotationTime)*time.Minute),
		rotatelogs.WithMaxAge(time.Duration(config.MaxAge)*time.Minute),
	)
	if err != nil {
		return nil, err
	}

	//同步标准输出
	var output io.Writer
	if config.Stdout {
		output = io.MultiWriter(logFile, os.Stdout)
	} else {
		output = logFile
	}

	//输出
	logger.SetOutput(output)
	//等级
	logger.SetLevel(logrus.Level(config.Level))
	//调用函数名
	//显示调用函数
	logger.SetReportCaller(config.ReportCaller)
	//格式化
	if config.IsJsonFormatter {
		logger.SetFormatter(&JSONFormatter{})
	} else {
		logger.SetFormatter(&TextFormatter{})
	}

	return &Logger{src: logger}, nil
}

/**
 * @Author: yaoqiang
 * @Description: 通过name获取logger, 如果没找到则返回default logger
 * @Date: 2021/10/23 下午7:23
 * @Param:
 * @return:
 **/
func WithLogger(loggerName string) *Logger {
	logger, ok := loggerMap[loggerName]
	if !ok {
		return loggerMap["default"]
	}
	return logger
}

// 想日志条目添加一个字段,
// 请注意, 在调用Debug, Print, Info, Warn, Error, Fatal或Panic之前, 它不会记录日志. 它只会创建一个日志条目.
// 如果需要多个字段, 请使用'WithFields'.
func WithField(key string, value interface{}) *Entry {
	return WithLogger("default").WithField(key, value)
}

// 想日志条目添加字段的结构. 它所做的就是为每个Field调用WithField.
func WithFields(fields Fields) *Entry {
	return WithLogger("default").WithFields(fields)
}

// 将一个错误作为单个字段添加到日志条目中. 它所做的就是为给定的'error'调用'WithError'.
func WithError(err error) *Entry {
	return WithLogger("default").WithError(err)
}

// 添加Context到Entry
func WithContext(ctx context.Context) *Entry {
	return WithLogger("default").WithContext(ctx)
}

// 覆盖日志条目的时间.
func WithTime(t time.Time) *Entry {
	return WithLogger("default").WithTime(t)
}

func Logf(level Level, format string, args ...interface{}) {
	WithLogger("default").src.Logf(logrus.Level(level), format, args...)
}

func Tracef(format string, args ...interface{}) {
	WithLogger("default").src.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	WithLogger("default").src.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	WithLogger("default").src.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	WithLogger("default").src.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	WithLogger("default").src.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	WithLogger("default").src.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	WithLogger("default").src.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	WithLogger("default").src.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	WithLogger("default").src.Panicf(format, args...)
}

func Log(level Level, args ...interface{}) {
	WithLogger("default").src.Log(logrus.Level(level), args...)
}

func Trace(args ...interface{}) {
	WithLogger("default").src.Trace(args...)
}

func Debug(args ...interface{}) {
	WithLogger("default").src.Debug(args...)
}

func Info(args ...interface{}) {
	WithLogger("default").src.Info(args...)
}

func Print(args ...interface{}) {
	WithLogger("default").src.Print(args...)
}

func Warn(args ...interface{}) {
	WithLogger("default").src.Warn(args...)
}

func Warning(args ...interface{}) {
	WithLogger("default").src.Warning(args...)
}

func Error(args ...interface{}) {
	WithLogger("default").src.Error(args...)
}

func Fatal(args ...interface{}) {
	WithLogger("default").src.Fatal(args...)
}

func Panic(args ...interface{}) {
	WithLogger("default").src.Panic(args...)
}

func Logln(level Level, args ...interface{}) {
	WithLogger("default").src.Logln(logrus.Level(level), args...)
}

func Traceln(args ...interface{}) {
	WithLogger("default").src.Traceln(args...)
}

func Debugln(args ...interface{}) {
	WithLogger("default").src.Debugln(args...)
}

func Infoln(args ...interface{}) {
	WithLogger("default").src.Infoln(args...)
}

func Println(args ...interface{}) {
	WithLogger("default").src.Println(args...)
}

func Warnln(args ...interface{}) {
	WithLogger("default").src.Warnln(args...)
}

func Warningln(args ...interface{}) {
	WithLogger("default").src.Warningln(args...)
}

func Errorln(args ...interface{}) {
	WithLogger("default").src.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	WithLogger("default").src.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	WithLogger("default").src.Panicln(args...)
}
