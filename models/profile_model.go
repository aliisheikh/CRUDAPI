package Models

type ProfileModel struct {
	ProfileId   int    `gorm:"primary_key"`
	ProfileName string `gorm:"type:varchar(255);column:profile_name"`
	Age         int    `gorm:"type:int(11);column:age"`
	Address     string `gorm:"type:varchar(255);column:address"`
	Phone       string `gorm:"type:varchar(255);column:phone"`
	UserID      int    `gorm:"foreign_key:Id"`
}
