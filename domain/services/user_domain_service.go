package services

import (
	"ProjectCRUD/infrastructure/models"
)

type UserService interface {
	Create(user Models.User) error
	Update(user Models.User) error
	Delete(userId int) error
	FindById(userId int) (*Models.User, error)
}
