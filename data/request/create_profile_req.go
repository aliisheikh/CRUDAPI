package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CreateProfileReq struct {
	ProfileName string `json:"profileName" binding:"required,min=2,max=200"`
	//Age         string `json:"age" binding:"required,min=1,max=150"`
	Address string `json:"address" binding:"required,min=1,max=200"`
	Phone   string `json:"phone" binding:"required,min=1,max=20"`
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
