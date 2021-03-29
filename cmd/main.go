package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/ping", status)

	router.Run(":3030")
}
