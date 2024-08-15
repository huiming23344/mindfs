package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/mindfs/dataServer/server"
)

func WriteData(c *gin.Context) {
	var data string
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid data",
		})
		return
	}
	inodeId := c.GetHeader("inodeId")
	// Write data to the database
	err = server.Write(inodeId, data)
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

func UpdateData(c *gin.Context) {
	// Update data in the database
}

func DeleteData(c *gin.Context) {
	// Delete data from the database
}
