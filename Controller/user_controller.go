package Controller

import (
	"ProjectCRUD/data/request"
	"ProjectCRUD/data/request/response"
	"ProjectCRUD/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Original Create function

func (userController *UserController) Create(ctx *gin.Context) {
	// Initialize a CreateUserReq instance
	var createuserrequest request.CreateUserReq

	// Bind JSON data from the request to createuserrequest
	if err := ctx.ShouldBindJSON(&createuserrequest); err != nil {
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
	if err := userController.userService.Create(createuserrequest); err != nil {
		// duplicate email error
		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}

		// If user creation fails, respond with an internal server error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email already exists"})
		return
	}

	// Respond with a success message and status code 201 Created
	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   createuserrequest, // You can choose to include the created user data here
	}
	ctx.JSON(http.StatusCreated, webResponse)
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
	var updateuserrequest request.UpdateUserReq
	if err := c.ShouldBindJSON(&updateuserrequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	// Set the ID field of UpdateUserRequest with the id value
	updateuserrequest.Id = id

	// Call the service method to update the user
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

	// Call the userService to delete the user
	err = userController.userService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (userController *UserController) FindAll(c *gin.Context) {
	usersRespose := userController.userService.FindAll()
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "success",
		Data:   usersRespose,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)

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
