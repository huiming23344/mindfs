package v1

import "github.com/gin-gonic/gin"

func OpenFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "open file",
	})
}
