package service

import (
	"ProjectCRUD/infrastructure/models"
	"ProjectCRUD/infrastructure/repositories"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"strings"
)

type UserServiceImp struct {
	usersRepository repositories.UserEpoImpl
	validate        *validator.Validate
}

func NewUserServiceImp(usersRepository repositories.UserEpoImpl, validate *validator.Validate) *UserServiceImp {
	return &UserServiceImp{
		usersRepository: usersRepository,
		validate:        validate,
	}
}
func (u *UserServiceImp) Create(user Models.User) error {
	// Validate required fields
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}

	// Validate user using validator
	if err := u.validate.Struct(user); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Prepare user model for persistence
	userModel := Models.User{
		Email: user.Email,
		Name:  user.Name,
	}

	// Save user to repository
	if err := u.usersRepository.Save1(&userModel); err != nil {
		// Handle specific errors
		switch {
		case strings.Contains(err.Error(), "Duplicate entry"):
			return errors.New("email already exists")
		default:
			return fmt.Errorf("failed to save user: %w", err)
		}
	}

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

func (u *UserServiceImp) Update(user Models.User) error {

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
