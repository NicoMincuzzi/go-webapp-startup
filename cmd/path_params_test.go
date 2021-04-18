package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSingleParam(t *testing.T) {
	router := gin.Default()
	router.GET("/user/:name", singleParam)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/buddy", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"name\":\"buddy\"}", w.Body.String())
}

func TestMoreParams(t *testing.T) {
	router := gin.Default()
	router.GET("/user/:name/*action", moreParams)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/buddy/walk", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"name\":\"buddy\",\"action\":\"/walk\"}", w.Body.String())
}
