package common

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 保存logger实例的全局变量, 分别写入终端和日志文件
var LoggerTer *logrus.Logger
var LoggerFile *logrus.Logger

// InitLogger 日志记录器配置初始化
// =>
func InitLogger() {
	// 创建新的日志记录器
	loggerTer := logrus.New()
	loggerFile := logrus.New()

	// 设置日志级别
	loggerTer.SetLevel(logrus.InfoLevel)
	loggerFile.SetLevel(logrus.InfoLevel)

	// 设置输出 call site 信息
	// loggerTer.SetReportCaller(true)
	// loggerFile.SetReportCaller(true)

	// 设置输出到文件的日志配置，并配置滚动
	logFile := &lumberjack.Logger{
		Filename:   "log/DracoNisus-Thuban.log", // 日志文件路径
		MaxSize:    10,                          // 每个日志文件的最大尺寸，单位是MB
		MaxBackups: 3,                           // 保留旧日志文件的最大个数
		MaxAge:     7,                           // 保留旧日志文件的最大天数
		Compress:   true,                        // 是否启用压缩
	}

	// 设置日志同时输出到终端和文件
	loggerTer.SetOutput(io.Writer(os.Stdout))
	loggerFile.SetOutput(io.Writer(logFile))

	// 设置终端日志格式为text
	loggerTer.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})

	// 设置文件日志格式为json
	loggerFile.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
	})

	// 保存配置好的日志记录器
	LoggerTer = loggerTer
	LoggerFile = loggerFile
}
