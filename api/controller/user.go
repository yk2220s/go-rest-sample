package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yk2220s/go-rest-sample/api/database"
	"github.com/yk2220s/go-rest-sample/api/model"
)

// ListUser fetch List of users
func ListUser(c *gin.Context) {
	db := database.Open()
	defer db.Close()

	var users []model.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser fetch List of users
func GetUser(c *gin.Context) {
	db := database.Open()
	defer db.Close()

	userID := c.Params.ByName("id")

	var user model.User
	if db.First(&user, userID).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"user": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
