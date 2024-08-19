package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/huiming23344/mindfs/metaServer/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	dataSrvApi := r.Group("/user")
	{
		dataSrvApi.POST("/add", v1.AddUser)
		dataSrvApi.POST("/delete/:name", v1.DeleteUser)
		dataSrvApi.GET("/list", v1.ListUser)
		dataSrvApi.POST("/group/add", v1.AddGroup)
		dataSrvApi.POST("/group/delete/:name", v1.DeleteGroup)
		dataSrvApi.POST("/group/addUser", v1.AddUserToGroup)
		dataSrvApi.POST("/group/deleteUser", v1.RemoveUserFromGroup)
	}
	return r
}
