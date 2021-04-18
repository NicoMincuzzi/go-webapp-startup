package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var message struct {
	Status string `json:"status"`
}

func status(c *gin.Context) {
	message.Status = "healthy"
	c.JSON(http.StatusOK, message)
}
