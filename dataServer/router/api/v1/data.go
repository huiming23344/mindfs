package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/mindfs/dataServer/apis"
	"github.com/huiming23344/mindfs/dataServer/server"
)

type DataServer struct {
	Id   string
	Ip   string
	Port int
}

type WriteDataReq struct {
	data        string
	dataServers []*DataServer
}

type CopyDataReq struct {
	dataServer DataServer
}

func WriteData(c *gin.Context) {
	req := WriteDataReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid data",
		})
		return
	}
	inodeId := c.GetHeader("inodeId")
	// Write data to the database
	err = server.Write(inodeId, req.data)
	for _, ds := range req.dataServers {
		if ds.Id != server.DataServer.ServiceId {
			err = apis.WriteData(req.data, apis.DataServer{
				Id:   ds.Id,
				Ip:   ds.Ip,
				Port: ds.Port,
			})
			if err != nil {
				c.JSON(500, gin.H{
					"message": "write data failed",
				})
				return
			}
		}
	}
	if err != nil {
		c.JSON(500, gin.H{
			"message": "write data failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "write data success",
	})
}

func ReadData(c *gin.Context) {
	var data string
	inodeId := c.GetHeader("inodeId")
	// Read data from the database
	data, err := server.Read(inodeId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "read data failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "read data success",
		"data":    data,
	})
	return
}

func CopyData(c *gin.Context) {
	copyReq := CopyDataReq{}
	err := c.ShouldBind(&copyReq)
	inodeId := c.GetHeader("inodeId")
	// Read data from the database
	data, err := server.Read(inodeId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "read data failed",
		})
		return
	}
	// Write data to the database
	err = apis.WriteData(data, apis.DataServer{
		Id:   copyReq.dataServer.Id,
		Ip:   copyReq.dataServer.Ip,
		Port: copyReq.dataServer.Port,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "copy data failed",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "copy data success",
		})
		return
	}
}

func UpdateData(c *gin.Context) {
	// Update data in the database
}

func DeleteData(c *gin.Context) {
	inodeId := c.GetHeader("inodeId")
	// Delete data in the database
	err := server.Delete(inodeId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "delete data failed",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "delete data success",
		})
		return
	}
}
