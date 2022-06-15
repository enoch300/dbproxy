/*
* @Author: wangqilong
* @Description:
* @File: clickhouse
* @Date: 2021/11/30 3:52 下午
 */

package ck

import (
	"dbproxy/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Insert(c *gin.Context) {
	var insertInfo model.InsertData
	if err := c.BindJSON(&insertInfo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "BindJSON: " + err.Error(),
		})
		return
	}

	if err := insertInfo.Insert(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Insert:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
