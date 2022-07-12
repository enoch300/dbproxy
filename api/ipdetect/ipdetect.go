package ipdetect

import (
	"dbproxy/model/ipdetect"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	Appid        string         `json:"appid"`
	SrcMachineId string         `json:"src_machine_id"`
	Data         []ipdetect.Row `json:"data"`
}

func Insert(c *gin.Context) {
	var request Request
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}

	if err := ipdetect.Insert(request.Data); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}
