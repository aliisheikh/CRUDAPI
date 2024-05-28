package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CreateUserReq struct {
	UserName string `json:"username" binding:"required,min=2,max=200"`
	Email    string `json:"email" binding:"required,min=2,max=200"`
	Name     string `json:"name" binding:"required,min=2,max=200"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		fieldname := err.(validator.ValidationErrors)[0].Field()
		c.JSON(http.StatusBadRequest, gin.H{"error": fieldname + " " + err.Error()})
		return
	}

	// Process the request
}
