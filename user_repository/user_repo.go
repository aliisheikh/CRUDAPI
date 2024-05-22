package user_repository

import Models "ProjectCRUD/models"

type UserRepo interface {
	Save(user1 Models.User)
	Update(user1 Models.User) error
	Delete(userId int)
	FindById(userId int) (Models.User, error)
	FindAll() ([]Models.User, error)
}
