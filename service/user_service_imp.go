package service

import (
	"ProjectCRUD/data/request"
	"ProjectCRUD/data/request/response"
	Models "ProjectCRUD/models"
	"ProjectCRUD/user_repository"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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

func (u *UserServiceImp) Create(users request.CreateUserReq) error {

	//if users.UserName == "" {
	//	return errors.New("username is required")
	//}
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
		//UserName: users.UserName,
		//Id:    users.Id,
		Email: users.Email,
		Name:  users.Name,
	}

	err = u.usersRepository.Save1(usersModel)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("email already exists")
		}
		fmt.Println(err)
		return err
	}
	//createUserRes := &response.UserResponse{
	//	Id:    usersModel.Id,
	//	Name:  usersModel.Name,
	//	Email: usersModel.Email,
	//}
	fmt.Println("TRDDDDD", usersModel.Id)

	return nil

}

// Delete implements userService

func (u *UserServiceImp) Delete(userID int) error {
	// Check if the user with the given ID exists
	_, err := u.usersRepository.FindById(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// If the user doesn't exist, return nil indicating successful deletion
		return nil
	} else if err != nil {
		// If an error occurred while retrieving the user, return the error
		return err
	}

	// Attempt to delete the user
	err = u.usersRepository.Delete(userID)
	if err != nil {
		// If an error occurred while deleting the user, return the error
		return err
	}

	// User successfully deleted
	return nil
}

//FindsAll Implements userService

func (u *UserServiceImp) FindAll() []response.UserResponse {
	result := u.usersRepository.FindAll()
	var users []response.UserResponse
	for _, v := range result {
		user := response.UserResponse{
			Id: v.Id,
			//UserName: v.UserName,
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
		Id: userData.Id,
		//UserName: userData.UserName,
		Email: userData.Email,
		Name:  userData.Name,
	}
	//fmt.Println(userResponse)
	return userResponse, nil
}

func (u *UserServiceImp) Update(user request.UpdateUserReq) error {

	if user.Id == 0 && user.Name == "" && user.Email == "" {
		return errors.New("at least one field is required for the update")
	}

	userData, err := u.usersRepository.FindById(user.Id)
	if err != nil {
		return err
	}
	//if userData.UserName != "" {
	//	userData.UserName = user.UserName
	//}
	if userData.Email != "" {
		userData.Email = user.Email
	}
	if userData.Name != "" {
		userData.Name = user.Name
	}

	if err := u.usersRepository.Save(*userData); err != nil {
		return err
	}
	return nil
}
