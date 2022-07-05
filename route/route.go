/*
* @Author: wangqilong
* @Description:
* @File: route
* @Date: 2021/11/30 3:32 下午
 */

package route

import (
	"dbproxy/api/detect"
	"dbproxy/api/httpquality"
	"dbproxy/middleware"
	"dbproxy/utils/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net"
)

func Listen() {
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	r.Use(gin.Recovery())
	r.Use(middleware.Log())

	v1 := r.Group("/api/v1")
	v1.POST("/ck/detect", detect.Insert)
	v1.POST("/ck/httpquality", httpquality.Insert)
	if err := r.Run(net.JoinHostPort(viper.GetString("server.ip"), viper.GetString("server.port"))); err != nil {
		logger.Global.Fatalf("Run app %s", err.Error())
	}
}
