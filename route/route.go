/*
* @Author: wangqilong
* @Description:
* @File: route
* @Date: 2021/11/30 3:32 下午
 */

package route

import (
	"dbproxy/api/ck"
	"dbproxy/middleware"
	"dbproxy/utils/config"
	"dbproxy/utils/log"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	r.Use(gin.Recovery())
	r.Use(middleware.Log())

	v1 := r.Group("/api/v1")
	v1.POST("/ck", ck.Insert)
	if err := r.Run(fmt.Sprintf("%v:%v", config.Cfg.Server.Ip, config.Cfg.Server.Port)); err != nil {
		log.L.Fatalf("Run app %s", err.Error())
	}
}
