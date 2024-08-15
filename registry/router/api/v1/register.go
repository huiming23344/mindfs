package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/mindfs/registry/server"
)

type ServiceReq struct {
	ServiceName string `json:"serviceName"`
	ServiceId   string `json:"serviceId"`
	IpAddress   string `json:"ipAddress"`
	Port        int    `json:"port"`
}

func Register(c *gin.Context) {
	var body ServiceReq
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "BindJSON failed",
		})
		return
	}
	server.RegisterService(server.ServiceReq{
		ServiceName: body.ServiceName,
		ServiceId:   body.ServiceId,
		IpAddress:   body.IpAddress,
		Port:        body.Port,
	})
	c.JSON(200, gin.H{
		"message": "register success",
	})
}

func Unregister(c *gin.Context) {
	var body ServiceReq
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "BindJSON failed",
		})
		return
	}
	success := server.UnregisterService(server.ServiceReq{
		ServiceName: body.ServiceName,
		ServiceId:   body.ServiceId,
		IpAddress:   body.IpAddress,
		Port:        body.Port,
	})
	if success {
		c.JSON(200, gin.H{
			"message": "unregister success",
		})
	} else {
		c.JSON(200, gin.H{
			"error": "unregister failed",
		})
	}
}
