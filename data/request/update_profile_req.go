package request

type UpdateProfileReq struct {
	ProfileId   int    `validation:"required" json:"profileId"`
	ProfileName string `validation:"required,email" json:"profileName"`
	Age         string `validation:"required,age" json:"age"`
	Phone       string `validation:"required,phone" json:"phone"`
	Address     string `validation:"required,address" json:"address"`
}
