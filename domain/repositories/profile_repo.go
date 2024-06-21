package repositories

import (
	"ProjectCRUD/infrastructure/models"
)

type ProfileRepository interface {
	Save2(prof *Models.ProfileModel) error
	UpdateP(prof Models.ProfileModel) error
	DeleteP(userId, profId int) error
	FindByIdP(userId, profId int) (prof Models.ProfileModel, err error)
	FinAllProfilesByUserId(userId int) (profiles []Models.ProfileModel, err error)
}
