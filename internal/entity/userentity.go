package entity

type RegisterEmailReq struct {
	NickName string `json:"nick_name" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email,min=6,max=35"`
	Pwd      string `json:"pwd" validate:"required,min=6,max=20"`
}

type LoginByEmailReq struct {
	Email string `json:"email" validate:"required,email,min=6"`
	Pwd   string `json:"pwd" validate:"required,min=6,max=20"`
}

type LoginRep struct {
	Token string `json:"token"`
}
