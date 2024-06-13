package service

import (
	"ProjectCRUD/data/request"
	Models "ProjectCRUD/models"
)

type ProfileService interface {
	CreateP(profile request.CreateProfileReq) error
	UpdateP(profile request.UpdateProfileReq) error
	DeleteP(userId, profileId int) error
	FindByIdP(userId, profileId int) (*Models.ProfileModel, error)
	FindAllProfilesByUserID(userID int) ([]Models.ProfileModel, error)
}
