package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/status", Status)
	router.GET("/user/:name", SingleParam)
	router.GET("/user/:name/*action", MoreParams)
	router.GET("/user", queryParams)

	router.POST("/user", modelBindingValidation)

	router.Run(":3030")
}
