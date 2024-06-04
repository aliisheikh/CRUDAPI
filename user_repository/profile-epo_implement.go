package user_repository

import (
	Models "ProjectCRUD/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ProfileEPOImpl struct {
	DB *gorm.DB
}

func NewProfileRepositoryImp(Db *gorm.DB) *ProfileEPOImpl {
	return &ProfileEPOImpl{DB: Db}
}

// DeleteP function for profile
func (p *ProfileEPOImpl) DeleteP(profId int) error {

	profiles := Models.ProfileModel{ProfileId: profId}
	result := p.DB.Where(&profiles).First(&profiles)
	if result.Error != nil {
		// Check if the user doesn't exist
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("profile not found")
		}
		// Return any other errors encountered during user retrieval
		return result.Error
	}

	// Delete the user
	result = p.DB.Delete(&profiles)
	if result.Error != nil {
		// Return any errors encountered during deletion
		return result.Error
	}

	// No error occurred, user deleted successfully
	return nil
}

func (p *ProfileEPOImpl) FindByIdP(profId int) (*Models.ProfileModel, error) {
	if profId == 0 {
		// Return an error indicating that the user ID is invalid
		return nil, errors.New("invalid user ID")
	}
	var profile Models.ProfileModel
	result := p.DB.First(&profile, profId)
	if result.Error != nil {
		// Check if the error is due to record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		// Handle other errors
		return nil, result.Error
	}
	return &profile, nil
}

// Update function for profile

func (p *ProfileEPOImpl) UpdateP(profile Models.ProfileModel) error {
	// Check if the user exists before updating
	existingUser, err := p.FindByIdP(profile.ProfileId)
	if err != nil {
		// Forward the error
		return err
	}

	// Update the existing user's data
	//	existingUser.UserName = user.UserName
	existingUser.ProfileName = profile.ProfileName
	//existingUser.Age = profile.Age
	existingUser.Phone = profile.Phone
	existingUser.Address = profile.Address

	// Save the updated user data
	result := p.DB.Save(existingUser)
	if result.Error != nil {
		// Handle the error
		return result.Error
	}
	return nil
}

// Save function for profile(UPDATE, DELETE,GET)

func (p *ProfileEPOImpl) Save2(profile Models.ProfileModel) error {
	// Check if the user already exists
	existingProfile, err := p.FindByIdP(profile.ProfileId)
	if err != nil {
		// If the user does not exist, create a new record
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := p.DB.Create(&profile)
			if result.Error != nil {
				return result.Error
			}
			return nil
		}
		// Handle other errors
		return err
	}

	// Update only the specified fields in the database
	result := p.DB.Model(&existingProfile).Updates(profile)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Save for Create

func (p *ProfileEPOImpl) SaveCreate(profile Models.ProfileModel) error {
	// Save the user directly, GORM will handle whether it's a new user or an existing one
	result := p.DB.Save(&profile)
	if result.Error != nil {
		//if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
		//	return errors.New("user already exists")
		//}
		return fmt.Errorf("Failed to save the Profile:%w", result.Error)
	}
	return nil
}
