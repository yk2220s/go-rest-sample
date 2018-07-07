package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yk2220s/go-rest-sample/api/controller"
)

// SetupRouter is return router.
func SetupRouter() *gin.Engine {
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
	r.GET("/users", controller.ListUser)
	r.GET("/users/:id", controller.GetUser)

	return r
}