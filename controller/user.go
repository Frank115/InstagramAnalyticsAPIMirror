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
	err = service.CreateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "El usuario fue creado"})
}
func DeleteUser(c *gin.Context) {
	// dar error si el user id no es valido
	var u domain.User

	id := c.Param("user_id")
	if id == "" { // solo me molesta si no hay nada, porque en ese caso gorm decide borrar todo
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "El usuario es invalido"})
		return
	}

	err := db.Where("user_id = ?", id).First(&u).Error
	if err != nil {
		//error handling?
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "El usuario no ha sido encontrado"})
	}
	//borrar de la db
	db.First(&u, "user_id = ?", id) // esto es para que u reciba la info y hacer delete (podria hacerlo solo una vez antes de err)
	db.Delete(&u)                   //puede fallar al borrar?
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "El usuario ha sido borrado"})
}
