package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/client"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})

	router.GET("/user/:name", func(context *gin.Context) {
		client.ServeWs(context.Writer, context.Request)
		context.String(http.StatusOK, "Hello %s", context.Param("name"))
	})

	router.Run(":8080")
}
