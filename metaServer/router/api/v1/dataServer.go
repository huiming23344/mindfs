package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/mindfs/metaServer/apis"
	"github.com/huiming23344/mindfs/metaServer/server"
)

func Invalid(c *gin.Context) {
	c.GetHeader("dataServerId")
	invalidDs := server.MetaServer.Servers[c.GetHeader("dataServerId")]
	if invalidDs == nil {
		c.JSON(400, gin.H{
			"message": "invalid data server ID",
		})
		return
	}
	for _, chunk := range invalidDs.Chunks {
		ds := apis.Discovery("dataServer")
		err := apis.CopyData(chunk.Id, apis.DataServer{
			Id:   ds[0].ServiceId,
			Ip:   ds[0].IpAddress,
			Port: ds[0].Port,
		})
		if err != nil {
			c.JSON(500, gin.H{
				"message": "copy data failed",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"message": "copy data success",
	})
	return
}
