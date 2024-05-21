package request

type UpdateUserReq struct {
	Id    int    `validation:"required" json:"id"`
	Name  string `validation:"required,max=200,min=1" json:"name"`
	Email string `validation:"required,email" json:"email"`
}