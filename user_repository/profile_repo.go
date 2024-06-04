package user_repository

import Models "ProjectCRUD/models"

type ProfileRepository interface {
	Save2(prof *Models.ProfileModel) error
	UpdateP(prof Models.ProfileModel) error
	DeleteP(profId int)
	FindByIdP(profId int) (prof Models.ProfileModel, err error)
}
