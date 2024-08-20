package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huiming23344/mindfs/metaServer/server"
)

func OpenFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "open file",
	})
}

func CreateFile(c *gin.Context) {
	var path, fileSystemName string
	path = c.GetHeader("path")
	fileSystemName = c.GetHeader("fileSystemName")
	dir := server.MetaServer.FileSys[fileSystemName]
	if dir == nil {
		c.JSON(404, gin.H{
			"message": "file system not found",
		})
		return
	}
	_, err := dir.FindDir(path)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "path not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "create file",
	})
}

func CreateDir(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create dir",
	})
}
