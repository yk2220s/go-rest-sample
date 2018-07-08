package repository

import (
	"github.com/yk2220s/go-rest-sample/api/model"
)

// UserRepository access UserEntity
type UserRepository interface {
	Fetch(cursor string, num int64) ([]*model.User, error)
	GetByID(id int64) (*model.User, error)
	GetByTitle(title string) (*model.User, error)
	Update(article *model.User) (*model.User, error)
	Store(a *model.User) (int64, error)
	Delete(id int64) (bool, error)
}
