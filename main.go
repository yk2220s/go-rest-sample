package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yk2220s/go-rest-sample/api/controller"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Index
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	// User resource.
	r.GET("/users", controller.UserList)

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		c.JSON(http.StatusOK, gin.H{"userId": id})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
