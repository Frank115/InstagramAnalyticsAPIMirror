package service_test

import (
	"testing"

	"github.com/ProyectoIT/InstagramAnalyticsAPI/dao/mysql"
	"github.com/ProyectoIT/InstagramAnalyticsAPI/domain"
	"github.com/ProyectoIT/InstagramAnalyticsAPI/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mysql.InitDB()
	user := domain.User{Username: "rakki", Password: "cono123"}
	err := service.CreateUser(&user)
	assert.Nil(t, err)
}
