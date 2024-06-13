package user_repository

import (
	Models "ProjectCRUD/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	//"regexp"
)

type ProfileEPOImpl struct {
	DB *gorm.DB
}

func NewProfileRepositoryImp(Db *gorm.DB) *ProfileEPOImpl {
	return &ProfileEPOImpl{DB: Db}
}

// DeleteP function for profile

func (p *ProfileEPOImpl) DeleteP(userid, profId int) error {

	profiles := Models.ProfileModel{ProfileId: profId, UserID: userid}
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

func (p *ProfileEPOImpl) FindByIdP(userId, profId int) (*Models.ProfileModel, error) {

	if userId == 0 {
		return nil, errors.New("invalid userId")
	}

	if profId == 0 {
		// Return an error indicating that the user ID is invalid
		return nil, errors.New("invalid profile ID")
	}
	var profile Models.ProfileModel
	result := p.DB.Preload("User").First(&profile, "Id = ? AND userId = ?", profId, userId)

	if result.Error != nil {
		// Check if the error is due to record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("profile not found")
		}
		// Handle other errors
		return nil, result.Error
	}
	return &profile, nil
}

// Update function for profile

func (p *ProfileEPOImpl) UpdateP(profile Models.ProfileModel) error {

	if profile.ProfileName == "" && profile.Phone == "" && profile.Address == "" {
		return errors.New("at least one field is required for the update")
	}
	// Check if the user exists before updating
	existingUser, err := p.FindByIdP(profile.UserID, profile.ProfileId)
	if err != nil {
		// Forward the error
		return err
	}

	existingUser.ProfileName = profile.ProfileName
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
	existingProfile, err := p.FindByIdP(profile.UserID, profile.ProfileId)
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

func (p *ProfileEPOImpl) SaveCreate(profile Models.ProfileModel, userID int) error {
	// Set the UserID for the profile
	profile.UserID = userID

	// Check if the UserID exists in the users table
	var user Models.User
	if err := p.DB.First(&user, profile.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user with the given ID does not exist")
		}
		return fmt.Errorf("failed to retrieve user: %w", err)
	}

	// Save the profile
	result := p.DB.Create(&profile)
	if result.Error != nil {
		return fmt.Errorf("failed to save the profile: %w", result.Error)
	}

	return nil
}

func (p *ProfileEPOImpl) FindAllProfilesByUserID(userId int) ([]Models.ProfileModel, error) {
	var profiles []Models.ProfileModel
	var user Models.User

	if err := p.DB.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User with the specified ID not found
			return nil, fmt.Errorf("user with ID %d not found", userId)
		}
		// Other database error occurred
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	// Use Preload to include the associated User data
	if err := p.DB.Model(&Models.ProfileModel{}).Preload("User").Where("userId = ?", userId).Find(&profiles).Error; err != nil {
		return nil, err
	}

	for _, profile := range profiles {
		profile.User = user
	}

	//for i := range profiles {
	//	profiles[i].User = user
	//}
	return profiles, nil
}
