package global

// 全局日志，在Boot初始化
var Log oaLog

type oaLog interface {
	Info(content string)
	Error(content string)
	Debug(content string)
}
