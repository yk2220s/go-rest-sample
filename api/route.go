package api

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yk2220s/go-rest-sample/api/application/controller"
)

// SetupRouter is return router.
func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	errlogfile, _ := os.Create("error.log")
	gin.DefaultErrorWriter = io.MultiWriter(errlogfile)

	r := gin.Default()

	// Register Error Handler
	// r.Use(handleErrors)

	// Index
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	// User resource.
	uController := controller.UserControllerFactory()
	r.GET("/users", uController.ListUser)
	r.GET("/users/:id", uController.GetUser)
	r.POST("/users", uController.CreateUser)
	r.PATCH("/users/:id", uController.UpdateUser)
	r.DELETE("/users/:id", uController.DeleteUser)

	return r
}
