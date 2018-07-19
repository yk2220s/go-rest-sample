package repository

import (
	"github.com/yk2220s/go-rest-sample/api/database"
	"github.com/yk2220s/go-rest-sample/api/domain/exception"
	"github.com/yk2220s/go-rest-sample/api/domain/model"
)

// UserRepository access UserEntity
type UserRepository interface {
	List(page int) []*model.User
	GetByID(userID int) (*model.User, exception.DomainError)
	Store(user *model.User) (*model.User, exception.DomainError)
	Update(user *model.User) (*model.User, exception.DomainError)
	Delete(userID int) (bool, exception.DomainError)
}

// UserRepositoryFactory make UserRepository
func UserRepositoryFactory() UserRepository {
	repo := UserRepositoryImpl{}

	return &repo
}

//////////////////////////////////////
// Implements
//////////////////////////////////////

// UserRepositoryImpl implements interface
type UserRepositoryImpl struct {
}

// List fetches list of users
func (repo *UserRepositoryImpl) List(page int) []*model.User {
	db := database.Open()
	defer db.Close()

	var users []*model.User

	limit := 10
	offset := limit * (page - 1)
	db.Limit(limit).Offset(offset).Find(&users)

	return users
}

// GetByID fetch user by id
func (repo *UserRepositoryImpl) GetByID(userID int) (*model.User, exception.DomainError) {
	db := database.Open()
	defer db.Close()

	var user model.User

	if db.First(&user, userID).RecordNotFound() {
		return nil, exception.NewNotFound()
	}

	return &user, nil
}

// Store save new user model.
func (repo *UserRepositoryImpl) Store(user *model.User) (*model.User, exception.DomainError) {
	db := database.Open()
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		return nil, exception.NewDatabaseError(err.Error())
	}

	return user, nil
}

// Update update user model
func (repo *UserRepositoryImpl) Update(user *model.User) (*model.User, exception.DomainError) {
	db := database.Open()
	defer db.Close()

	if err := db.Save(&user).Error; err != nil {
		return nil, exception.NewDatabaseError(err.Error())
	}

	return user, nil
}

// Delete delete user model
func (repo *UserRepositoryImpl) Delete(userID int) (bool, exception.DomainError) {
	db := database.Open()
	defer db.Close()

	var user model.User

	if db.First(&user, userID).RecordNotFound() {
		return false, exception.NewNotFound()
	}

	if err := db.Delete(&user).Error; err != nil {
		return false, exception.NewDatabaseError(err.Error())
	}

	return true, nil
}
