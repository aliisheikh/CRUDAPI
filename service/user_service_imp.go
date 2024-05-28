package service

import (
	"ProjectCRUD/data/request"
	"ProjectCRUD/data/request/response"
	Models "ProjectCRUD/models"
	"ProjectCRUD/user_repository"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
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

//var ErrDuplicateEmail = errors.New("email address already exists")
//
//func (u *UserServiceImp) Create(users request.CreateUserReq) error {
//	err := u.validate.Struct(users)
//	if err != nil {
//		return err
//	}
//
//	usersModel := Models.User{
//		UserName: users.UserName,
//		Email:    users.Email,
//		Name:     users.Name,
//	}
//
//	err = u.usersRepository.Save1(usersModel)
//	if err != nil {
//		// Check if the error message indicates a unique constraint violation
//		if isDuplicateEmailError(err) {
//			return ErrDuplicateEmail
//		}
//		// Handle other errors
//		return err
//	}
//
//	return nil
//}
//
//// Function to check if the error message indicates a duplicate email error
//func isDuplicateEmailError(err error) bool {
//	return err != nil && (strings.Contains(err.Error(), "Duplicate entry") ||
//		strings.Contains(err.Error(), "duplicate key value"))
//}

// Create implement userService
//

func (u *UserServiceImp) Create(users request.CreateUserReq) error {

	if users.UserName == "" {
		return errors.New("username is required")
	}
	if users.Name == "" {
		return errors.New("name is required")
	}
	if users.Email == "" {
		return errors.New("email is required")

	}
	fmt.Println("kkkkk", users)
	err := u.validate.Struct(users)
	if err != nil {
		return err
	}

	usersModel := Models.User{
		UserName: users.UserName,
		Email:    users.Email,
		Name:     users.Name,
	}

	err = u.usersRepository.Save1(usersModel)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("email already exists")
		}
		fmt.Println(err)
		return err
	}

	fmt.Println("TRDDDDD", usersModel.Id)

	return nil

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
			Id:       v.Id,
			UserName: v.UserName,
			Name:     v.Name,
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
		Id:       userData.Id,
		UserName: userData.UserName,
		Email:    userData.Email,
		Name:     userData.Name,
	}
	//fmt.Println(userResponse)
	return userResponse, nil
}
func (u *UserServiceImp) Update(user request.UpdateUserReq) error {
	userData, err := u.usersRepository.FindById(user.Id)
	if err != nil {
		return err
	}

	userData.UserName = user.UserName
	userData.Email = user.Email
	userData.Name = user.Name

	if err := u.usersRepository.Save(*userData); err != nil {
		return err
	}
	return nil
}
