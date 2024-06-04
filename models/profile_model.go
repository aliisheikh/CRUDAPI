package Models

type ProfileModel struct {
	//gorm.Model
	ProfileId   int    `gorm:"primary_key"`
	ProfileName string `gorm:"type:varchar(255);column:profileName"`
	//Age         string `gorm:"type:varchar(11);column:age"`
	Address string `gorm:"type:varchar(255);column:address"`
	Phone   string `gorm:"type:varchar(255);column:phone"`
	UserID  int    `gorm:"column:user_Id"`
}

func (ProfileModel) TableName() string {
	return "profile" // Set the exact table name here
}
