package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yk2220s/go-rest-sample/api/model"
)

// UserList fetch List of users
func UserList(c *gin.Context) {
	users := [2]model.User{
		model.User{UserID: 1, Name: "Sam"},
		model.User{UserID: 2, Name: "Elle"},
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
