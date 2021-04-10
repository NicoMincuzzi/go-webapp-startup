package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/status", status)
	router.GET("/user/:name", singleParam)
	router.GET("/user/:name/*action", moreParams)
	router.GET("/user", queryParams)

	router.POST("/user", modelBindingValidation)

	router.Run(":3030")
}
