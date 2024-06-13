package Models

type ProfileModel struct {
	ProfileId   int    `gorm:"primary_key;column:Id"`
	ProfileName string `gorm:"type:varchar(255);column:profileName"`
	Address     string `gorm:"type:varchar(255);column:address"`
	Phone       string `gorm:"type:varchar(255);column:phone"`
	UserID      int    `gorm:"type:int;column:userId;foreignkey:UserId;constraint:OnDelete:CASCADE;"`
	User        User   `gorm:"foreignkey:userId;constraint:OnDelete:CASCADE;"` // Added constraint for cascade deletion
}

func (ProfileModel) TableName() string {

	return "profiles"
}
