package Models

type ProfileModel struct {
	ProfileId   int    `gorm:"primary_key;column:Id"`
	ProfileName string `gorm:"type:varchar(255);column:profileName" binding:"required,min=2,max=200"`
	Address     string `gorm:"type:varchar(255);column:address" binding:"required,min=1,max=200"`
	Phone       string `gorm:"type:varchar(255);column:phone" binding:"required,min=1,max=20"`
	UserID      int    `gorm:"type:int;column:userId;foreignkey:UserId;constraint:OnDelete:CASCADE;"`
	UserId      int    `validation:"required" json:"userId"`
	User        User   `gorm:"foreignkey:userId;constraint:OnDelete:CASCADE;"` // Added constraint for cascade deletion

}

func (ProfileModel) TableName() string {

	return "profiles"
}
