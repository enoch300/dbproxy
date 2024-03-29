/*
* @Author: wangqilong
* @Description:
* @File: log
* @Date: 2021/11/30 3:10 下午
 */

package logger

import (
	"bufio"
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
	"time"
)

var Global *logrus.Logger

type MineFormatter struct{}

const TimeFormat = "2006-01-02 15:04:05"

func (s *MineFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s] [%s] %s\n", time.Now().Local().Format(TimeFormat), strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func writer(logPath string, level string, save uint) *rotatelogs.RotateLogs {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fileSuffix := time.Now().In(cstSh).Format("2006-01-02") + ".log"
	logier, err := rotatelogs.New(
		logPath+"_"+level+"-"+fileSuffix,
		rotatelogs.WithLinkName(logPath+"_"+level), // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(int(save)),    // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Hour*24),  // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return logier
}

func NewLogger(logPath string, app string, save uint) *logrus.Logger {
	var log = logrus.New()
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Printf("os.OpenFile %v\n", err)
	}
	output := bufio.NewWriter(src)
	log.SetOutput(output)

	logPath = path.Join(logPath, app)
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer(logPath, "debug", save), // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer(logPath, "info", save),
		logrus.WarnLevel:  writer(logPath, "warn", save),
		logrus.ErrorLevel: writer(logPath, "error", save),
		logrus.FatalLevel: writer(logPath, "fatal", save),
		logrus.PanicLevel: writer(logPath, "panic", save),
	}, &MineFormatter{})
	log.AddHook(lfHook)

	return log
}

func InitLog() {
	Global = NewLogger(viper.GetString("log.path"), "dbproxy", 3)
}
