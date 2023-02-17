package log

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"iCloudDisk/pkg/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func newLogFileWriter(logCfg *config.LogConfig) *rotatelogs.RotateLogs {
	logFullPath := path.Join(logCfg.Dir, logCfg.FileName)
	fmt.Printf("log full path: %s\n", logFullPath)
	logier, err := rotatelogs.New(
		logFullPath+"_%Y%m%d%H%M",
		rotatelogs.WithLinkName(logFullPath),
		rotatelogs.WithRotationCount(logCfg.RotationCount),
		rotatelogs.WithRotationTime(time.Second*time.Duration(logCfg.Rotate)))
	if err != nil {
		fmt.Printf("[ERROR]new rotatelogs error: %s\n", err.Error())
		os.Exit(-1)
	}
	return logier
}

type LogFmter struct {
	callStack bool
}

func newLogFmter(printCallStack bool) *LogFmter {
	fmter := new(LogFmter)
	fmter.callStack = printCallStack
	return fmter
}

func (f *LogFmter) Format(entry *logrus.Entry) ([]byte, error) {
	var (
		file string
		line int
		ok   bool
	)
	if f.callStack {
		_, file, line, ok = runtime.Caller(9)
		if ok {
			_, file = filepath.Split(file)
		}
	}
	b := entry.Buffer
	if b == nil {
		b = &bytes.Buffer{}
	}
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.000"))
	b.WriteString(f.Level(entry.Level))
	b.WriteByte(' ')

	if f.callStack && ok {
		b.WriteByte('[')
		b.WriteString(file)
		b.WriteByte(':')
		b.WriteString(strconv.Itoa(line))
		b.WriteByte(']')
	}

	b.WriteString(entry.Message)
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *LogFmter) Level(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "[DEBUG]"
	case logrus.InfoLevel:
		return "[INFO ]"
	case logrus.WarnLevel:
		return "[WARN ]"
	case logrus.ErrorLevel:
		return "[ERROR]"
	case logrus.FatalLevel:
		return "[FATAL]"
	}
	return "[NIL  ]"
}
