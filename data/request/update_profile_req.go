package request

type UpdateProfileReq struct {
	ProfileId   int    `validation:"required" json:"profileId"`
	ProfileName string `validation:"required,profile_name" json:"profileName"`
	Phone       string `validation:"required,phone" json:"phone"`
	Address     string `validation:"required,address" json:"address"`
	UserId      int    `validation:"required" json:"userId"`
}
