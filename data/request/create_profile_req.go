package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CreateProfileReq struct {
	ProfileName string `json:"profileName" binding:"required,min=2,max=200"`
	Age         int    `json:"age" binding:"required,min=1,max=20"`
	Address     string `json:"address" binding:"required,min=1,max=200"`
	Phone       string `json:"phone" binding:"required,min=1,max=20"`
}

func CreateProfile(c *gin.Context) {
	var req CreateProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fieldName := err.(validator.ValidationErrors)[0].Field()
		c.JSON(http.StatusBadRequest, gin.H{"error": fieldName + " " + err.Error()})
		return
	}
}
