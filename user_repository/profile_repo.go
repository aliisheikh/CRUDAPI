package user_repository

import Models "ProjectCRUD/models"

type ProfileRepository interface {
	Save2(prof *Models.ProfileModel)
	UpdateP(prof Models.ProfileModel)
	DeleteP(prof int)
	FindByIdP(profId int) (prof Models.ProfileModel, err error)
}
