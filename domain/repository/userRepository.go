package repository

import "github.com/rhodeon/moviescreen/domain/models"

type UserRepository interface {
	Register(user *models.User) error
	GetByEmail(email string) (models.User, error)
	Update(user *models.User) error
}
