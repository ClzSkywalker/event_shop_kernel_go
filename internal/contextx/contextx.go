package contextx

import (
	"context"

	"gorm.io/gorm"
)

// 自定义context
// 全局参数

type Contextx struct {
	context.Context `json:"-"`
	Language        string   `json:"language,omitempty"`
	TID             string   `json:"tid,omitempty"` // team id
	UID             string   `json:"uid,omitempty"` // user id
	Tx              *gorm.DB `json:"-"`
}
