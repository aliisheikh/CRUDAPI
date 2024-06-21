package Models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type User struct {
	//gorm.Model
	Id int `gorm:"type:int;primary_key;column:userId"`
	//UserName string `gorm:"type:varchar(255);column:username"`
	Email   string         `gorm:"type:varchar(255);column:email;unique" binding:"required,min=2,max=200"`
	Name    string         `gorm:"type:varchar(255);column:name" binding:"required,min=2,max=200"`
	profile []ProfileModel `gorm:"foreignkey:UserID"`
}

func Model() {
	fmt.Println("model")
}

func CreateUser(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		fieldname := err.(validator.ValidationErrors)[0].Field()
		c.JSON(http.StatusBadRequest, gin.H{"error": fieldname + " " + err.Error()})
		return
	}

	// Process the request
}
