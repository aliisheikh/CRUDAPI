package Models

import (
	"fmt"
)

type User struct {
	Id       int    `gorm:"type:int;primary_key"`
	UserName string `gorm:"type:varchar(255);column:username"`
	Email    string `gorm:"type:varchar(255);column:email;unique"`
	Name     string `gorm:"type:varchar(255);column:name"`
}

//
//profiles []profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

//
//type profile struct {
//	ID    int    `gorm:"type:int;primary_key"`
//	Name  string `gorm:"type:varchar(255);not null"`
//	Email string `gorm:"type:varchar(255);not null"`
//}

func Model() {
	fmt.Println("model")

}
