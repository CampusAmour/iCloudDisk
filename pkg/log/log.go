package log

import (
	"os"

	"github.com/sirupsen/logrus"

	"iCloudDisk/pkg/config"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	FatalLevel = "fatal"
)

// InitConsoleLog 初始化console日志 在读取配置之前使用
func InitConsoleLog(printCallStack ...bool) {
	printCall := true
	if len(printCallStack) != 0 {
		printCall = printCallStack[0]
	}

	logrus.SetFormatter(newLogFmter(printCall))
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

// InitLog 初始化日志模块
func InitLog(cfg *config.LogConfig) {
	switch cfg.Level {
	case DebugLevel:
		logrus.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		logrus.SetLevel(logrus.InfoLevel)
	case WarnLevel:
		logrus.SetLevel(logrus.WarnLevel)
	case ErrorLevel:
		logrus.SetLevel(logrus.ErrorLevel)
	case FatalLevel:
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.StandardLogger().SetNoLock()

	if cfg.FileName == "console" {
		InitConsoleLog(cfg.CallStack)
		return
	}
	logrus.SetOutput(newLogFileWriter(cfg))
}

func Debug(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// Close 退出logrus
func Close() {
	logrus.Exit(0)
}
