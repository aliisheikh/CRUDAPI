package Models

import (
	"fmt"
)

type User struct {
	//gorm.Model
	Id int `gorm:"type:int;primary_key;column:userId"`
	//UserName string `gorm:"type:varchar(255);column:username"`
	Email   string         `gorm:"type:varchar(255);column:email;unique"`
	Name    string         `gorm:"type:varchar(255);column:name"`
	profile []ProfileModel `gorm:"foreignkey:UserID"`
}

func Model() {
	fmt.Println("model")
}
