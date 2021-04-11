package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SingleParam(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func MoreParams(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
