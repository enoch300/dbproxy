/*
* @Author: wangqilong
* @Description:
* @File: LogMiddle
* @Date: 2021/6/7 5:54 下午
 */

package middleware

import (
	"dbproxy/utils/logger"
	"github.com/gin-gonic/gin"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		logger.Global.Infof("%s | %s | %s | %d | %s | %d | %s", c.ClientIP(), c.Request.Method, c.Request.RequestURI, c.Writer.Status(), c.Request.Proto, latency, c.Request.UserAgent())
	}
}
