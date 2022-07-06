package routers

import (
	"SkyWings/controller"
	"SkyWings/settings"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(settings.Conf.Mode)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/ping", controller.NetPingPong)
	v1.POST("/signup", controller.SignUpHandler)

	r.NoRoute(controller.NoRoute)
	return r
}
