package response

type UserResponse struct {
	Id int `json:"id"`
	//UserName string `json:"username"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
