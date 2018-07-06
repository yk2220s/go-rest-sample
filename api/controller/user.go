package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserList fetch List of users
func UserList(c *gin.Context) {
	users := [2]string{"user1", "user2"}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
