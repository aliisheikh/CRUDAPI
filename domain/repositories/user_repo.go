package repositories

import (
	"ProjectCRUD/infrastructure/models"
)

type UserRepo interface {
	Save(user1 *Models.User) error
	Update(user1 Models.User) error
	Delete(userId int)
	FindById(userId int) (Models.User, error)
	//FindAll() ([]Models.User, error)
}
