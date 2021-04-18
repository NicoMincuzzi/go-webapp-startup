package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func queryParams(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Request.URL.Query().Get("lastname")
	nickname := c.Query("nickname") // shortcut for c.Request.URL.Query().Get("lastname")

	message := firstname + lastname + "(" + nickname + ")"
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
