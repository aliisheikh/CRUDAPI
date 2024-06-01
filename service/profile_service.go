package service

import (
	"ProjectCRUD/data/request"
	Models "ProjectCRUD/models"
)

type ProfileService interface {
	CreateP(profile request.CreateProfileReq) error
	UpdateP(profile request.UpdateProfileReq) error
	DeleteP(profileId int) error
	FindByIdP(profileId int) (*Models.ProfileModel, error)
}
