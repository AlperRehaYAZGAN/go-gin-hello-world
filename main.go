package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", HelloWorldIndexHandler)

	// start server on port APP_PORT
	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
		APP_PORT = "9090"
	}
	if err := r.Run(":" + APP_PORT); err != nil {
		log.Fatal(err)
	}
}

// HelloWorldIndexHandler is a simple health check endpoint
func HelloWorldIndexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "World",
	})
}
