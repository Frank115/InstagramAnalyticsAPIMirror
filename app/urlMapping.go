package app

import (
	"github.com/ProyectoIT/InstagramAnalyticsAPI/controller"
	"github.com/gin-gonic/gin"
)

// MapEndpoints - map the endpoints
func MapEndpoints(r *gin.Engine) {
	r.GET("/ping", controller.Ping)
	r.POST("/users", controller.Post)

}
