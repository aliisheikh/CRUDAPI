package user_repository

import (
	"ProjectCRUD/data/request"
	Models "ProjectCRUD/models"
	"errors"
	"gorm.io/gorm"
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

func (u *UserEpoImpl) FindById(userId int) (use Models.User, err error) {
	var users Models.User
	result := u.DB.Find(&users, userId)
	if result != nil {
		return use, nil
	} else {
		return users, errors.New("ID not Found")
	}

}

func (u *UserEpoImpl) Save(user Models.User) error {
	result := u.DB.Create(&user)
	if err := result; err != nil {
		return err.Error
	}
	return nil
}

func (u *UserEpoImpl) Update(user Models.User) error {
	var Updateuser = request.UpdateUserReq{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	result := u.DB.Model(&user).Updates(Updateuser)
	if err := result; err != nil {
		return err.Error
	}
	return nil
}

func NewUserEpoImpl(db *gorm.DB) *UserEpoImpl {
	return &UserEpoImpl{DB: db}

}
