package service

import (
	"ProjectCRUD/data/request"
	Models "ProjectCRUD/models"
	"ProjectCRUD/user_repository"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProfileServiceImp struct {
	profilerepository user_repository.ProfileEPOImpl
	validate          *validator.Validate
}

func NewProfileServiceImp(profileRepo user_repository.ProfileEPOImpl, validate *validator.Validate) *ProfileServiceImp {
	return &ProfileServiceImp{
		profilerepository: profileRepo,
		validate:          validate,
	}
}

func (p *ProfileServiceImp) CreateP(profile request.CreateProfileReq) error {

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
	//if profile.Age != "" {
	//	return errors.New("age is required")
	//
	//}

	fmt.Println("kkkkk", profile)
	//if err := p.validate.Struct(profile); err != nil {
	//	return err
	//}
	//ageInt, err := strconv.Atoi(profile.Age)
	//if err != nil {
	//	return errors.New("age is invalid")
	//}
	profileModel := Models.ProfileModel{

		ProfileName: profile.ProfileName,
		Phone:       profile.Phone,
		Address:     profile.Address,
		//Age:         strconv.Itoa(ageInt),
	}

	err := p.profilerepository.SaveCreate(profileModel)
	if err != nil {
		//if strings.Contains(err.Error(), "Duplicate entry") {
		//	return errors.New("profileName already exists")
		//}
		fmt.Println(err)
		return err
	}

	fmt.Println("TRDDDDD", profileModel.ProfileId)

	return nil

}

// DeleteP implements ProfileService

func (p *ProfileServiceImp) DeleteP(profileId int) error {
	// Check if the user with the given ID exists
	_, err := p.profilerepository.FindByIdP(profileId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// If the user doesn't exist, return nil indicating successful deletion
		return nil
	} else if err != nil {
		// If an error occurred while retrieving the user, return the error
		return err
	}

	// Attempt to delete the user
	err = p.profilerepository.DeleteP(profileId)
	if err != nil {
		// If an error occurred while deleting the user, return the error
		return err
	}

	// User successfully deleted
	return nil
}

func (p *ProfileServiceImp) FindByIdP(profileId int) (*Models.ProfileModel, error) {
	userData, err := p.profilerepository.FindByIdP(profileId)
	if err != nil {
		return nil, err
	}
	profileResponse := &Models.ProfileModel{
		ProfileId:   userData.ProfileId,
		ProfileName: userData.ProfileName,
		Phone:       userData.Phone,
		Address:     userData.Address,
		//Age:         userData.Age,
	}
	//fmt.Println(userResponse)
	return profileResponse, nil
}

//UpdateP Function for Profile

func (p *ProfileServiceImp) UpdateP(profile request.UpdateProfileReq) error {
	userData, err := p.profilerepository.FindByIdP(profile.ProfileId)
	if err != nil {
		return err
	}
	//if userData.UserName != "" {
	//	userData.UserName = user.UserName
	//}
	if userData.ProfileName != "" {
		userData.ProfileName = profile.ProfileName
	}
	if userData.Phone != "" {
		userData.Phone = profile.Phone
	}
	if userData.Address != "" {
		userData.Address = profile.Address
	}
	//if profile.Age != "" {
	//	// Convert profile.Age from string to int
	//	age, err := strconv.Atoi(profile.Age)
	//	if err != nil {
	//		// Handle error if conversion fails
	//		return err
	//	}

	// Assign converted age to userData.Age
	//	userData.Age = strconv.Itoa(age)
	//}
	if err := p.profilerepository.Save2(*userData); err != nil {
		return err
	}
	return nil
}
