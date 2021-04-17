package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	router := gin.Default()
	router.GET("/status", status)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":\"healthy\"}", w.Body.String())
}
