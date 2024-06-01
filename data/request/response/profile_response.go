package response

type ProfileResponse struct {
	ProfileId   int    `json:"profileId"`
	ProfileName string `json:"profileName"`
	Age         int    `json:"age"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
}
