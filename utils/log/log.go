/*
* @Author: wangqilong
* @Description:
* @File: log
* @Date: 2021/11/30 3:10 下午
 */

package log

import (
	"dbproxy/utils/config"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
	"time"
)

var L *logrus.Logger
var once sync.Once

func NewLogger() {
	once.Do(func() {
		L = logrus.New()
		L.SetLevel(logrus.DebugLevel)
		logCfg := &lumberjack.Logger{
			// 日志输出文件路径
			Filename: config.Cfg.Log.FileName,
			// 日志文件最大 size, 单位是 MB
			MaxSize: config.Cfg.Log.MaxSize, // megabytes
			// 最大过期日志保留的个数
			MaxBackups: config.Cfg.Log.MaxBackups,
			// 保留过期文件的最大时间间隔,单位是天
			MaxAge: config.Cfg.Log.MaxAge, //days
			// 是否需要压缩滚动日志, 使用的 gzip 压缩
			Compress: config.Cfg.Log.Compress, // disabled by default
		}
		L.SetFormatter(&nested.Formatter{
			HideKeys:        true,
			TimestampFormat: time.RFC3339,
		})
		L.SetOutput(logCfg)
	})
}
