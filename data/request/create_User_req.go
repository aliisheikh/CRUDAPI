package request

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserReq struct {
	Name  string `json:"name" binding:"required,min=2,max=200"`
	Email string `json:"email" binding:"required,min=2,max=200"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process the request
}
