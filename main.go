package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})

	router.GET("/user/:name", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello %s", context.Param("name"))
	})

	router.Run(":8080")
}
