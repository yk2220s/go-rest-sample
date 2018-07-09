package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/yk2220s/go-rest-sample/api/database"
	"github.com/yk2220s/go-rest-sample/api/domain/model"
	"github.com/yk2220s/go-rest-sample/api/domain/repository"
)

// ListUser fetch List of users
func ListUser(c *gin.Context) {
	var repo repository.UserRepository
	repo = repository.UserRepositoryFactory()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	users, _ := repo.List(page)

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser fetch List of users
func GetUser(c *gin.Context) {
	db := database.Open()
	defer db.Close()

	userID, _ := strconv.Atoi(c.Params.ByName("id"))

	var user model.User
	if db.First(&user, userID).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"user": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

type paramPostUser struct {
	User struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"omitempty,max=255,email"`
	} `json:"user"`
}

// CreateUser create User record.
func CreateUser(c *gin.Context) {
	db := database.Open()
	defer db.Close()

	var puser paramPostUser
	if err := c.ShouldBindJSON(&puser); err == nil {
		user := model.User{
			Name:  puser.User.Name,
			Email: puser.User.Email,
		}
		db.Create(&user)

		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

type paramPatchUser struct {
	User struct {
		Name  string `json:"name" binding:""`
		Email string `json:"email" binding:"omitempty,max=255,email"`
	} `json:"user"`
}

// UpdateUser create User record.
func UpdateUser(c *gin.Context) {
	db := database.Open()
	defer db.Close()

	userID, _ := strconv.Atoi(c.Params.ByName("id"))

	var puser paramPatchUser
	if err := c.ShouldBindJSON(&puser); err == nil {
		var user model.User
		if db.First(&user, userID).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"user": nil})
			return
		}

		user.Name = puser.User.Name
		user.Email = puser.User.Email

		db.Save(&user)

		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
