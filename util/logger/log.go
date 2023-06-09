package logger

import (
	"bytes"
	"domo1/config"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type MyFormatter struct{}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	var newLog string

	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		// fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d] %s\n",
			timestamp, entry.Level, entry.Caller.File, entry.Caller.Line, entry.Message) //format位置-----------------------------------
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

// Init logrus logger.
func InitLogger() {
	// 设置日志格式。
	logrus.SetFormatter(&MyFormatter{})

	switch config.Conf.LogConfig.Level {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	}
	logrus.SetReportCaller(true) // 打印文件、行号和主调函数。

	// 实现日志滚动。
	// Refer to https://www.cnblogs.com/jssyjam/p/11845475.html.
	logger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%v/%v%v", config.Conf.LogConfig.Dir, time.Now().Format("2006-01-02"), config.Conf.LogConfig.Name), // 日志输出文件路径，以日期命名。
		MaxSize:    config.Conf.LogConfig.MaxSize,                                                                                  // 日志文件最大 size(MB)，缺省 100MB。
		MaxBackups: config.Conf.LogConfig.MaxBackups,                                                                               // 最大过期日志保留的个数。
		MaxAge:     config.Conf.LogConfig.MaxAge,                                                                                   // 保留过期文件的最大时间间隔，单位是天。
		LocalTime:  true,                                                                                                           // 是否使用本地时间来命名备份的日志。
	}
	logrus.SetOutput(logger)
}
