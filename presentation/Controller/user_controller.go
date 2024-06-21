package Controller

import (
	"ProjectCRUD/domain/services"
	"ProjectCRUD/infrastructure/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	userService services.UserService
	//deletedUsers map[int]bool
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Original Create function
func (userController *UserController) Create(ctx *gin.Context) {
	// Initialize a User instance
	var createUserRequest Models.User

	// Bind JSON data from the request to createUserRequest
	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
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
			errorMsg = "Failed to parse JSON"
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
		return
	}

	// Create the user using the userService
	if err := userController.userService.Create(createUserRequest); err != nil {
		// Handle specific errors
		switch {
		case strings.Contains(err.Error(), "Duplicate entry"):
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email already exist"})
		}
		return
	}

	// Respond with a success message and status code 201 Created
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Update user

func (userController *UserController) Update(c *gin.Context) {

	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Parse the JSON request body into update user request struct

	// change request to model
	var updateuserrequest Models.User
	if err := c.ShouldBindJSON(&updateuserrequest); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	if updateuserrequest.Id == 0 && updateuserrequest.Email == "" && updateuserrequest.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one field is required for the update"})
		return
	}
	// Set the ID field of UpdateUserRequest with the id value
	updateuserrequest.Id = id

	// Call the services method to update the user
	if err := userController.userService.Update(updateuserrequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Fetch the updated user data
	updatedUser, err := userController.userService.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated user data"})
		return
	}

	// Respond with success message and updated user data
	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

// DELETE USER

func (userController *UserController) Delete(c *gin.Context) {
	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Attempt to delete the user
	err = userController.userService.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"error": "Failed to delete the user"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User already been deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User and associated profiles deleted successfully"})
}
func (userController *UserController) FindById(c *gin.Context) {
	userId := c.Param("userId")
	fmt.Println("userId", userId)
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	u, err := userController.userService.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find the user"})
		return
	}

	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})
}
