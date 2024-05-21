package service

import (
	"ProjectCRUD/data/request"
	"ProjectCRUD/data/request/response"
	Models "ProjectCRUD/models"
	"ProjectCRUD/user_repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImp struct {
	usersRepository user_repository.UserEpoImpl
	validate        *validator.Validate
}

func NewUserServiceImp(usersRepository user_repository.UserEpoImpl, validate *validator.Validate) *UserServiceImp {
	return &UserServiceImp{
		usersRepository: usersRepository,
		validate:        validate,
	}
}

// Create implement userService
func (u *UserServiceImp) Create(users request.CreateUserReq) {
	err := u.validate.Struct(users)
	if err != nil {
		return
	}
	usersModel := Models.User{
		Name:  users.Name,
		Email: users.Email,
	}
	u.usersRepository.Save(usersModel)
	return

}

// Delete implements userService

func (u *UserServiceImp) Delete(usersId int) error {
	err := u.usersRepository.Delete(usersId)
	if err != nil {
		return err
	}
	return nil
}

//FindsAll Implements userService

func (u *UserServiceImp) FindAll() []response.UserResponse {
	result := u.usersRepository.FindAll()
	var users []response.UserResponse
	for _, v := range result {
		user := response.UserResponse{
			Id:   v.Id,
			Name: v.Name,
		}
		users = append(users, user)
	}
	return users
}

func (u *UserServiceImp) FindById(usersId int) (*Models.User, error) {
	userData, err := u.usersRepository.FindById(usersId)
	if err != nil {
		return nil, err
	}
	userResponse := &Models.User{
		Id:    userData.Id,
		Name:  userData.Name,
		Email: userData.Email,
	}
	//fmt.Println(userResponse)
	return userResponse, nil
}
func (u *UserServiceImp) Update(users request.UpdateUserReq) {
	userData, err := u.usersRepository.FindById(users.Id)
	if err != nil {
		return
	}
	userData.Name = users.Name
	userData.Email = users.Email
	u.usersRepository.Save(userData)
	return
}
