package controller

import (
	"net/http"
	"strconv"

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
	err = service.CreateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "El usuario fue creado"})
}
func DeleteUser(c *gin.Context) {
	// dar error si el user id no es valido
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil { // solo me molesta si no hay nada, porque en ese caso gorm decide borrar todo
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "El usuario es invalido", "error": err.Error()})
		return
	}
	//error handling over db
	if err := service.DeleteUser(id); err != nil {
		c.JSON(err.Status, gin.H{"status": "error", "message": err.Message, "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "El usuario ha sido borrado"})
}
func Put(c *gin.Context) {
	var u domain.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if u.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Debe indicar el user_id"})
		return
	}
	err = service.UpdateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "El usuario fue modificado"})
}
