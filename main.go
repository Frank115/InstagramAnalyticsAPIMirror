package main

import "github.com/gin-gonic/gin"
import "github.com/ProyectoIT/InstagramAnalyticsAPI/controller"


func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}