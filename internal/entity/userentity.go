package entity

type RegisterEmailReq struct {
	NickName string `json:"nick_name" validator:"required,min=3,max=20"`
	Email    string `json:"email" validator:"required,email,min=6,max=35"`
	Pwd      string `json:"pwd" validator:"required,min=6,max=20"`
}

type LoginByEmailReq struct {
	Email string `json:"email" validator:"required,email,min=6"`
	Pwd   string `json:"pwd" validator:"required,min=6,max=20"`
}

type LoginRep struct {
	Token string `json:"token"`
}
