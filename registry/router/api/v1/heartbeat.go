package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/mindfs/registry/server"
)

type HeartbeatReq struct {
	ServiceId string `json:"serviceId"`
	IpAddress string `json:"ipAddress"`
	Port      int    `json:"port"`
}

func Heartbeat(c *gin.Context) {
	var HBbody HeartbeatReq
	err := c.BindJSON(&HBbody)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "BindJSON failed",
		})
		return
	}
	server.HeartbeatService(HBbody.ServiceId)
	c.JSON(200, gin.H{
		"message": "heartbeat success",
	})
}
