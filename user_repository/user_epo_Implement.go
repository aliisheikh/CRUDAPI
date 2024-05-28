package user_repository

import (
	Models "ProjectCRUD/models"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type UserEpoImpl struct {
	DB *gorm.DB
}

func (u UserEpoImpl) Delete(userId int) error {
	var u_user Models.User
	result := u.DB.Where("id=?", userId).Delete(&u_user)

	if err := result; err != nil {
		return err.Error
	}
	return nil

}

func (u *UserEpoImpl) FindAll() []Models.User {
	//panic("can't implement me")
	var users []Models.User
	result := u.DB.Find(&users)
	if result != nil {
		panic(result.Error)
	}
	return users
}

func (u *UserEpoImpl) FindById(userId int) (*Models.User, error) {
	if userId == 0 {
		// Return an error indicating that the user ID is invalid
		return nil, errors.New("invalid user ID")
	}

	var user Models.User
	result := u.DB.First(&user, userId)
	if result.Error != nil {
		// Check if the error is due to record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		// Handle other errors
		return nil, result.Error
	}
	return &user, nil
}

// For just simple POST

func (u *UserEpoImpl) Save1(user Models.User) error {
	// Save the user directly, GORM will handle whether it's a new user or an existing one
	result := u.DB.Save(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			return errors.New("user already exists")
		}
		return result.Error
	}
	return nil
}

// Save for Update,GET and DELETE
func (u *UserEpoImpl) Save(user Models.User) error {
	// Check if the user already exists
	existingUser, err := u.FindById(user.Id)
	if err != nil {
		// If the user does not exist, create a new record
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := u.DB.Create(&user)
			if result.Error != nil {
				return result.Error
			}
			return nil
		}
		// Handle other errors
		return err
	}
	// Update the existing user's data
	existingUser.UserName = user.UserName
	existingUser.Email = user.Email
	existingUser.Name = user.Name

	// Save the updated user data
	result := u.DB.Save(existingUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserEpoImpl) Update(user Models.User) error {
	// Check if the user exists before updating
	existingUser, err := u.FindById(user.Id)
	if err != nil {
		// Forward the error
		return err
	}

	// Update the existing user's data
	existingUser.UserName = user.UserName
	existingUser.Email = user.Email
	existingUser.Name = user.Name

	// Save the updated user data
	result := u.DB.Save(existingUser)
	if result.Error != nil {
		// Handle the error
		return result.Error
	}
	return nil
}
func NewUserEpoImpl(db *gorm.DB) *UserEpoImpl {
	return &UserEpoImpl{DB: db}
}
