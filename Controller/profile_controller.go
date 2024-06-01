package Controller

import (
	"ProjectCRUD/data/request"
	"ProjectCRUD/data/request/response"
	"ProjectCRUD/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
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
	// Initialize a CreateUserReq instance
	var createProfileRequest request.CreateProfileReq

	// Bind JSON data from the request to createProfileRequest
	if err := ctx.ShouldBindJSON(&createProfileRequest); err != nil {
		// If JSON binding fails, respond with a bad request error
		var errorMsg string
		if verr, ok := err.(validator.ValidationErrors); ok {
			var fields []string
			for _, fieldErr := range verr {
				fieldName := fieldErr.StructField()
				fields = append(fields, fieldName)
			}
			errorMsg = fmt.Sprintf("Missing or invalid fields: %s", strings.Join(fields, ", "))
		} else {
			errorMsg = "Failed to Parse JSON"
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
		return
	}

	// Create the user using the userService
	if err := profilecontroller.profileService.CreateP(createProfileRequest); err != nil {
		// duplicate email error
		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusConflict, gin.H{"error": "ProfileName already exists"})
			return
		}

		// If user creation fails, respond with an internal server error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ProfileName already exists"})
		return
	}

	// Respond with a success message and status code 201 Created
	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   createProfileRequest, // You can choose to include the created user data here
	}
	ctx.JSON(http.StatusCreated, webResponse)
}

// Update profile controller

func (profilecontroller *ProfileController) UpdateP(c *gin.Context) {
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse the JSON request body into update user request struct
	var updateprofilerequest request.UpdateProfileReq
	if err := c.ShouldBindJSON(&updateprofilerequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	// Set the ID field of UpdateUserRequest with the id value
	updateprofilerequest.ProfileId = id

	// Call the service method to update the user
	if err := profilecontroller.profileService.UpdateP(updateprofilerequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Fetch the updated user data
	updatedProfile, err := profilecontroller.profileService.FindByIdP(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated user data"})
		return
	}

	// Respond with success message and updated user data
	c.JSON(http.StatusOK, gin.H{"data": updatedProfile})
}

//DELETEP function for profile

// DELETE USER
func (profilecontroller *ProfileController) DeleteP(c *gin.Context) {
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Attempt to delete the user
	err = profilecontroller.profileService.DeleteP(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"error": "Failed to delete the Profile"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "profile already been deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile deleted successfully"})
}

// FindById of Profile

func (profilecontroller *ProfileController) FindByIdP(c *gin.Context) {
	userId := c.Param("userId")
	fmt.Println("userId", userId)
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	p, err := profilecontroller.profileService.FindByIdP(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find the user"})
		return
	}

	if p == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}
