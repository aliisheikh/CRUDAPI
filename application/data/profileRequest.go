package data

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

func validateUserID(fl validator.FieldLevel) bool {
	if fl.Parent().Kind() == reflect.Struct {
		// Get the struct field by name
		field, _ := fl.Parent().Type().FieldByName(fl.FieldName())
		// If the field is UserID and its value is zero, skip validation
		if field.Name == "UserId" && fl.Field().Int() == 0 {
			return true
		}
	}
	return false
}

type CreateProfileReq struct {
	ProfileName string `json:"profileName" binding:"required,min=2,max=200"`
	//Age         string `json:"age" binding:"required,min=1,max=150"`
	Address string `json:"address" binding:"required,min=1,max=200"`
	Phone   string `json:"phone" binding:"required,min=1,max=20"`
	UserId  int    `json:"userId" binding:"-"`
}

func CreateProfile(c *gin.Context) {
	var req CreateProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			// Extract the first validation error
			fieldName := verr[0].Field()
			// Return the validation error message
			c.JSON(http.StatusBadRequest, gin.H{"error": fieldName + " " + err.Error()})
			return
		}
		// Handle non-validation errors
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		return
	}

	// Your logic to create the profile goes here...
	// You can access the validated profile data from the 'req' variable
	c.JSON(http.StatusCreated, gin.H{"message": "Profile created successfully"})
}
