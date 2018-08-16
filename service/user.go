package service

import (
	"net/http"
	"strconv"

	"github.com/ProyectoIT/InstagramAnalyticsAPI/dao/mysql"
	"github.com/ProyectoIT/InstagramAnalyticsAPI/domain"
)

func CreateUser(user *domain.User) (err error) {
	tx := mysql.DB.Begin()
	if err = tx.Create(user).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}
func DeleteUser(id int) (err domain.ApiError) {
	var count int
	e := mysql.DB.Model(&domain.User{}).Where("user_id = ?", id).Count(&count).Error
	if count == 0 {
		err.Status = http.StatusNotFound
		err.Err = e
		err.Message = "El usuario: " + strconv.Itoa(id) + " no ha sido encontrado."
		return
	}
	tx := mysql.DB.Begin()
	if e := tx.Where("user_id = ?", id).Delete(domain.User{}).Error; e != nil {
		err.Err = e
		err.Status = http.StatusInternalServerError
		err.Message = "Error al borrar el usuario " + strconv.Itoa(id)
		tx.Rollback()
		return
	}
	if e = tx.Commit().Error; e != nil {
		err.Err = e
		err.Status = http.StatusInternalServerError
		err.Message = "Error al borrar el usuario " + strconv.Itoa(id)
	}
	return
}
