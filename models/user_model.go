package Models

import "fmt"

type User struct {
	Id       int       `gorm:"type:int;primary_key"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Email    string    `gorm:"type:varchar(255);not null"`
	profiles []profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type profile struct {
	ID    int    `gorm:"type:int;primary_key"`
	Name  string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255);not null"`
}

func Model() {
	fmt.Println("model")

}
