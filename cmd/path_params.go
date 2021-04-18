package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct{
	Name   string `json:"name"`
	Action string `json:"action,omitempty"`
}

func singleParam(c *gin.Context) {
	name := c.Param("name")

	var user User
	user.Name = name
	c.JSON(http.StatusOK, user)
}

func moreParams(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")

	var user User
	user.Name = name
	user.Action = action
	c.JSON(http.StatusOK, user)
}
