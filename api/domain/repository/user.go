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
	// Store(user *model.User) (int64, error)
	// Update(user *model.User) (*model.User, error)
	// Delete(id int64) (bool, error)
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
