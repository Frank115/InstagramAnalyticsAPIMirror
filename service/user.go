package service

import (
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
