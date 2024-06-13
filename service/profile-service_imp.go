package service

import (
	"ProjectCRUD/data/request"
	Models "ProjectCRUD/models"
	"ProjectCRUD/user_repository"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ProfileServiceImp struct {
	profilerepository user_repository.ProfileEPOImpl
	userRepo          user_repository.UserEpoImpl
	validate          *validator.Validate
}

func NewProfileServiceImp(profileRepo user_repository.ProfileEPOImpl, userRepo user_repository.UserEpoImpl, validate *validator.Validate) *ProfileServiceImp {
	return &ProfileServiceImp{
		profilerepository: profileRepo,
		userRepo:          userRepo,
		validate:          validate,
	}
}

func (p *ProfileServiceImp) CreateP(profile request.CreateProfileReq) error {

	//if profile.UserId <= 0 {
	//	return errors.New("Invalid UserId")
	//}
	// Check if the field is empty
	if profile.ProfileName == "" {
		return errors.New("profileName is required")
	}
	if profile.Phone == "" {
		return errors.New("phone no. is required")

	}
	if profile.Address == "" {
		return errors.New("address is required")

	}

	if len(profile.Phone) < 1 || len(profile.Phone) > 11 {
		return errors.New("phone number must be between 1 and 11 characters")
	}
	fmt.Println("kkkkk", profile)

	profileModel := Models.ProfileModel{

		ProfileName: profile.ProfileName,
		Phone:       profile.Phone,
		Address:     profile.Address,
		UserID:      profile.UserId,
	}
	userId := profile.UserId
	err := p.profilerepository.SaveCreate(profileModel, userId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//fmt.Println("Profile Created Successfully with UserId", profile.)

	return nil

}

// DeleteP implements ProfileService
func (p *ProfileServiceImp) DeleteP(userId, profileId int) error {
	// Check if the profile with the given ID belongs to the user with the given ID
	profile, err := p.profilerepository.FindByIdP(userId, profileId)
	if err != nil {
		// If an error occurred while retrieving the profile, return the error
		return err
	}

	if profile == nil {
		// If the profile doesn't exist or doesn't belong to the user, return an error
		return errors.New("profile not found")
	}

	// Attempt to delete the profile
	err = p.profilerepository.DeleteP(userId, profileId)
	if err != nil {
		// If an error occurred while deleting the profile, return the error
		return err
	}

	// Profile successfully deleted
	return nil
}

func (p *ProfileServiceImp) FindByIdP(userId, profileId int) (*Models.ProfileModel, error) {
	userData, err := p.profilerepository.FindByIdP(userId, profileId)
	if err != nil {
		return nil, err
	}

	user, err := p.userRepo.FindById(userData.UserID)
	if err != nil {
		return nil, err

	}

	profileResponse := &Models.ProfileModel{
		ProfileId:   userData.ProfileId,
		ProfileName: userData.ProfileName,
		Phone:       userData.Phone,
		Address:     userData.Address,
		UserID:      userData.UserID,
		User:        *user,
	}
	//fmt.Println(userResponse)
	return profileResponse, nil
}

// UpdateP Function for Profile

func (p *ProfileServiceImp) UpdateP(profile request.UpdateProfileReq) error {
	// Check if no fields are provided for the update
	if profile.ProfileName == "" && profile.Phone == "" && profile.Address == "" {
		return errors.New("at least one field is required for the update")
	}

	// Check if the provided phone number is valid
	if profile.Phone != "" && (len(profile.Phone) < 1 || len(profile.Phone) > 11) {
		return errors.New("phone number must be between 1 and 11 characters")
	}

	userData, err := p.profilerepository.FindByIdP(profile.UserId, profile.ProfileId)
	if err != nil {
		return err
	}

	if profile.ProfileName != "" {
		userData.ProfileName = profile.ProfileName
	}
	if profile.Phone != "" {
		userData.Phone = profile.Phone
	}
	if profile.Address != "" {
		userData.Address = profile.Address
	}

	if err := p.profilerepository.Save2(*userData); err != nil {
		return err
	}
	return nil
}

func (s *ProfileServiceImp) FindAllProfilesByUserID(userID int) ([]Models.ProfileModel, error) {
	// Call the repository method to find all profiles by userID
	return s.profilerepository.FindAllProfilesByUserID(userID)

}
