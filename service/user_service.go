package service

import (
	"ProjectCRUD/data/request"
	"ProjectCRUD/data/request/response"
	Models "ProjectCRUD/models"
)

type UserService interface {
	Create(user request.CreateUserReq)
	Update(user request.UpdateUserReq) error
	Delete(userId int) error
	FindById(userId int) (*Models.User, error)
	FindAll() []response.UserResponse
}
