package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/huiming23344/mindfs/metaServer/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	userApi := r.Group("/user")
	{
		userApi.POST("/add", v1.AddUser)
		userApi.POST("/delete/:name", v1.DeleteUser)
		userApi.GET("/list", v1.ListUser)
		userApi.POST("/group/add", v1.AddGroup)
		userApi.POST("/group/delete/:name", v1.DeleteGroup)
		userApi.POST("/group/addUser", v1.AddUserToGroup)
		userApi.POST("/group/deleteUser", v1.RemoveUserFromGroup)
	}
	dataSrvApi := r.Group("/dataServer")
	{
		dataSrvApi.POST("/invalid", v1.Invalid)
	}

	return r
}
