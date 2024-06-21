package services

import (
	"ProjectCRUD/application/data"
	"ProjectCRUD/infrastructure/models"
)

type ProfileService interface {
	CreateP(profile data.CreateProfileReq) error
	UpdateP(profile Models.ProfileModel) error
	DeleteP(userId, profileId int) error
	FindByIdP(userId, profileId int) (*Models.ProfileModel, error)
	FindAllProfilesByUserID(userID int) ([]Models.ProfileModel, error)
}
