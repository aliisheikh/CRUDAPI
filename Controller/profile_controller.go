package Controller

import (
	"ProjectCRUD/data/request"
	"ProjectCRUD/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProfileController struct {
	profileService service.ProfileService
}

func NewProfileController(profileService service.ProfileService) *ProfileController {
	return &ProfileController{
		profileService: profileService,
	}
}

// Create Profile Function

func (profilecontroller *ProfileController) CreateP(ctx *gin.Context) {
	// Extract the userId from the URL parameter
	userIDStr := ctx.Param("userId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Initialize a CreateProfileReq instance
	var createProfileRequest request.CreateProfileReq

	// Bind JSON data from the request to createProfileRequest
	if err := ctx.ShouldBindJSON(&createProfileRequest); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Set the UserID from the URL parameter
	if len(createProfileRequest.Phone) < 1 || len(createProfileRequest.Phone) > 11 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Phone number must be between 1 and 11 characters"})
		return
	}

	createProfileRequest.UserId = userID

	// Create the profile using the profileService
	if err := profilecontroller.profileService.CreateP(createProfileRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Profile"})
		return
	}

	// Respond with a success message and status code 201 Created
	ctx.JSON(http.StatusCreated, gin.H{"message": "Profile created successfully"})
}

// Update profile controller

func (profilecontroller *ProfileController) UpdateP(c *gin.Context) {
	profId := c.Param("profId")
	userId := c.Param("userId")

	profIdInt, err := strconv.Atoi(profId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse the JSON request body into update user request struct
	var updateProfileRequest request.UpdateProfileReq
	if err := c.ShouldBindJSON(&updateProfileRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	if updateProfileRequest.ProfileName == "" && updateProfileRequest.Phone == "" && updateProfileRequest.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one field is required for the update"})
		return
	}

	if len(updateProfileRequest.Phone) < 1 || len(updateProfileRequest.Phone) > 11 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number must be between 1 and 11 characters"})
		return
	}

	// Set the ID fields of UpdateProfileRequest with the values from the path
	updateProfileRequest.ProfileId = profIdInt
	updateProfileRequest.UserId = userIdInt

	// Call the service method to update the user
	if err := profilecontroller.profileService.UpdateP(updateProfileRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// DELETE Profile
func (profilecontroller *ProfileController) DeleteP(c *gin.Context) {
	profId := c.Param("profId")
	userId := c.Param("userId")

	profIdInt, err := strconv.Atoi(profId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Attempt to delete the profile
	err = profilecontroller.profileService.DeleteP(userIdInt, profIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete the profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}

// FindById of Profile

func (profilecontroller *ProfileController) FindByIdP(c *gin.Context) {
	profId := c.Param("profId")
	userId := c.Param("userId")

	profIdInt, err := strconv.Atoi(profId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Retrieve the profile by ID
	profile, err := profilecontroller.profileService.FindByIdP(userIdInt, profIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find the profile"})
		return
	}

	// If profile is not found, return an error
	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

// Find all profiles against the single user
func (profileController *ProfileController) FindAllProfilesByUserID(ctx *gin.Context) {
	// Extract the user ID from the request parameters
	userIDStr := ctx.Param("userId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the service method to fetch all profiles by user ID
	profiles, err := profileController.profileService.FindAllProfilesByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch profiles", "Internal": err.Error()})
		return
	}

	// If no profiles are found, return an appropriate message
	if len(profiles) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No profiles found for this user"})
		return
	}

	// Respond with the list of profiles and associated user details
	ctx.JSON(http.StatusOK, gin.H{"data": profiles})
}
