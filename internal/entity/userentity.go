package entity

import "github.com/golang-jwt/jwt/v5"

type TokenInfo struct {
	JMap jwt.MapClaims
	UID  string // 用户id
	TID  string // 团队id
}

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

type RegisterByUidRep struct {
	LoginRep
	Uid string `json:"uid"`
}

type LoginByEmailReq struct {
	Email string `json:"email" validate:"required,email,min=6"`
	Pwd   string `json:"pwd" validate:"required,min=6,max=20"`
}

type LoginByPhoneReq struct {
	Phone string `json:"phone" validate:"required,min=6"`
	Pwd   string `json:"pwd" validate:"required,min=6,max=20"`
}

type LoginByUidReq struct {
	Uid string `json:"uid" validate:"re"`
}

type LoginRep struct {
	Token string `json:"token"`
}

type BindEmailReq struct {
	Email string `json:"email" validate:"required,email,min=6"`
}

type BindPhoneReq struct {
	Phone string `json:"phone" validate:"required,min=6"`
}
