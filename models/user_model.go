package Models

import (
	"fmt"
)

type User struct {
	Id int `gorm:"type:int;primary_key"`
	//UserName string `gorm:"type:varchar(255);column:username"`
	Email    string `gorm:"type:varchar(255);column:email;unique"`
	Name     string `gorm:"type:varchar(255);column:name"`
	Profiles []ProfileModel
}

func Model() {
	fmt.Println("model")

}
