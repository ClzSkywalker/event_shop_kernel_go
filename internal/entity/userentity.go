package entity

type registerBody struct {
	NickName string `json:"nick_name" validate:"required,min=3,max=20"`
	Pwd      string `json:"pwd" validate:"required,min=6,max=20"`
}

type RegisterByEmailReq struct {
	registerBody
	Email string `json:"email" validate:"required,email,min=6,max=35"`
}

type RegisterByPhoneReq struct {
	registerBody
	Phone string `json:"phone" validate:"required,email,min=6,max=35"`
}

type LoginByEmailReq struct {
	Email string `json:"email" validate:"required,email,min=6"`
	Pwd   string `json:"pwd" validate:"required,min=6,max=20"`
}

type LoginRep struct {
	Token string `json:"token"`
}
