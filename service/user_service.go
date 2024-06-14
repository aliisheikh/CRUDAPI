package service

import (
	"ProjectCRUD/data/request"
	Models "ProjectCRUD/models"
)

type UserService interface {
	Create(user request.CreateUserReq) error
	Update(user request.UpdateUserReq) error
	Delete(userId int) error
	FindById(userId int) (*Models.User, error)
}
