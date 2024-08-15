package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/mindfs/registry/server"
)

func Discovery(c *gin.Context) {
	var serviceRspList []ServiceReq
	serviceName := c.GetHeader("serviceName")
	fmt.Printf("serviceName: %s\n", serviceName)
	if serviceName == "" {
		serviceList := server.GetAllService()
		for _, service := range serviceList {
			serviceRspList = append(serviceRspList, ServiceReq{
				ServiceName: service.ServiceName,
				ServiceId:   service.ServiceId,
				IpAddress:   service.IpAddress,
				Port:        service.Port,
			})
		}
	} else {
		service := server.GetService(serviceName)
		if service != nil {
			serviceRspList = append(serviceRspList, ServiceReq{
				ServiceName: service.ServiceName,
				ServiceId:   service.ServiceId,
				IpAddress:   service.IpAddress,
				Port:        service.Port,
			})
		}
	}
	jsonData, err := json.Marshal(serviceRspList)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "json marshal failed",
		})
		return
	}
	c.String(200, string(jsonData))
}
