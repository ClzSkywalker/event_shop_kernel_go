package entity

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/golang-jwt/jwt/v5"
)

type TokenInfo struct {
	JMap jwt.MapClaims `json:"-"`
	UID  string        `json:"uid"` // 用户id
	TID  string        `json:"tid"` // 团队id
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

type UserItem struct {
	CreatedBy    string              `json:"created_by,omitempty"`
	TeamIdPort   string              `json:"team_id_port,omitempty"`
	NickName     string              `json:"nick_name,omitempty"`
	MemberType   constx.UserType     `json:"member_type,omitempty"`
	RegisterType constx.RegisterTypt `json:"register_type,omitempty"`
	Picture      string              `json:"picture,omitempty"`
	Email        string              `json:"email,omitempty"`
	Phone        string              `json:"phone,omitempty"`
	Version      string              `json:"version,omitempty"`
}

type UserResp struct {
	UserItem
}

type UserUpdateReq struct {
	UserItem
}
