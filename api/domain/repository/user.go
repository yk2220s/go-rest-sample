package repository

import (
	"github.com/yk2220s/go-rest-sample/api/database"
	"github.com/yk2220s/go-rest-sample/api/domain/derror"
	"github.com/yk2220s/go-rest-sample/api/domain/model"
)

// UserRepository access UserEntity
type UserRepository interface {
	List(page int) ([]*model.User, error)
	GetByID(userID int) (*model.User, derror.DomainError)
	Store(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(userID int) (bool, error)
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
func (repo *UserRepositoryImpl) List(page int) ([]*model.User, error) {
	db := database.Open()
	defer db.Close()

	var users []*model.User

	limit := 10
	offset := limit * (page - 1)
	db.Limit(limit).Offset(offset).Find(&users)

	return users, nil
}

// GetByID fetch user by id
func (repo *UserRepositoryImpl) GetByID(userID int) (*model.User, derror.DomainError) {
	db := database.Open()
	defer db.Close()

	var user model.User

	if db.First(&user, userID).RecordNotFound() {
		return nil, derror.NewNotFound()
	}

	return &user, nil
}

// Store save new user model.
func (repo *UserRepositoryImpl) Store(user *model.User) (*model.User, error) {
	db := database.Open()
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Update update user model
func (repo *UserRepositoryImpl) Update(user *model.User) (*model.User, error) {
	db := database.Open()
	defer db.Close()

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Update update user model
func (repo *UserRepositoryImpl) Delete(userID int) (bool, error) {
	db := database.Open()
	defer db.Close()

	var user model.User

	if db.First(&user, userID).RecordNotFound() {
		return false, derror.NewNotFound()
	}

	db.Delete(&user)

	return true, nil
}
