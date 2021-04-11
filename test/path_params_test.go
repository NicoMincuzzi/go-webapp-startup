package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-webapp/cmd"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSingleParam(t *testing.T) {
	router := gin.Default()
	router.GET("/user/:name", main.SingleParam)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/buddy", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"name\":\"buddy\"}", w.Body.String())
}

func TestMoreParams(t *testing.T) {
	router := gin.Default()
	router.GET("/user/:name/*action", main.MoreParams)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/buddy/walk", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"buddy is /walk\"}", w.Body.String())
}
