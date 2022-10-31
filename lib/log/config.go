/*
*

	@author: yaoqiang
	@date: 2021/10/22
	@note:

*
*/
package log

type Config struct {
	Loggers       map[string]*LoggerConfig
	DefaultLogger string
}

type LoggerConfig struct {
	//日志输出文件
	OutFile string
	//日志等级
	//0 - PanicLevel - 最高严重等级, 记录日志后调用panic, 然后传递给Debug, Info, ...
	//1 - FatalLevel - 记录日志后调用'logger.Exit(1)'. 即使几日志等级设置为panic也会退出.
	//2 - ErrorLevel - 用于应该明确支出的错误. 通常用于hooks将错误发送到其它日志收集服务.
	//3 - WarnLevel  - 值得关注的非关键条目
	//5 - InfoLevel  - 用于应用程序内部运行的一般操作
	//6 - DebugLevel - 通常只在调试时调用. 非常详细的日志记录.
	//7 - TraceLevel - 比Debug更细粒度的信息事件.
	Level int //日志等级[0:Panic, 1:Fatal, 2:Error, 3:Warn, 4:Info, 5:Debug, 6:Trace]
	//记录调用的函数名及文件行数(不可用)
	ReportCaller bool
	//日志分割时间(单位:分钟)
	RotationTime int64
	//日志分割大小(字节)
	RotationSize int64
	//日志最大保留时间(单位:分钟)
	MaxAge int64
	//同步到标准输出
	Stdout bool
	//是否按json格式序列化
	IsJsonFormatter bool
}
