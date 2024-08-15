package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/huiming23344/mindfs/registry/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	apiv1 := r.Group("/api")
	apiv1.Use()
	apiv1.POST("register", v1.Register)
	apiv1.POST("unregister", v1.Unregister)
	apiv1.GET("discovery", v1.Discovery)
	apiv1.GET("heartbeat", v1.Heartbeat)
	return r
}
