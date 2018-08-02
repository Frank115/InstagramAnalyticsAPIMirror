package controller

import (
	"net/http"

	"github.com/ProyectoIT/InstagramAnalyticsAPI/domain"
	"github.com/ProyectoIT/InstagramAnalyticsAPI/service"
	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	var u domain.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	err = service.SaveUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "El usuario fue creado"})
}
