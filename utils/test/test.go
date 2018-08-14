package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/ProyectoIT/InstagramAnalyticsAPI/app"
	"github.com/ProyectoIT/InstagramAnalyticsAPI/dao/mysql"
	"github.com/gin-gonic/gin"
)

func StartTest() *gin.Engine {
	r := gin.Default()
	app.MapEndpoints(r)
	mysql.InitDB()
	return r
}

func MakeRequest(r *gin.Engine, method string, url string, body string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	r.ServeHTTP(w, req)
	return w
}
