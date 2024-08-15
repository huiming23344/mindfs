package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/huiming23344/mindfs/dataServer/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	dataSrvApi := r.Group("/dataServer")
	dataSrvApi.POST("/writeData", v1.WriteData)
	dataSrvApi.GET("/readData", v1.ReadData)
	dataSrvApi.POST("/updateData", v1.UpdateData)
	dataSrvApi.GET("/deleteData", v1.DeleteData)
	return r
}
