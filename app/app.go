package app

import (
	"log"

	"github.com/ProyectoIT/InstagramAnalyticsAPI/config"
	"github.com/ProyectoIT/InstagramAnalyticsAPI/dao/mysql"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	MapEndpoints(r)
	if err := mysql.InitDB(); err != nil {
		log.Fatal(err)
		return
	}
	r.Run(":" + config.PORT) // listen and serve on 0.0.0.0:8080

}
