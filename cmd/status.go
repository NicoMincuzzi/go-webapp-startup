package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func status (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
