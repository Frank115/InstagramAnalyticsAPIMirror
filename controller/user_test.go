package controller_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ProyectoIT/InstagramAnalyticsAPI/utils/test"
	"github.com/stretchr/testify/assert"
)

func TestPostUser(t *testing.T) {
	jsonUser := `{"username":"rakki","password":"cono123"}`
	r := test.StartTest()
	w := test.MakeRequest(r, "POST", "/users", jsonUser)
	assert.Equal(t, http.StatusCreated, w.Code)
	var mp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &mp)
	assert.Equal(t, "success", mp["status"])
	assert.Equal(t, "El usuario fue creado", mp["message"])
}

func TestPostUserFailsIfParamBadType(t *testing.T) {
	jsonUser := `{"username":123,"password":"cono123"}`
	r := test.StartTest()
	w := test.MakeRequest(r, "POST", "/users", jsonUser)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
