package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/yk2220s/go-rest-sample/api/database"
	"github.com/yk2220s/go-rest-sample/api/domain/model"
	"github.com/yk2220s/go-rest-sample/api/domain/repository"
)

// UserController is present data to show.
type UserController struct {
	uRepository repository.UserRepository
}

// UserControllerFactory create UserController injected depency
func UserControllerFactory() UserController {
	var urepo repository.UserRepository
	urepo = repository.UserRepositoryFactory()

	return UserController{urepo}
}

// ListUser fetch List of users
func (controller UserController) ListUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	users, _ := controller.uRepository.List(page)

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser fetch List of users
func (controller UserController) GetUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Params.ByName("id"))

	user, err := controller.uRepository.GetByID(userID)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(err.StatusCode(), gin.H{"user": nil})
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
