package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryParams(t *testing.T) {
	router := gin.Default()
	router.GET("/user", queryParams)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user?firstname=Nick&lastname=Mincuzzi&nickname=Thor", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"NickMincuzzi(Thor)\"}", w.Body.String())
}
