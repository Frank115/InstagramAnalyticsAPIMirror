package app

import (
	"github.com/ProyectoIT/InstagramAnalyticsAPI/config"
	"github.com/ProyectoIT/InstagramAnalyticsAPI/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.POST("/users", controller.Post)
	r.Run(":" + config.Port) // listen and serve on 0.0.0.0:8080
}
